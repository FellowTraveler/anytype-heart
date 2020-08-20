package threads

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/anytypeio/go-anytype-library/core/smartblock"
	"github.com/anytypeio/go-anytype-library/wallet"
	"github.com/anytypeio/go-slip21"
	"github.com/libp2p/go-libp2p-core/crypto"

	corenet "github.com/textileio/go-threads/core/net"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/crypto/symmetric"
)

type threadDerivedIndex uint32

const (
	// profile page is publicly accessible as service/read keys derived from account public key
	threadDerivedIndexProfilePage threadDerivedIndex = 0
	threadDerivedIndexHome        threadDerivedIndex = 1
	threadDerivedIndexArchive     threadDerivedIndex = 2
	threadDerivedIndexAccount     threadDerivedIndex = 3

	threadDerivedIndexSetPages threadDerivedIndex = 20

	anytypeThreadSymmetricKeyPathPrefix = "m/SLIP-0021/anytype"
	// TextileAccountPathFormat is a path format used for Anytype keypair
	// derivation as described in SEP-00XX. Use with `fmt.Sprintf` and `DeriveForPath`.
	// m/SLIP-0021/anytype/<predefined_thread_index>/%d/<label>
	anytypeThreadPathFormat = anytypeThreadSymmetricKeyPathPrefix + `/%d/%s`

	anytypeThreadServiceKeySuffix = `service`
	anytypeThreadReadKeySuffix    = `read`
	anytypeThreadIdKeySuffix      = `id`
)

type DerivedSmartblockIds struct {
	Account  string
	Profile  string
	Home     string
	Archive  string
	SetPages string
}

var threadDerivedIndexToThreadName = map[threadDerivedIndex]string{
	threadDerivedIndexProfilePage: "profile",
	threadDerivedIndexHome:        "home",
	threadDerivedIndexArchive:     "archive",
}
var threadDerivedIndexToSmartblockType = map[threadDerivedIndex]smartblock.SmartBlockType{
	threadDerivedIndexProfilePage: smartblock.SmartBlockTypeProfilePage,
	threadDerivedIndexHome:        smartblock.SmartBlockTypeHome,
	threadDerivedIndexArchive:     smartblock.SmartBlockTypeArchive,
	threadDerivedIndexSetPages:    smartblock.SmartBlockTypeSet,
}
var ErrAddReplicatorsAttemptsExceeded = fmt.Errorf("add replicatorAddr attempts exceeded")

func (s *service) EnsurePredefinedThreads(ctx context.Context, newAccount bool) (DerivedSmartblockIds, error) {
	s.Lock()
	defer s.Unlock()

	ids := DerivedSmartblockIds{}
	// account
	account, justCreated, err := s.derivedThreadEnsure(ctx, threadDerivedIndexAccount, newAccount, false)
	if err != nil {
		return ids, err
	}

	ids.Account = account.ID.String()
	if s.db == nil {
		err = s.threadsDbInit()
		if err != nil {
			return ids, fmt.Errorf("threadsDbInit failed: %w", err)
		}

		err = s.threadsDbMigration(account.ID.String())
		if err != nil {
			return ids, fmt.Errorf("threadsDbMigration failed: %w", err)
		}
	}

	if !newAccount {
		accountThreadPullDone := make(chan struct{})
		// accountSelect common case
		go func() {
			defer close(accountThreadPullDone)
			// pull only after adding collection to handle all events
			_, err = s.pullThread(context.TODO(), account.ID)
			if err != nil {
				log.Errorf("failed to pull accountThread")
			}
		}()

		if justCreated {
			// this is the case of accountSelect after accountRecovery
			// we need to wait for account thread pull to be done
			<-accountThreadPullDone
			if err != nil {
				return ids, err
			}
		}
	}

	// profile
	profile, _, err := s.derivedThreadEnsure(ctx, threadDerivedIndexProfilePage, newAccount, true)
	if err != nil {
		return ids, err
	}
	ids.Profile = profile.ID.String()

	// home
	home, _, err := s.derivedThreadEnsure(ctx, threadDerivedIndexHome, newAccount, true)
	if err != nil {
		return ids, err
	}
	ids.Home = home.ID.String()

	// archive
	archive, _, err := s.derivedThreadEnsure(ctx, threadDerivedIndexArchive, newAccount, true)
	if err != nil {
		return ids, err
	}
	ids.Archive = archive.ID.String()

	// set pages
	setPages, _, err := s.derivedThreadEnsure(ctx, threadDerivedIndexSetPages, newAccount, true)
	if err != nil {
		return ids, err
	}
	ids.SetPages = setPages.ID.String()

	return ids, nil
}

