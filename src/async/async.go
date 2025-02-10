package async

import "time"

type Keeper struct{}

func AsyncLaunch[D any, E any](f func() (*D, *E), timeout int64) *Promise[D, E] {
	completionChan := make(chan Completion)
	interruptionChan := make(chan bool)
	p := NewPromise[D, E](completionChan, interruptionChan, timeout)

	go func() {
		d, e := f()
		if e != nil {
			p.errorChannel <- *e
		} else {
			p.dataChannel <- *d
		}
	}()

	go func() {
		select {
		case data := <-p.dataChannel:
			p.data = &data
			p.isCompleted = true
			completionChan <- Completed
		case err := <-p.errorChannel:
			p.err = &err
			p.isCompleted = true
			completionChan <- Failed
		case <-p.interruptionChannel:
			p.isInterrupted = true
			completionChan <- UserInterrupt
		case <-time.After(p.timeout):
			completionChan <- Timeout
		}
	}()

	return p
}
