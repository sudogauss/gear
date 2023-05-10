package dsturcts_go

import (
	"sudogauss/gsu/dstructs"
	"testing"
)

// Tests if pair is created correctly
func TestCreatePair(t *testing.T) {
	_pair := dstructs.NewPair(0, "test")
	if _pair.First != 0 {
		t.Errorf("Pair key is equal to %d, expected 0", _pair.First)
	}
	if _pair.Second != "test" {
		t.Errorf("Pair value is equal to %s, expected 0", _pair.Second)
	}
}