func ProfileThreadIDFromAccountPublicKey(pubk crypto.PubKey) (thread.ID, error) {
	accountPub, err := pubk.Raw()
	if err != nil {
		return thread.Undef, err
	}

	node, err := slip21.DeriveForPath(fmt.Sprintf(anytypeThreadPathFormat, threadDerivedIndexProfilePage, anytypeThreadIdKeySuffix), accountPub)
	if err != nil {
		return thread.Undef, err
	}

	// we use symmetric key because it is also has the size of 32 bytes
	return threadIDFromBytes(thread.Raw, threadDerivedIndexToSmartblockType[threadDerivedIndexProfilePage], node.SymmetricKey())
}

func ProfileThreadKeysFromAccountPublicKey(pubk crypto.PubKey) (service *symmetric.Key, read *symmetric.Key, err error) {
	masterKey, err := pubk.Raw()
	if err != nil {
		return
	}

	return threadDeriveKeys(threadDerivedIndexProfilePage, masterKey)
}

func ProfileThreadKeysFromAccountAddress(address string) (service *symmetric.Key, read *symmetric.Key, err error) {
	pubk, err := wallet.NewPubKeyFromAddress(wallet.KeypairTypeAccount, address)
	if err != nil {
		return
	}

	return ProfileThreadKeysFromAccountPublicKey(pubk)
}

func ProfileThreadIDFromAccountAddress(address string) (thread.ID, error) {
	pubk, err := wallet.NewPubKeyFromAddress(wallet.KeypairTypeAccount, address)
	if err != nil {
		return thread.Undef, err
	}
	return ProfileThreadIDFromAccountPublicKey(pubk)
}

func (s *service) derivedThreadKeyByIndex(index threadDerivedIndex) (service *symmetric.Key, read *symmetric.Key, err error) {
	if index == threadDerivedIndexProfilePage {
		// anyone should be able to read profile
		// so lets derive its encryption keys from the account public key instead
		masterKey, err2 := s.account.GetPublic().Raw()
		if err2 != nil {
			err = err2
			return
		}
		return threadDeriveKeys(index, masterKey)
	}

	var masterKey = make([]byte, 32)
	pkey, err2 := s.account.Raw()
	if err2 != nil {
		err = err2
		return
	}
	copy(masterKey, pkey[:32])

	return threadDeriveKeys(index, masterKey)
}

func (s *service) derivedThreadIdByIndex(index threadDerivedIndex) (thread.ID, error) {
	if s.account == nil {
		return thread.Undef, fmt.Errorf("account key not set")
	}

	if index == threadDerivedIndexProfilePage {
		accountKey, err := s.account.GetPublic().Raw()
		if err != nil {
			return thread.Undef, err
		}

		return threadDeriveId(index, accountKey)
	}

	var masterKey = make([]byte, 32)
	pkey, err := s.account.Raw()
	if err != nil {
		return thread.Undef, err
	}
	copy(masterKey, pkey[:32])

	return threadDeriveId(index, masterKey)
}

