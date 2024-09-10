package spinlock

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	flag atomic.Uint32
}

func (lock *SpinLock) TryLock() bool {
	return lock.flag.CompareAndSwap(0, 1)
}

func (lock *SpinLock) Lock() {
	for {
		if lock.flag.CompareAndSwap(0, 1) {
			return
		}

		runtime.Gosched()
	}
}

func (lock *SpinLock) Unlock() {
	lock.flag.Store(0)
}
