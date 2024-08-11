package stack

import (
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New[int32]()
	assert.EqualValues(t, 0, stack.Size())

	count := rand.Uint32N(50) + 50
	elements := make([]int32, 0, count)

	for i := uint32(1); i <= count; i++ {
		element := rand.Int32N(100) - 50
		elements = append(elements, element)
		stack.Push(element)
		assert.EqualValues(t, i, stack.Size())
	}

	slices.Reverse(elements)

	for i, element := range elements {
		stackElement, ok := stack.Pop()
		assert.True(t, ok)
		assert.Equal(t, element, stackElement)
		assert.EqualValues(t, count-uint32(i)-1, stack.Size())
	}

	stackElement, ok := stack.Pop()
	assert.False(t, ok)
	assert.Zero(t, stackElement)
}