func (s *service) derivedThreadByName(name string) (thread.Info, error) {
	for index, tname := range threadDerivedIndexToThreadName {
		if name == tname {
			return s.derivedThreadWithIndex(index)
		}
	}

	return thread.Info{}, fmt.Errorf("thread not found")
}

func (s *service) derivedThreadWithIndex(index threadDerivedIndex) (thread.Info, error) {
	id, err := s.derivedThreadIdByIndex(index)
	if err != nil {
		return thread.Info{}, err
	}

	return s.t.GetThread(context.TODO(), id)
}

func (s *service) derivedThreadEnsure(ctx context.Context, index threadDerivedIndex, newAccount bool, pull bool) (thrd thread.Info, justCreated bool, err error) {
	if newAccount {
		thrd, err := s.derivedThreadCreate(index)
		return thrd, true, err
	}

	return s.derivedThreadAddExistingFromLocalOrRemote(ctx, index, pull)
}

func (s *service) derivedThreadCreate(index threadDerivedIndex) (thread.Info, error) {
	id, err := s.derivedThreadIdByIndex(index)
	if err != nil {
		return thread.Info{}, err
	}

	thrd, err := s.t.GetThread(context.Background(), id)
	if err == nil && thrd.Key.Service() != nil {
		// we already have the thread locally, we can safely pull it in background
		return thrd, nil
	}

	serviceKey, readKey, err := s.derivedThreadKeyByIndex(index)
	if err != nil {
		return thread.Info{}, err
	}

	// intentionally do not pass the original ctx, because we don't want to stuck in the middle of thread creation
	thrd, err = s.t.CreateThread(context.Background(),
		id,
		corenet.WithThreadKey(thread.NewKey(serviceKey, readKey)),
		corenet.WithLogKey(s.device))
	if err != nil {
		return thread.Info{}, err
	}

	// because this thread just have been created locally we can safely put all networking in the background
	go func() {
		if s.replicatorAddr == nil {
			return
		}

		err = s.addReplicatorWithAttempts(context.Background(), thrd, s.replicatorAddr, 0)
		if err != nil {
			log.Warnf("derivedThreadCreate failed to add replicatorAddr: %s", err.Error())
		}
	}()

	return thrd, nil
}

func (s *service) derivedThreadAddExistingFromLocalOrRemote(ctx context.Context, index threadDerivedIndex, pull bool) (info thread.Info, justCreated bool, err error) {
	id, err := s.derivedThreadIdByIndex(index)
	if err != nil {
		return thread.Info{}, false, err
	}

	addReplicatorAnPullAfter := func(thrd thread.Info) {
		if s.replicatorAddr != nil {
			// if thread doesn't yet have s replicatorAddr this function will continuously try to add it in the background
			err = s.addReplicatorWithAttempts(context.Background(), thrd, s.replicatorAddr, 0)
			if err != nil {
				log.Errorf("existing thread failed to add replicatorAddr: ", err.Error())
				return
			}
		}

		if !pull {
			return
		}

		// lets try to pull it once the replicatorAddr have been added
		// in case it fails this thread will be still pulled every PullInterval
		err := s.PullThread(ctx, thrd.ID)
		if err != nil {
			log.Errorf("existing thread failed to pull: ", err.Error())
			return
		}
	}

	thrd, err := s.t.GetThread(ctx, id)
	if err == nil && thrd.Key.Service() != nil {
		// we already have the thread locally, we can safely pull it in background
		go addReplicatorAnPullAfter(thrd)
		return thrd, false, nil
	}

	serviceKey, readKey, err := s.derivedThreadKeyByIndex(index)
	if err != nil {
		return
	}

	// we must recover it from
	// intentionally do not pass the original ctx, because we don't want to stuck in the middle of thread creation
	thrd, err = s.t.CreateThread(context.Background(),
		id,
		corenet.WithThreadKey(thread.NewKey(serviceKey, readKey)),
		corenet.WithLogKey(s.device))
	if err != nil {
		return
	}

	justCreated = true

	if s.replicatorAddr != nil {
		err = s.addReplicatorWithAttempts(ctx, thrd, s.replicatorAddr, 3)
		if err != nil {
			// remove the thread we have just created because we've supposed to successfully pull it from the replicatorAddr
			err2 := s.t.DeleteThread(context.Background(), id)
			if err2 != nil {
				log.Errorf("failed to delete thread: %s", err.Error())
			}
			return
		}
	}

	if !pull {
		return
	}

	err = s.PullThread(ctx, thrd.ID)
	if err != nil {
		// remove the thread we have just created because we've supposed to successfully pull it from the replicatorAddr
		err2 := s.t.DeleteThread(context.Background(), id)
		if err2 != nil {
			log.Errorf("failed to delete thread: %s", err.Error())
		}
		return
	}
	return thrd, true, nil
}

