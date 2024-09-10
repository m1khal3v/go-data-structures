package doublylinkedlist

type SortFunction[T comparable] func(first, second T) bool

func mergeSort[T comparable](headRef **element[T], sort SortFunction[T]) {
	head := *headRef

	if head == nil || head.next == nil {
		return
	}

	first, second := halfDivide(head)

	mergeSort(&first, sort)
	mergeSort(&second, sort)

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

	result.next.previous = result
	result.previous = nil

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
