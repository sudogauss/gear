package async

import (
	"fmt"
	"time"
)

type Promise[D any, E any] struct {
	id                 uint64
	completionChannel  chan (Completion)
	completionState    Completion
	cancelationChannel chan (bool)
	cancellable        bool
	timeout            time.Duration
	dataChannel        chan (D)
	errorChannel       chan (E)
	data               *D
	err                *E
	completionCallback *func(D)
	exceptionCallback  *func(E)
}

func NewPromise[D any, E any](id uint64, threadId uint64, cancellable bool, timeout int64) *Promise[D, E] {
	p := &Promise[D, E]{
		id:                 id,
		completionChannel:  make(chan Completion),
		cancelationChannel: make(chan bool),
		cancellable:        cancellable,
		isCompleted:        false,
		isCanceled:         false,
		timeout:            time.Duration(timeout),
		dataChannel:        make(chan D),
		errorChannel:       make(chan E),
		data:               nil,
		err:                nil,
		completionCallback: nil,
		exceptionCallback:  nil,
	}

	go func() {
		res := <-p.completionChannel
	}()

	if p.cancellable {
		go func() {
			c := <-p.cancelationChannel
			if c {
				// Implement cancelation logic here
				// Cancel thread by id
				// Note that Promise routine goroutine is locked on thread
			}
		}()
	}

	return p
}

func (p *Promise[D, E]) GetId() uint64 {
	return p.id
}

func (p *Promise[D, E]) IsCompleted() bool {
	return p.isCompleted
}

func (p *Promise[D, E]) IsCanceled() bool {
	return p.isCanceled
}

func (p *Promise[D, E]) GetCompletionChanne() chan Completion {
	return p.completionChannel
}

func (p *Promise[D, E]) GetCancelationChannel() chan bool {
	if p.cancellable {
		return nil
	} else {
		return p.cancelationChannel
	}
}

func (p *Promise[D, E]) GetData() *D {
	return p.data
}

func (p *Promise[D, E]) GetError() *E {
	return p.err
}

func (p *Promise[D, E]) GetTimeout() int64 {
	return p.timeout.Nanoseconds()
}

func (p *Promise[D, E]) Cancel() error {
	if p.cancellable {
		p.cancelationChannel <- true
		return nil
	} else {
		return fmt.Errorf("Promise %d is not cancellable")
	}
}

func (p *Promise[D, E]) Then(completionCallback func(D)) *Promise[D, E] {
	if p.completionCallback == nil {
		p.completionCallback = &completionCallback
	} else {
		fmt.Println("Warning: you cannot apply more than one completion callback")
	}

	// go func() {
	// 	select {
	// 		case
	// 	}
	// }()
	return p
}

func (p *Promise[D, E]) Except(exceptionCallback func(E)) *Promise[D, E] {
	if p.exceptionCallback == nil {
		p.exceptionCallback = &exceptionCallback
	} else {
		fmt.Println("Warning: you cannot apply more than one completion callback")
	}
	return p
}
