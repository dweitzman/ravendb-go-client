package ravendb

// CompletableFuture helps porting Java code
type CompletableFuture struct {
	// if > 1, it has been finished
	done AtomicInteger
	// used to wait for Future to finish
	chDone chan bool
	// result generated by the Future, only valid if completed
	result interface{}
	err    error
}

func NewCompletableFuture() *CompletableFuture {
	return &CompletableFuture{
		// channel with capacity 1 so that markAsDone() can finish the goroutine
		// without waiting for someone to call get()
		chDone: make(chan bool, 1),
	}
}

func NewCompletableFutureAlreadyCompleted(result interface{}) *CompletableFuture {
	res := NewCompletableFuture()
	res.markAsDone(result)
	return res
}

func (f *CompletableFuture) isDone() bool {
	return f.done.get() > 0
}

func (f *CompletableFuture) markAsDone(result interface{}) {
	f.result = result
	f.done.set(1)
	f.chDone <- true
}

func (f *CompletableFuture) markAsDoneWithError(err error) {
	f.err = err
	f.done.set(1)
	f.chDone <- true
}

func (f *CompletableFuture) get() (interface{}, error) {
	if f.isDone() {
		return f.result, f.err
	}
	// wait for the Future to finish
	<-f.chDone
	return f.result, f.err
}
