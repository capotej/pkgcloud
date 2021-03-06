package pkgcloud

import "testing"

func TestDistroID_deb(t *testing.T) {
	var tests = map[string]int{
		"ubuntu/warty":    1,
		"ubuntu/trusty":   20,
		"debian/jessie":   25,
		"any/any":         35,
		"linuxmint/petra": 157,
		"raspbian/buster": 156,
	}
	for name, id := range tests {
		actualID, _ := distroID(".deb", name)
		if actualID != id {
			t.Errorf("distro id of %s != %d", name, id)
		}
	}
}

func TestDistroID_rpm(t *testing.T) {
	var tests = map[string]int{
		"el/7":         140,
		"fedora/22":    147,
		"scientific/5": 138,
		"ol/7":         146,
	}
	for name, id := range tests {
		actualID, _ := distroID(".rpm", name)
		if actualID != id {
			t.Errorf("distro id of %s != %d", name, id)
		}
	}
}
