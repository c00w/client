package uidmap

import (
	"bytes"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"sort"
)

var offsets []uint32

func init() {
	var offset uint32
	offsets = make([]uint32, len(lengths), len(lengths))
	for index, length := range lengths {
		offsets[index] = offset
		offset += uint32(length)
	}
}

func Find(uid keybase1.UID) libkb.NormalizedUsername {
	searchFor := uid.ToBytes()
	uidLen := len(searchFor)
	l := len(uids) / uidLen
	uidAt := func(i int) []byte {
		start := i * uidLen
		return uids[start : start+uidLen]
	}
	doCmp := func(i int) int {
		return bytes.Compare(searchFor, uidAt(i))
	}
	n := sort.Search(l, func(i int) bool {
		return doCmp(i) <= 0
	})
	if n == l || doCmp(n) != 0 {
		return libkb.NormalizedUsername("")
	}
	offset := offsets[n]
	s := usernames[offset : offset+uint32(lengths[n])]
	return libkb.NewNormalizedUsername(s)
}
