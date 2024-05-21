package lists

import "errors"

type singleNode[T any] struct {
	value T
	next  *singleNode[T]
}

func newSingleNode[T any](value T) *singleNode[T] {
	return &singleNode[T]{value: value}
}

type SingleLinkedList[T any] struct {
	size int
	head *singleNode[T]
	tail *singleNode[T]
}

func (sl *SingleLinkedList[T]) Size() int {
	return sl.size
}

func (sl *SingleLinkedList[T]) IsEmpty() bool {
	return sl.size > 0
}

var ErrTailIsNotLastNode = errors.New("tail isn't last node")

func (sl *SingleLinkedList[T]) PushTail(value T) error {
	node := newSingleNode(value)
	if sl.size < 1 {
		sl.init(node)
		return nil
	}
	if sl.tail.next != nil {
		return ErrTailIsNotLastNode
	}

	sl.tail.next = node
	sl.tail = sl.tail.next
	sl.size++
	return nil
}

type OnEach[T any] func(value *T) error

func (sl *SingleLinkedList[T]) Map(each OnEach[T]) error {
	for node := sl.head; node != nil; node = node.next {
		if err := each(&node.value); err != nil {
			return err
		}
	}

	return nil
}

func (sl *SingleLinkedList[T]) PushHead(value T) {
	node := newSingleNode(value)
	if sl.size < 1 {
		sl.init(node)
	} else {
		node.next = sl.head
		sl.head = node
		sl.size++
	}
}

var ErrIndexOutOfRange = errors.New("index out of range")

func (sl *SingleLinkedList[T]) Insert(index int, value T) error {
	if index < 0 || sl.size < index {
		return ErrIndexOutOfRange
	}
	if index == 0 {
		sl.PushHead(value)
		return nil
	}
	if index == sl.size {
		return sl.PushTail(value)
	}

	node := sl.head
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	insertedNode := newSingleNode(value)
	insertedNode.next = node.next
	node.next = insertedNode
	sl.size++
	return nil
}

func (sl *SingleLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || sl.size-1 < index {
		return *new(T), ErrIndexOutOfRange
	}
	if index == 0 {
		return sl.head.value, nil
	}
	if index == sl.size-1 {
		return sl.tail.value, nil
	}

	var node *singleNode[T]
	for node = sl.head; index > 0; index-- {
		node = node.next
	}
	return node.value, nil
}

func (sl *SingleLinkedList[T]) Delete(index int) (T, error) {
	if index < 0 || sl.size-1 < index {
		return *new(T), ErrIndexOutOfRange
	}
	var value T
	if index == 0 {
		node := sl.head
		sl.head = node.next
		value = node.value
		node = nil
		sl.size--
		return value, nil
	}

	var node *singleNode[T]
	counter := index
	for node = sl.head; counter-1 > 0; counter-- {
		node = node.next
	}
	if index == sl.size-1 {
		value = node.next.value
		node.next = nil
		sl.tail = node
		sl.size--
		return value, nil
	}

	deletedNode := node.next
	node.next = deletedNode.next
	value = deletedNode.value
	deletedNode = nil
	sl.size--
	return value, nil
}

func (sl *SingleLinkedList[T]) init(node *singleNode[T]) {
	sl.tail = node
	sl.head = node
	sl.size++
}

func NewSingleList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}
