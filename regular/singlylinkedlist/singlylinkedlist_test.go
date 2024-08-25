package singlylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglyLinkedList_LPush(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.LPush(0)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_LPop(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	item, ok := list.LPop()
	assert.True(t, ok)
	assert.Equal(t, 1, item)
	assert.Equal(t, []int{2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_LPopEmpty(t *testing.T) {
	list := New[uint8]()
	_, ok := list.LPop()
	assert.False(t, ok)
}

func TestSinglyLinkedList_Prepend(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Prepend(0, -1, -2)
	assert.Equal(t, []int{-2, -1, 0, 1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_RPush(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.RPush(6)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, list.Values())
}

func TestSinglyLinkedList_RPop(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	item, ok := list.RPop()
	assert.True(t, ok)
	assert.Equal(t, 5, item)
	assert.Equal(t, []int{1, 2, 3, 4}, list.Values())
}

func TestSinglyLinkedList_RPopEmpty(t *testing.T) {
	list := New[uint8]()
	_, ok := list.RPop()
	assert.False(t, ok)
}

func TestSinglyLinkedList_Append(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Append(6, 7, 8)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, list.Values())
}

func TestSinglyLinkedList_Push(t *testing.T) {
	list := New(1, 2, 4, 5)
	list.Push(2, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_Pop(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	item, ok := list.Pop(2)
	assert.True(t, ok)
	assert.Equal(t, 3, item)
	assert.Equal(t, []int{1, 2, 4, 5}, list.Values())
}

func TestSinglyLinkedList_PopEmpty(t *testing.T) {
	list := New[uint8]()
	_, ok := list.Pop(1)
	assert.False(t, ok)
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	list := New(-2, -1, 0, 4, 5, 6)
	list.Insert(3, 1, 2, 3)
	assert.Equal(t, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6}, list.Values())
}

func TestSinglyLinkedList_Set(t *testing.T) {
	list := New(1, 2, 0, 4, 5)
	list.Set(2, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_Get(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	item, ok := list.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 3, item)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	ok := list.Remove(2)
	assert.True(t, ok)
	assert.Equal(t, []int{1, 2, 4, 5}, list.Values())
}

func TestSinglyLinkedList_Swap(t *testing.T) {
	list := New(1, 2, 5, 4, 3)
	ok := list.Swap(2, 4)
	assert.True(t, ok)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_IndexOf(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	index, ok := list.IndexOf(3)
	assert.True(t, ok)
	assert.EqualValues(t, 2, index)
}

func TestSinglyLinkedList_Contains(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	ok := list.Contains(2, 3, 4)
	assert.True(t, ok)
}

func TestSinglyLinkedList_NotContains(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	ok := list.Contains(0, 6, 1)
	assert.False(t, ok)
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	assert.EqualValues(t, 5, list.Size())
	list.Clear()
	assert.Empty(t, list.Values())
	assert.EqualValues(t, 0, list.Size())
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Reverse()
	assert.Equal(t, []int{5, 4, 3, 2, 1}, list.Values())
}

func TestSinglyLinkedList_Sort(t *testing.T) {
	list := New(1, 3, 5, 4, 2)
	list.Sort(func(first, second int) bool {
		return first <= second
	})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Values())
}

func TestSinglyLinkedList_ReverseSort(t *testing.T) {
	list := New(1, 3, 5, 4, 2)
	list.Sort(func(first, second int) bool {
		return first >= second
	})
	assert.Equal(t, []int{5, 4, 3, 2, 1}, list.Values())
}
