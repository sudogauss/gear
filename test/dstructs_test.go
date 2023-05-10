package dstructs_test

import (
	"sudogauss/gsu/dstructs"
	"testing"
)

// Tests whether pair is created correctly
func TestCreatePair(t *testing.T) {
	_pair := dstructs.NewPair(0, "test")
	if _pair.First != 0 {
		t.Errorf("Pair key is equal to %d, expected 0", _pair.First)
	}
	if _pair.Second != "test" {
		t.Errorf("Pair value is equal to %s, expected 0", _pair.Second)
	}
}

// Tests whether pair is stringified correctly
func TestToStringPair(t *testing.T) {
	_pair := dstructs.NewPair(0, "test")
	_str := _pair.ToString()
	if _str != "{0, test}" {
		t.Errorf("Pair was not correctly stringified. Expected {0, test}, got %s", _str)
	}
}

// Tests whether pair keys are compared correctly
func TestComparePairs(t *testing.T) {
	_first_pair := dstructs.NewPair("a_key", "test")
	_second_pair := dstructs.NewPair("b_key", "test")
	_third_pair := dstructs.NewPair("c_key", "test")

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func(c chan int) {
		c <- _second_pair.Compare(_second_pair)
	}(c1)

	go func(c chan int) {
		c <- _second_pair.Compare(_first_pair)
	}(c2)

	go func(c chan int) {
		c <- _second_pair.Compare(_third_pair)
	}(c3)

	_expect_eq, _expect_gt, _expect_lt := <-c1, <-c2, <-c3

	if _expect_eq != 0 {
		t.Errorf("Expected %s == %s, output %d", _second_pair.ToString(), _second_pair.ToString(), _expect_eq)
	}

	if _expect_gt != 1 {
		t.Errorf("Expected %s > %s, output %d", _second_pair.ToString(), _first_pair.ToString(), _expect_gt)
	}

	if _expect_lt != -1 {
		t.Errorf("Expected %s < %s, output %d", _second_pair.ToString(), _third_pair.ToString(), _expect_lt)
	}
}
