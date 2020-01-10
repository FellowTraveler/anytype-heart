package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	ipfsCore "github.com/ipfs/go-ipfs/core"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/crypto"
	tcore "github.com/textileio/go-textile/core"
	tmobile "github.com/textileio/go-textile/mobile"
	logging2 "github.com/whyrusleeping/go-logging"
)

var log = logging.Logger("anytype-core")

const privateKey = `/key/swarm/psk/1.0.0/
/base16/
fee6e180af8fc354d321fde5c84cab22138f9c62fec0d1bc0e99f4439968b02c`

var BootstrapNodes = []string{
	"/ip4/68.183.2.167/tcp/4001/ipfs/12D3KooWB2Ya2GkLLRSR322Z13ZDZ9LP4fDJxauscYwUMKLFCqaD",
	"/ip4/157.230.124.182/tcp/4001/ipfs/12D3KooWKLLf9Qc6SHaLWNPvx7Tk4AMc9i71CLdnbZuRiFMFMnEf",
}

type PredefinedBlockIds struct {
	Home    string
	Archive string
}

type Anytype struct {
	Textile            *tmobile.Mobile
	predefinedBlockIds PredefinedBlockIds
	logLevels          map[string]string
}

func (a *Anytype) ipfs() *ipfsCore.IpfsNode {
	return a.Textile.Node().Ipfs()
}

func (a *Anytype) textile() *tcore.Textile {
	return a.Textile.Node()
}

// PredefinedBlockIds returns default blocks like home and archive
// ⚠️ Will return empty struct in case it runs before Anytype.Run()
func (a *Anytype) PredefinedBlockIds() PredefinedBlockIds {
	return a.predefinedBlockIds
}

func New(repoPath string, account string) (*Anytype, error) {
	// todo: remove this temp workaround after release of go-ipfs v0.4.23
	crypto.MinRsaKeyBits = 1024

	msg := messenger{}
	tm, err := tmobile.NewTextile(&tmobile.RunConfig{filepath.Join(repoPath, account), false, nil}, &msg)
	if err != nil {
		return nil, err
	}

	levels := os.Getenv("ANYTYPE_LOG_LEVEL")
	logLevels := make(map[string]string)
	if levels != "" {
		for _, level := range strings.Split(levels, ";") {
			parts := strings.Split(level, "=")
			if len(parts) == 1 {
				for _, subsystem := range logging.GetSubsystems() {
					if strings.HasPrefix(subsystem, "anytype-") {
						logLevels[subsystem] = parts[0]
					}
				}
			} else if len(parts) == 2 {
				logLevels[parts[0]] = parts[1]
			}
		}
	}

	return &Anytype{Textile: tm, logLevels: logLevels}, nil
}

func (a *Anytype) SetLogLevel(subsystem string, level string) {
	a.logLevels[subsystem] = strings.ToUpper(level)

	if a.Textile.Node().Started() {
		a.applyLogLevel()
	}
}

func (a *Anytype) applyLogLevel() {
	if len(a.logLevels) == 0 {
		logging.SetAllLoggers(logging2.ERROR)
		return
	}

	for subsystem, level := range a.logLevels {
		err := logging.SetLogLevel(subsystem, level)
		if err != nil {
			log.Fatalf("incorrect log level for %s: %s", subsystem, level)
		}
	}
}

func (a *Anytype) Run() error {
	swarmKeyFilePath := filepath.Join(a.textile().RepoPath(), "swarm.key")
	err := ioutil.WriteFile(swarmKeyFilePath, []byte(privateKey), 0644)
	if err != nil {
		return err
	}

	err = a.Textile.Start()
	if err != nil {
		return err
	}
	a.applyLogLevel()

	err = a.ipfs().Repo.SetConfigKey("Addresses.Bootstrap", BootstrapNodes)
	if err != nil {
		return err
	}

	go func() {
		for {
			if !a.textile().Started() {
				break
			}

			if !a.ipfs().IsOnline {
				time.Sleep(time.Second)
				continue
			}

			_, err = a.textile().RegisterCafe("12D3KooWB2Ya2GkLLRSR322Z13ZDZ9LP4fDJxauscYwUMKLFCqaD", "2MsR9h7mfq53oNt8vh7RfdPr57qPsn28X3dwbviZWs3E8kEu6kpdcDHyMx7Qo")
			if err != nil {
				log.Errorf("failed to register cafe: %s", err.Error())
				time.Sleep(time.Second * 5)
				continue
			}
			break
		}
	}()

	/*tgateway.Host = &tgateway.Gateway{
		Node: a.Textile.Node(),
	}
	tgateway.Host.Start(a.Textile.Node().Config().Addresses.Gateway)
	fmt.Println("Textile Gateway: " + a.Textile.Node().Config().Addresses.Gateway)*/

	err = a.createPredefinedBlocks()
	if err != nil {
		return err
	}

	return nil
}

func (a *Anytype) Stop() error {
	return a.textile().Stop()
}
