package util

import (
	"fmt"
	textutil "github.com/anytypeio/go-anytype-middleware/util/text"
	"github.com/libp2p/go-libp2p-core/peer"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-threads/core/thread"
)

var log = logging.Logger("anytype-util")

func MultiAddressAddThread(addr ma.Multiaddr, tid thread.ID) (ma.Multiaddr, error) {
	if addr == nil {
		return nil, fmt.Errorf("addr is nil")
	}

	existingThreadId, err := addr.ValueForProtocol(thread.Code)
	if err != ma.ErrProtocolNotFound {
		if existingThreadId == tid.String() {
			return addr, nil
		} else {
			return nil, fmt.Errorf("addr already has another thread component with differnet ID")
		}
	}

	threadComp, err := ma.NewComponent(thread.Name, tid.String())
	if err != nil {
		return nil, err
	}
	return addr.Encapsulate(threadComp), nil
}

func MultiAddressTrimThread(addr ma.Multiaddr) (ma.Multiaddr, error) {
	parts := strings.Split(addr.String(), "/"+thread.Name)
	trimmed, err := ma.NewMultiaddr(parts[0])
	if err != nil {
		return nil, err
	}
	return trimmed, nil
}

func GetLog(logs []thread.LogInfo, id peer.ID) thread.LogInfo {
	for _, l := range logs {
		if l.ID == id {
			return l
		}
	}

	return thread.LogInfo{}
}

func MultiAddressHasReplicator(addrs []ma.Multiaddr, multiaddr ma.Multiaddr) bool {
	replicatorP2PAddr, err := multiaddr.ValueForProtocol(ma.P_P2P)
	if err != nil {
		log.Errorf("invalid replicator addr: %s", multiaddr.String())
		return false
	}

	for _, addr := range addrs {
		p2pAddr, err := addr.ValueForProtocol(ma.P_P2P)
		if err != nil {
			continue
		}
		if p2pAddr == replicatorP2PAddr {
			return true
		}
	}
	return false
}

func MultiAddressesContains(addrs []ma.Multiaddr, addr ma.Multiaddr) bool {
	for _, a := range addrs {
		if a.Equal(addr) {
			return true
		}
	}
	return false
}

func MultiAddressesToStrings(addrs []ma.Multiaddr) []string {
	var s []string
	for _, addr := range addrs {
		s = append(s, addr.String())
	}

	return s
}

func ThreadIdsToStings(ids []thread.ID) []string {
	var s []string
	for _, id := range ids {
		s = append(s, id.String())
	}

	return s
}

func TruncateText(text string, length int) string {
	var ellipsis = " …"
	if textutil.UTF16RuneCountString(text) <= length {
		return text
	}

	var lastWordIndex, lastNonSpace, currentLen, endTextPos int
	for i, r := range text {
		currentLen++
		if unicode.IsSpace(r) {
			lastWordIndex = lastNonSpace
		} else if unicode.In(r, unicode.Han, unicode.Hangul, unicode.Hiragana, unicode.Katakana) {
			lastWordIndex = i
		} else {
			lastNonSpace = i + utf8.RuneLen(r)
		}

		if currentLen > length {
			if lastWordIndex == 0 {
				endTextPos = i
			} else {
				endTextPos = lastWordIndex
			}
			out := text[0:endTextPos]

			return out + ellipsis
		}
	}

	return text
}

// UniqueStrings returns the new slice without duplicates, while preserving the order.
// The second and further occurrences are considered a duplicate
func UniqueStrings(items []string) []string {
	var um = make(map[string]struct{}, len(items))
	var unique = make([]string, 0, len(um))
	for _, item := range items {
		if _, exists := um[item]; exists {
			continue
		}
		unique = append(unique, item)
		um[item] = struct{}{}
	}

	return unique
}
