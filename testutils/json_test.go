package testutils

import (
	"encoding/json"
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func TestAssertJsonEquals(t *testing.T) {
	profile := randomdata.GenerateProfile(randomdata.Male)
	copyProfile := profile
	t1, err := json.Marshal(profile)
	if err != nil {
		t.Error(err)
	}
	t2, err := json.Marshal(copyProfile)
	if err != nil {
		t.Error(err)
	}
	AssertJSONEquals(t1, t2, t)

	// a profile with female cannot euqla male
	notEqualProfile := randomdata.GenerateProfile(randomdata.Female)
	ne, err := json.Marshal(notEqualProfile)
	if err != nil {
		t.Error(err)
	}
	brokenTesting := testing.T{}
	// not euqls
	AssertJSONEquals(t1, ne, &brokenTesting)
	if !brokenTesting.Failed() {
		t.Error("expected test to fail")
	}
}
