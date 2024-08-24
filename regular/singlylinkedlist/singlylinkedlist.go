package singlylinkedlist

import "slices"

type element[T comparable] struct {
	value T
	next  *element[T]
}

type SinglyLinkedList[T comparable] struct {
	head *element[T]
	tail *element[T]
	size uint64
}

func NewSinglyLinkedList[T comparable](values ...T) *SinglyLinkedList[T] {
	list := &SinglyLinkedList[T]{}
	list.Append(values...)

	return list
}

func (list *SinglyLinkedList[T]) LPush(value T) {
	newElement := &element[T]{
		value: value,
		next:  list.head,
	}

	list.head = newElement

	if list.size == 0 {
		list.tail = newElement
	}

	list.size++
}

func (list *SinglyLinkedList[T]) LPop() (T, bool) {
	if list.size == 0 {
		return *new(T), false
	}

	value := list.head.value
	if list.size == 1 {
		list.Clear()

		return value, true
	}

	list.head = list.head.next
	list.size--

	return value, true
}

func (list *SinglyLinkedList[T]) Prepend(values ...T) {
	for _, value := range values {
		list.LPush(value)
	}
}

func (list *SinglyLinkedList[T]) RPush(value T) {
	newElement := &element[T]{
		value: value,
	}

	if list.size == 0 {
		list.head = newElement
	} else {
		list.tail.next = newElement
	}

	list.tail = newElement
	list.size++
}

func (list *SinglyLinkedList[T]) RPop() (T, bool) {
	if list.size == 0 {
		return *new(T), false
	}

	value := list.tail.value
	if list.size == 1 {
		list.Clear()

		return value, true
	}

	var tail *element[T]
	for tail = list.head.next; tail != list.tail; tail = tail.next {
	}

	tail.next = nil
	list.tail = tail
	list.size--

	return value, true
}

func (list *SinglyLinkedList[T]) Append(values ...T) {
	for _, value := range values {
		list.RPush(value)
	}
}

func (list *SinglyLinkedList[T]) Push(index uint64, value T) bool {
	if index > list.size {
		return false
	}

	switch index {
	case 0:
		list.LPush(value)
	case list.size:
		list.RPush(value)
	default:
		var previous *element[T]
		var currentIndex uint64
		previousIndex := index - 1
		for previous, currentIndex = list.head, uint64(0); currentIndex != previousIndex; previous, currentIndex = previous.next, currentIndex+1 {
		}

		newElement := &element[T]{
			value: value,
			next:  previous.next,
		}
		previous.next = newElement

		list.size++
	}

	return true
}

func (list *SinglyLinkedList[T]) Pop(index uint64) (T, bool) {
	if index >= list.size {
		return *new(T), false
	}

	switch index {
	case 0:
		return list.LPop()
	case list.size - 1:
		return list.RPop()
	default:
		var previous *element[T]
		var currentIndex uint64
		previousIndex := index - 1
		for previous, currentIndex = list.head, uint64(0); currentIndex != previousIndex; previous, currentIndex = previous.next, currentIndex+1 {
		}

		element := previous.next
		previous.next = element.next

		list.size--

		return element.value, true
	}
}

func (list *SinglyLinkedList[T]) Insert(index uint64, values ...T) bool {
	if index > list.size {
		return false
	}

	switch index {
	case 0:
		slices.Reverse(values)
		for _, value := range values {
			list.LPush(value)
		}
	case list.size:
		for _, value := range values {
			list.RPush(value)
		}
	default:
		var previous *element[T]
		var currentIndex uint64
		previousIndex := index - 1
		for previous, currentIndex = list.head, uint64(0); currentIndex != previousIndex; previous, currentIndex = previous.next, currentIndex+1 {
		}

		for _, value := range values {
			newElement := &element[T]{
				value: value,
				next:  previous.next,
			}
			previous.next = newElement
			previous = newElement

			list.size++
		}

	}

	return true
}

func (list *SinglyLinkedList[T]) Set(index uint64, value T) bool {
	if index >= list.size {
		return false
	}

	switch index {
	case 0:
		list.head = &element[T]{
			value: value,
			next:  list.head.next,
		}
	default:
		var previous *element[T]
		var currentIndex uint64
		previousIndex := index - 1
		for previous, currentIndex = list.head, uint64(0); currentIndex != previousIndex; previous, currentIndex = previous.next, currentIndex+1 {
		}

		newElement := &element[T]{
			value: value,
			next:  previous.next.next,
		}
		previous.next = newElement
	}

	return true
}

func (list *SinglyLinkedList[T]) Get(index uint64) (T, bool) {
	if index >= list.size {
		return *new(T), false
	}

	switch index {
	case 0:
		return list.head.value, true
	case list.size - 1:
		return list.tail.value, true
	default:
		var current *element[T]
		var currentIndex uint64

		for current, currentIndex = list.head.next, uint64(1); currentIndex != index; current, currentIndex = current.next, currentIndex+1 {
		}
		return current.value, true
	}
}

func (list *SinglyLinkedList[T]) Remove(index uint64) bool {
	if index >= list.size {
		return false
	}

	if list.size == 1 {
		list.Clear()
		return true
	}

	switch index {
	case 0:
		list.head = list.head.next
	default:
		var previous *element[T]
		var currentIndex uint64
		previousIndex := index - 1
		for previous, currentIndex = list.head, uint64(0); currentIndex != previousIndex; previous, currentIndex = previous.next, currentIndex+1 {
		}

		previous.next = previous.next.next
	}

	list.size--

	return true
}

func (list *SinglyLinkedList[T]) Swap(index1, index2 uint64) bool {
	if index1 == index2 || index1 >= list.size || index2 >= list.size {
		return false
	}

	maxIndex := max(index1, index2)
	var element1, element2 *element[T]

	for current, currentIndex := list.head, uint64(0); currentIndex <= maxIndex; current, currentIndex = current.next, currentIndex+1 {
		switch currentIndex {
		case index1:
			element1 = current
		case index2:
			element2 = current
		}
	}

	element1.value, element2.value = element2.value, element1.value

	return true
}

func (list *SinglyLinkedList[T]) IndexOf(value T) (uint64, bool) {
	if list.size == 0 {
		return 0, false
	}

	for current, currentIndex := list.head, uint64(0); current != nil; current, currentIndex = current.next, currentIndex+1 {
		if current.value == value {
			return currentIndex, true
		}
	}

	return 0, false
}

func (list *SinglyLinkedList[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true
	}

	if list.size == 0 {
		return false
	}

	for current := list.head; current != nil; current = current.next {
		for index, value := range values {
			if current.value == value {
				values[index] = values[len(values)-1]
				values = values[:len(values)-1]
				break
			}
		}

		if len(values) == 0 {
			return true
		}
	}

	return false
}

func (list *SinglyLinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *SinglyLinkedList[T]) Reverse() {
	switch list.size {
	case 0, 1:
		return
	case 2:
		list.tail.next = list.head
		list.head.next = nil

		list.head, list.tail = list.tail, list.head
	default:
		list.tail = list.head

		var head *element[T]
		for current := list.head; current != nil; {
			temp := current
			current = current.next
			temp.next = head
			head = temp
		}

		list.head = head
	}
}

func (list *SinglyLinkedList[T]) Values() []T {
	values := make([]T, list.size)
	for index, element := 0, list.head; element != nil; index, element = index+1, element.next {
		values[index] = element.value
	}

	return values
}

func (list *SinglyLinkedList[T]) Size() uint64 {
	return list.size
}

func (list *SinglyLinkedList[T]) Sort(sort SortFunction[T]) {
	mergeSort(&list.head, sort)
}
