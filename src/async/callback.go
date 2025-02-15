package async

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type CallbackErrorType string

const (
	CallbackDoesNotExist     CallbackErrorType = "Callback does not exist"
	CallbackIsBlockedAlready CallbackErrorType = "Callback is blocked already"
)

type CallbackError struct {
	callContext       string
	callbackErrorType CallbackErrorType
}

func NewCallbackError(callbackErrorType CallbackErrorType) *CallbackError {
	pcs := make([]uintptr, 10)
	callersNum := runtime.Callers(3, pcs)
	if callersNum == 0 {
		return &CallbackError{
			callContext:       "",
			callbackErrorType: callbackErrorType,
		}
	}

	pcs = pcs[:callersNum]
	frames := runtime.CallersFrames(pcs)
	var ctx strings.Builder
	for {
		frame, more := frames.Next()
		ctx.WriteString(fmt.Sprintf("Call from %s on line %d\n", frame.Function, frame.Line))
		if !more {
			break
		}
	}

	return &CallbackError{
		callContext:       ctx.String(),
		callbackErrorType: callbackErrorType,
	}
}

func (err *CallbackError) Error() string {
	return fmt.Sprintf("An error %s occurred on:\n %s", err.callbackErrorType, err.callContext)
}

type Blockable[T any] interface {
	Call(T) error
	BlockedCall(T) error
}

type SingleBlockableCallback[T any] struct {
	callback *func(T)
	lock     sync.Mutex
}

func NewSingleBlockableCallback[T any](callback *func(T)) *SingleBlockableCallback[T] {
	return &SingleBlockableCallback[T]{
		callback: callback,
	}
}

func (sbCallback *SingleBlockableCallback[T]) SetNewCallback(callback func(T)) {
	sbCallback.callback = &callback
}

func (sbCallback *SingleBlockableCallback[T]) Call(param T) error {

	if sbCallback == nil {
		return NewCallbackError(CallbackDoesNotExist)
	}

	if sbCallback.callback == nil {
		return NewCallbackError(CallbackDoesNotExist)
	}

	(*sbCallback.callback)(param)

	return nil
}

func (sbCallback *SingleBlockableCallback[T]) BlockedCall(param T) error {

	if sbCallback == nil {
		return NewCallbackError(CallbackDoesNotExist)
	}

	if sbCallback.callback == nil {
		return NewCallbackError(CallbackDoesNotExist)
	}

	if !sbCallback.lock.TryLock() {
		return NewCallbackError(CallbackIsBlockedAlready)
	}

	(*sbCallback.callback)(param)
	sbCallback.lock.Unlock()

	return nil
}