func threadIDFromBytes(variant thread.Variant, blockType smartblock.SmartBlockType, b []byte) (thread.ID, error) {
	blen := len(b)
	// two 8 bytes (max) numbers plus num
	buf := make([]byte, 2*binary.MaxVarintLen64+blen)
	n := binary.PutUvarint(buf, thread.V1)
	n += binary.PutUvarint(buf[n:], uint64(variant))
	n += binary.PutUvarint(buf[n:], uint64(blockType))

	cn := copy(buf[n:], b)
	if cn != blen {
		return thread.Undef, fmt.Errorf("copy length is inconsistent")
	}

	return thread.Cast(buf[:n+blen])
}

func threadCreateID(variant thread.Variant, blockType smartblock.SmartBlockType) (thread.ID, error) {
	rnd := make([]byte, 32)
	_, err := rand.Read(rnd)
	if err != nil {
		panic("random read failed")
	}

	rndlen := len(rnd)

	// two 8 bytes (max) numbers plus rnd
	buf := make([]byte, 2*binary.MaxVarintLen64+rndlen)
	n := binary.PutUvarint(buf, thread.V1)
	n += binary.PutUvarint(buf[n:], uint64(variant))
	n += binary.PutUvarint(buf[n:], uint64(blockType))

	cn := copy(buf[n:], rnd)
	if cn != rndlen {
		return thread.Undef, fmt.Errorf("copy length is inconsistent")
	}

	return thread.Cast(buf[:n+rndlen])
}

// threadDeriveKeys derive service and read encryption keys derived from key
func threadDeriveKeys(index threadDerivedIndex, masterKey []byte) (service *symmetric.Key, read *symmetric.Key, err error) {
	if len(masterKey) != 32 {
		err = fmt.Errorf("masterKey length should be 32 bytes, got %d instead", len(masterKey))
		return
	}

	nodeKey, err2 := slip21.DeriveForPath(fmt.Sprintf(anytypeThreadPathFormat, index, anytypeThreadServiceKeySuffix), masterKey)
	if err2 != nil {
		err = err2
		return
	}

	service, err = symmetric.FromBytes(append(nodeKey.SymmetricKey()))
	if err != nil {
		return
	}

	nodeKey, err = slip21.DeriveForPath(fmt.Sprintf(anytypeThreadPathFormat, index, anytypeThreadReadKeySuffix), masterKey)
	if err != nil {
		return
	}

	read, err = symmetric.FromBytes(nodeKey.SymmetricKey())
	if err != nil {
		return
	}

	return
}

func threadDeriveId(index threadDerivedIndex, accountKey []byte) (thread.ID, error) {
	node, err := slip21.DeriveForPath(fmt.Sprintf(anytypeThreadPathFormat, index, anytypeThreadIdKeySuffix), accountKey)
	if err != nil {
		return thread.Undef, err
	}

	// we use symmetric key because it is also has the size of 32 bytes
	return threadIDFromBytes(thread.Raw, threadDerivedIndexToSmartblockType[index], node.SymmetricKey())
}
