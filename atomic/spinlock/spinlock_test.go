package spinlock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpinLock(t *testing.T) {
	var lock SpinLock
	lock.Lock()
	assert.False(t, lock.TryLock())
	lock.Unlock()
	assert.True(t, lock.TryLock())
	assert.False(t, lock.TryLock())
	lock.Unlock()
}
