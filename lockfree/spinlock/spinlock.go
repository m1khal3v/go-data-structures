package spinlock

import "sync/atomic"

type SpinLock struct {
	atomic.Uintptr
}

func (lock *SpinLock) TryLock() bool {
	return lock.CompareAndSwap(0, 1)
}

func (lock *SpinLock) Lock() {
	for {
		if lock.CompareAndSwap(0, 1) {
			return
		}
	}
}

func (lock *SpinLock) Unlock() {
	lock.Store(0)
}
