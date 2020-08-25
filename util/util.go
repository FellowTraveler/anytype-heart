package util

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/anytypeio/go-anytype-library/logging"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-threads/core/thread"
)

var log = logging.Logger("anytype-util")

func MultiAddressAddThread(addr ma.Multiaddr, tid thread.ID) (ma.Multiaddr, error) {
	if addr == nil {
		return nil, fmt.Errorf("addr is nil")
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

func MultiAddressHasReplicator(addrs []ma.Multiaddr, multiaddr ma.Multiaddr) bool {
	for _, addr := range addrs {
		addr, err := MultiAddressTrimThread(addr)
		if err != nil {
			log.Error("failed to trim multiaddr: %s", err.Error())
			continue
		}

		if addr.Equal(multiaddr) {
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

func TruncateText(text string, length int) string {
	var ellipsis = " …"
	if utf8.RuneCountInString(text) <= length {
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

func UniqueStrings(items []string) []string {
	var um = make(map[string]struct{}, len(items))
	for _, item := range items {
		um[item] = struct{}{}
	}

	var unique = make([]string, 0, len(um))
	for item := range um {
		unique = append(unique, item)
	}

	return unique
}

func NewImmediateTicker(d time.Duration) *immediateTicker {
	c := make(chan time.Time, 1)
	s := make(chan struct{})

	ticker := time.NewTicker(d)
	c <- time.Now()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				c <- t
			case <-s:
				ticker.Stop()
				return
			}
		}
	}()

	return &immediateTicker{C: c, s: s}
}

type immediateTicker struct {
	C    chan time.Time
	s    chan struct{}
	stop sync.Once
}

func (t *immediateTicker) Stop() {
	t.stop.Do(func() { close(t.s) })
}
