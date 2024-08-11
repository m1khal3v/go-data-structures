package stack

type element[T any] struct {
	value T
	next  *element[T]
}

type Stack[T any] struct {
	head *element[T]
	size uint64
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(value T) {
	stack.head = &element[T]{value: value, next: stack.head}
	stack.size++
}

func (stack *Stack[T]) Pop() (T, bool) {
	if stack.head == nil {
		return *new(T), false
	}

	value := stack.head.value
	stack.head = stack.head.next
	stack.size--

	return value, true
}

func (stack *Stack[T]) Size() uint64 {
	return stack.size
}
