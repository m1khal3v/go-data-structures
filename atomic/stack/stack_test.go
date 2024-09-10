package stack

import (
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"slices"
	"sync"
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

func TestStackWithConcurrency(t *testing.T) {
	stack := New[int32]()
	assert.EqualValues(t, 0, stack.Size())

	count := rand.Uint32N(50) + 50
	elements := make([]int32, 0, count)

	for i := uint32(1); i <= count; i++ {
		element := rand.Int32N(100) - 50
		elements = append(elements, element)
	}

	var wg sync.WaitGroup
	stackElementsChan := make(chan int32, count)
	for _, element := range elements {
		wg.Add(2)
		go func() {
			defer wg.Done()
			stack.Push(element)
		}()
		go func() {
			defer wg.Done()
			for {
				element, ok := stack.Pop()
				if ok {
					stackElementsChan <- element
					return
				}
			}
		}()
	}
	wg.Wait()
	close(stackElementsChan)
	assert.EqualValues(t, 0, stack.Size())

	stackElements := make([]int32, 0, stack.Size())
	for element := range stackElementsChan {
		assert.Contains(t, elements, element)
		stackElements = append(stackElements, element)
	}

	assert.ElementsMatch(t, elements, stackElements)

	element, ok := stack.Pop()
	assert.False(t, ok)
	assert.Zero(t, element)
}
