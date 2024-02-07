package mutex

import (
	"sync"
)

type AtomicInt struct {
	value int
	mx    sync.Mutex
}

func (a *AtomicInt) Add(i int) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.value += i
}
