package queue

import (
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New[int32]()
	assert.EqualValues(t, 0, queue.Size())

	count := rand.Uint32N(50) + 50
	elements := make([]int32, 0, count)

	for i := uint32(1); i <= count; i++ {
		element := rand.Int32N(100) - 50
		elements = append(elements, element)
		queue.Push(element)
		assert.EqualValues(t, i, queue.Size())
	}

	for i, element := range elements {
		queueElement, ok := queue.Pop()
		assert.True(t, ok)
		assert.Equal(t, element, queueElement)
		assert.EqualValues(t, count-uint32(i)-1, queue.Size())
	}

	queueElement, ok := queue.Pop()
	assert.False(t, ok)
	assert.Zero(t, queueElement)
}
