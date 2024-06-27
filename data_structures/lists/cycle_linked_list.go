package lists

import "errors"

type cycledLinkedNode[T any] struct {
	value T
	prev  *cycledLinkedNode[T]
	next  *cycledLinkedNode[T]
}

func newCycledLinkedNode[T any](value T) *cycledLinkedNode[T] {
	return &cycledLinkedNode[T]{value: value}
}

type CycleLinkedList[T any] struct {
	head *cycledLinkedNode[T]
	size int
}

func NewCycleLinkedList[T any]() *CycleLinkedList[T] {
	return &CycleLinkedList[T]{}
}

func (cl *CycleLinkedList[T]) isEmpty() bool {
	return cl.size < 1
}

func (cl *CycleLinkedList[T]) Size() int {
	return cl.size
}

func (cl *CycleLinkedList[T]) Add(element T) {
	node := newCycledLinkedNode(element)
	if cl.isEmpty() {
		cl.head = node
		cl.head.prev = node
		cl.head.next = node
		cl.size++
		return
	}
	node.next = cl.head
	node.prev = cl.head.prev
	cl.head.prev.next = node
	cl.head.prev = node
	cl.head = node
	cl.size++
}

var ErrListIsEmpty = errors.New("list is empty")

func (cl *CycleLinkedList[T]) Value() (T, error) {
	if cl.isEmpty() {
		return *new(T), ErrListIsEmpty
	}
	return cl.head.value, nil
}

type CycleLinkFunc[T any] func(T)

func (cl *CycleLinkedList[T]) ForEach(clFunc CycleLinkFunc[T]) {
	if cl.isEmpty() {
		return
	}
	node := cl.head
	clFunc(node.value)
	for i := cl.size; i > 1; i-- {
		node = node.next
		clFunc(node.value)
	}
}

func (cl *CycleLinkedList[T]) ReverseForEach(clFunc CycleLinkFunc[T]) {
	if cl.isEmpty() {
		return
	}
	node := cl.head.prev
	clFunc(node.value)
	for i := cl.size; i > 1; i-- {
		node = node.prev
		clFunc(node.value)
	}
}

func (cl *CycleLinkedList[T]) Rotate(delta int) {
	if cl.isEmpty() {
		return
	}
	delta %= cl.size
	if delta == 0 {
		return
	}
	if delta < 0 {
		for ; delta < 0; delta++ {
			cl.head = cl.head.prev
		}
	} else {
		for i := 0; i < delta; i++ {
			cl.head = cl.head.next
		}
	}
}

func (cl *CycleLinkedList[T]) Remove() bool {
	if cl.isEmpty() {
		return false
	}
	currentNode := cl.head
	if cl.size == 1 {
		cl.head = nil
		cl.size--
		return true
	}

	cl.head = currentNode.next
	cl.head.prev.next = currentNode.next
	cl.head.prev = currentNode.prev
	cl.size--
	return true
}

func (cl *CycleLinkedList[T]) RemoveAll() {
	for cl.Remove() {
	}
}
