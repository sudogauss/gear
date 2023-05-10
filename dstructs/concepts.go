package dstructs

// Comparable interface determines whether data structure can be
// compared with another one or not.
type Comparable interface {
	Compare(interface{}) int
}

// Stringable interface determines whether data structure can be
// represented as string or not.
type Stringable interface {
	ToString(interface{}) string
}
