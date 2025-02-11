package pair

import "golang.org/x/exp/constraints"

type Pair[K constraints.Ordered, V any] struct {
	First  K
	Second V
}

func New[K constraints.Ordered, V any](_First K, _Second V) *Pair[K, V] {
	return &Pair[K, V]{
		First:  _First,
		Second: _Second,
	}
}

func (p *Pair[K, V]) Compare(other *Pair[K, V]) int {

	if p.First == other.First {
		return 0
	} else if p.First < other.First {
		return -1
	} else {
		return 1
	}
}
