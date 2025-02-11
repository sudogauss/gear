package async

// ResponseType string enum.
type responseType string

// ResponseType Enum definition.
const (
	OK      responseType = "OK"
	ERR     responseType = "ERR"
	UNKNOWN responseType = "UNKNOWN"
)

// Async response interface definition with any type.
// Unfold method is used to get underlying data.
type response[T any] interface {
	unfold() (responseType, *T)
}

// Struct representing a success async response.
// Underlying data can be of any type.
type ok[T any] struct {
	data T
}

// Struct representing an error async response.
// Underlying data can be of any type.
type err[T any] struct {
	data T
}

// Response interface implementation for the Ok async response.
func (a *ok[T]) unfold() (responseType, *T) {
	if a == nil {
		return UNKNOWN, nil
	}
	return OK, &a.data
}

// Response interface implementation for the Err async response.
func (a *err[T]) unfold() (responseType, *T) {
	if a == nil {
		return UNKNOWN, nil
	}
	return ERR, &a.data
}

// Enum representing Promise completion reason.
type Completion int

// Completion Enum definition.
const (
	Running Completion = iota
	Failed
	Timeout
	Canceled
	Completed
)
