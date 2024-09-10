package queue

import "sync/atomic"

type element[T any] struct {
	value T
	next  atomic.Pointer[element[T]]
}

type Queue[T any] struct {
	head atomic.Pointer[element[T]]
	tail atomic.Pointer[element[T]]
	size atomic.Uint64
}

func New[T any]() *Queue[T] {
	nothing := &element[T]{}
	queue := &Queue[T]{}
	queue.head.Store(nothing)
	queue.tail.Store(nothing)

	return queue
}

func (queue *Queue[T]) Push(value T) {
	element := &element[T]{value: value}

	for {
		tail := queue.tail.Load()

		if tail.next.CompareAndSwap(nil, element) {
			queue.tail.CompareAndSwap(tail, element)
			queue.size.Add(1)
			return
		} else {
			queue.tail.CompareAndSwap(tail, tail.next.Load())
		}
	}
}

func (queue *Queue[T]) Pop() (T, bool) {
	for {
		head := queue.head.Load()
		tail := queue.tail.Load()
		nextHead := head.next.Load()

		if head == tail {
			if nextHead == nil {
				return *new(T), false
			} else {
				queue.tail.CompareAndSwap(tail, nextHead)
			}
		} else {
			value := nextHead.value
			if queue.head.CompareAndSwap(head, nextHead) {
				queue.size.Add(^uint64(0))
				return value, true
			}
		}
	}
}

func (queue *Queue[T]) Size() uint64 {
	return queue.size.Load()
}
