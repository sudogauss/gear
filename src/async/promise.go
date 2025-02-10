package async

import (
	"fmt"
	"time"
)

type Promise[D any, E any] struct {
	// Routine            func() (D, E)
	completionChannel   chan (Completion)
	interruptionChannel chan (bool)
	isCompleted         bool
	isInterrupted       bool
	timeout             time.Duration
	dataChannel         chan (D)
	errorChannel        chan (E)
	data                *D
	err                 *E
	completionCallback  *func(D)
	exceptionCallback   *func(E)
}

func NewPromise[D any, E any](completionChan chan (Completion), interruptionChannel chan (bool), timeout int64) *Promise[D, E] {
	return &Promise[D, E]{
		completionChannel:   completionChan,
		interruptionChannel: interruptionChannel,
		isCompleted:         false,
		isInterrupted:       false,
		timeout:             time.Duration(timeout),
		dataChannel:         make(chan D),
		errorChannel:        make(chan E),
		data:                nil,
		err:                 nil,
		completionCallback:  nil,
		exceptionCallback:   nil,
	}
}

func (p *Promise[D, E]) IsCompleted() bool {
	return p.isCompleted
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
