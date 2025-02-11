package async

import "time"

type Keeper struct{}

func AsyncLaunch[D any, E any](f func() (*D, *E), timeout int64) *Promise[D, E] {

	p := NewPromise[D, E](0, 0, true, timeout)

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
			p.completionState = Completed
		case err := <-p.errorChannel:
			p.err = &err
			p.completionState = Failed
		case <-p.cancelationChannel:
			p.completionState = Canceled
		case <-time.After(p.timeout):
			p.completionState = Timeout
		}

		p.completionChannel <- p.completionState
	}()

	return p
}
