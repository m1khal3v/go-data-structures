package singlylinkedlist

import "sync"

type SortFunction[T comparable] func(first, second T) bool

func mergeSort[T comparable](headRef **element[T], sort SortFunction[T]) {
	head := *headRef

	if head == nil || head.next == nil {
		return
	}

	first, second := halfDivide(head)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		mergeSort(&first, sort)
	}()
	go func() {
		defer wg.Done()
		mergeSort(&second, sort)
	}()
	wg.Wait()

	*headRef = sortAndMerge(first, second, sort)
}

func sortAndMerge[T comparable](first, second *element[T], sort SortFunction[T]) *element[T] {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}

	var result *element[T]
	if sort(first.value, second.value) {
		result = first
		result.next = sortAndMerge(first.next, second, sort)
	} else {
		result = second
		result.next = sortAndMerge(first, second.next, sort)
	}

	return result
}

func halfDivide[T comparable](head *element[T]) (*element[T], *element[T]) {
	slow, fast := head, head.next

	for fast != nil {
		fast = fast.next

		if fast != nil {
			slow = slow.next
			fast = fast.next
		}
	}

	first, second := head, slow.next
	slow.next = nil

	return first, second
}
