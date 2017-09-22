package uidmap

import (
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"testing"
)

func TestFind(t *testing.T) {
	var findTests = []struct {
		uid      string
		username string
	}{
		{"0000919f77953a6961b086b579c3db00", "vladionescu"},
		{"008275e07f931b20807c3b81635c6300", "kobak"},
		{"23260c2ce19420f97b58d7d95b68ca00", "chris"},
		{"23157442436087ee8bcb46c5a193b119", "ZarathustraSpoke"},
		{"dbb165b7879fe7b1174df73bed0b9500", "max"},
		{"fffd6589590eaf361af59c6c22c05300", "SteveClement"},
		{"06eb0cf37180f36567a27bd4598ea700", "hicksfilosopher"},
		{"06ebf3277527fdebc08ab8f68c779100", "wileywiggins"},
		{"06ec3f699708c220e7b8126ab084d900", "pdg"},
		{"06ecc3c2bcbebae81187477e5e340800", "fletch"},
		{"06eeaa6bf23da490727dbc57852f2800", "svrist"},
		{"06f0c3aedcd6b8fc9656594108b32300", "peelr"},
		{"06f188a09ce38152f90a6f8353b08900", "jayholler"},
		{"06f19f7fd78c0d2ff5e8d91e5ae08600", "fredrikhegren"},
		{"06f2412ce9baadfc120739bec1664700", "dwradcliffe"},
		{"06f246cc34d13b7f23bb8a53547bb800", "oneofone"},
		{"eeeeeeeeeeeeeeeeeeeeeeeeeeeeee00", ""},
		{"11111111112222233334445556666719", ""},
	}

	for _, findTest := range findTests {
		uid, err := keybase1.UIDFromString(findTest.uid)
		if err != nil {
			t.Fatal(err)
		}
		found := Find(uid)
		if !found.Eq(libkb.NewNormalizedUsername(findTest.username)) {
			t.Fatalf("Failure for %v: %s != %s", uid, findTest.username, found)
		}
	}
}
