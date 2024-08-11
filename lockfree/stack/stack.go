package stack

import "sync/atomic"

type element[T any] struct {
	value T
	next  *element[T]
}

type Stack[T any] struct {
	head atomic.Pointer[element[T]]
	size atomic.Uint64
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(value T) {
	element := &element[T]{value: value}

	for {
		head := stack.head.Load()
		element.next = head

		if stack.head.CompareAndSwap(head, element) {
			stack.size.Add(1)
			return
		}
	}
}

func (stack *Stack[T]) Pop() (T, bool) {
	for {
		head := stack.head.Load()
		if head == nil {
			return *new(T), false
		}

		if stack.head.CompareAndSwap(head, head.next) {
			stack.size.Add(^uint64(0))
			return head.value, true
		}
	}
}

func (stack *Stack[T]) Size() uint64 {
	return stack.size.Load()
}
