package queue

type element[T any] struct {
	value T
	next  *element[T]
}

type Queue[T any] struct {
	head *element[T]
	tail *element[T]
	size uint64
}

func New[T any]() *Queue[T] {
	nothing := &element[T]{}

	return &Queue[T]{
		head: nothing,
		tail: nothing,
	}
}

func (queue *Queue[T]) Push(value T) {
	queue.tail.next = &element[T]{value: value}
	queue.tail = queue.tail.next
	queue.size++
}

func (queue *Queue[T]) Pop() (T, bool) {
	if queue.size == 0 {
		return *new(T), false
	}

	value := queue.head.next.value
	queue.head = queue.head.next
	queue.size--

	return value, true
}

func (queue *Queue[T]) Size() uint64 {
	return queue.size
}
