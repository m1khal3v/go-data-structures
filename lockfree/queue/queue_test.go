package queue

import (
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"sync"
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

func TestQueueWithConcurrency(t *testing.T) {
	queue := New[int32]()
	assert.EqualValues(t, 0, queue.Size())

	count := rand.Uint32N(50) + 50
	elements := make([]int32, 0, count)

	for i := uint32(1); i <= count; i++ {
		element := rand.Int32N(100) - 50
		elements = append(elements, element)
	}

	var wg sync.WaitGroup
	queueElementsChan := make(chan int32, count)
	for _, element := range elements {
		wg.Add(2)
		go func() {
			defer wg.Done()
			queue.Push(element)
		}()
		go func() {
			defer wg.Done()
			for {
				element, ok := queue.Pop()
				if ok {
					queueElementsChan <- element
					return
				}
			}
		}()
	}
	wg.Wait()
	close(queueElementsChan)
	assert.EqualValues(t, 0, queue.Size())

	queueElements := make([]int32, 0, queue.Size())
	for element := range queueElementsChan {
		assert.Contains(t, elements, element)
		queueElements = append(queueElements, element)
	}

	assert.ElementsMatch(t, elements, queueElements)

	element, ok := queue.Pop()
	assert.False(t, ok)
	assert.Zero(t, element)
}
