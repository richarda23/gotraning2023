package generics

import (
	"hello/testhelper"
	"testing"
)

type Node1[T comparable] struct {
	Item T
	Next *Node1[T]
}

type LinkedList[T comparable] struct {
	First *Node1[T]
}

func (l *LinkedList[T]) Prepend(item T) {
	n := &Node1[T]{Item: item}
	n.Next = l.First
	l.First = n
}

func (l *LinkedList[T]) Append(item T) {
	n := l.First
	if n == nil {
		l.Prepend((item))
	}
	for n.Next != nil {
		n = n.Next
	}
	toAdd := &Node1[T]{Item: item}
	n.Next = toAdd
}

// Maintains the order of the items in the linked list
func (l *LinkedList[T]) PrependAll(items []T) {
	for i := len(items) - 1; i >= 0; i-- {
		l.Prepend(items[i])
	}
}

// remove first item
func (l *LinkedList[T]) Shift() T {
	var nilN T
	rc := l.First
	if rc == nil {
		return nilN
	}
	l.First = rc.Next
	rc.Next = nil
	return rc.Item
}

// remove last item
func (l *LinkedList[T]) Pop() T {
	var nilN T

	n := l.First
	if n == nil {
		return nilN
	}
	for n.Next != nil && n.Next.Next != nil {
		n = n.Next
	}
	itemToRemove := n.Next
	n.Next = nil
	return itemToRemove.Item
}

func (l *LinkedList[T]) Size() int {
	if l.First == nil {
		return 0
	}
	count := 0
	n := l.First
	for {
		count++
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	return count
}

func TestLinkedList(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Prepend(22)
		list.Prepend(33)
		testhelper.AssertInteger(t, 33, list.First.Item)
	})
	t.Run("Length", func(t *testing.T) {
		list := &LinkedList[int]{}
		testhelper.AssertInteger(t, 0, list.Size())
		list.Prepend(22)
		testhelper.AssertInteger(t, 1, list.Size())
		list.Prepend(33)
		testhelper.AssertInteger(t, 2, list.Size())
	})

	t.Run("Prepend all ", func(t *testing.T) {
		list := &LinkedList[int]{}
		toAdd := []int{1, 2, 3, 4, 5}
		list.PrependAll(toAdd)
		testhelper.AssertInteger(t, 1, list.First.Item)
	})

	t.Run("Append ", func(t *testing.T) {
		list := &LinkedList[int]{}
		toAdd := []int{1, 2, 3, 4}
		list.PrependAll(toAdd)
		list.Append(17)
		testhelper.AssertInteger(t, 5, list.Size())
		testhelper.AssertInteger(t, 17, list.Pop())
		testhelper.AssertInteger(t, 4, list.Size())
	})

	t.Run("Shift", func(t *testing.T) {
		list := &LinkedList[int]{}
		testhelper.AssertNil[int](t, list.Shift())
		testhelper.AssertNil[int](t, list.Pop())
		testhelper.AssertInteger(t, 0, list.Size())
		toAdd := []int{1, 2, 3, 4, 5}
		list.PrependAll(toAdd)
		first := list.Shift()
		testhelper.AssertInteger(t, 1, first)
		testhelper.AssertInteger(t, 4, list.Size())
	})
}
