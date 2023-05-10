package dstructs

import "golang.org/x/exp/constraints"

// Basic pair representation with two elements: key and value.
// Key [First K] is constrainted to be ordered.
// Value [Second V] can be any type.
type Pair[K constraints.Ordered, V any] struct {
	First  K
	Second V
}

// Allocates new Pair[K, V] of key type K and value type V and returns
// a *Pair[K, V] pointer to this pair.
// Pair's key is equal to _First, pair's value is equal to _Second.
func NewPair[K constraints.Ordered, V any](_First K, _Second V) *Pair[K, V] {
	return &Pair[K, V]{
		First:  _First,
		Second: _Second,
	}
}

// Compares a pair with another pair by comparing their keys of type
// K which is constrainted to be ordered.
// Returned int indicates the comparison result:
// p == other => 0
// p < other => -1
// p > other => 1
func (p *Pair[K, V]) Compare(other *Pair[K, V]) int {

	if p.First == other.First {
		return 0
	} else if p.First < other.First {
		return -1
	} else {
		return 1
	}
}
