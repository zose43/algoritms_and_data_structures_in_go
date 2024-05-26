package lists

import "errors"

type DoubleNode[T any] struct {
	next     *DoubleNode[T]
	previous *DoubleNode[T]
	value    T
}

func NewDoubleNode[T any](value T) *DoubleNode[T] {
	return &DoubleNode[T]{value: value}
}

type DoubleLinkedList[T any] struct {
	size int
	tail *DoubleNode[T]
	head *DoubleNode[T]
}

func (dl *DoubleLinkedList[T]) Length() int {
	return dl.size
}

func (dl *DoubleLinkedList[T]) IsEmpty() bool {
	return dl.size > 0
}

var (
	ErrHeadIsUndefined = errors.New("head node is nil")
	ErrHeadIsNotHead   = errors.New("head previous isn't nil")
)

func (dl *DoubleLinkedList[T]) PushHead(value T) error {
	node := NewDoubleNode(value)
	if dl.size < 1 {
		dl.head = node
		dl.tail = node
		dl.size++
		return nil
	}
	if dl.head == nil {
		return ErrHeadIsUndefined
	}
	if dl.head.previous != nil {
		return ErrHeadIsNotHead
	}

	dl.head.previous = node
	node.next = dl.head
	dl.head = node
	dl.size++
	return nil
}

var ErrTailIsUndefined = errors.New("tail node is nil")

func (dl *DoubleLinkedList[T]) PushTail(value T) error {
	node := NewDoubleNode(value)
	if dl.size < 1 {
		dl.tail = node
		dl.head = node
		dl.size++
		return nil
	}
	if dl.tail == nil {
		return ErrTailIsUndefined
	}
	if dl.tail.next != nil {
		return ErrTailIsNotLastNode
	}

	dl.tail.next = node
	node.previous = dl.tail
	dl.tail = node
	dl.size++
	return nil
}

func (dl *DoubleLinkedList[T]) Insert(index int, value T) error {
	if err := dl.checkRange(index); err != nil {
		return err
	}
	if index == 0 {
		return dl.PushHead(value)
	}
	if index == dl.size-1 {
		return dl.PushTail(value)
	}

	var node *DoubleNode[T]
	for node = dl.head; index > 0; index-- {
		node = node.next
	}
	inserted := NewDoubleNode(value)
	inserted.previous = node.previous
	inserted.next = node
	node.previous.next = inserted
	node.previous = inserted
	dl.size++
	return nil
}

func (dl *DoubleLinkedList[T]) Get(index int) (T, error) {
	if err := dl.checkRange(index); err != nil {
		return *new(T), err
	}
	if index == 0 {
		return dl.head.value, nil
	}
	if index == dl.size-1 {
		return dl.tail.value, nil
	}

	var node *DoubleNode[T]
	for node = dl.head; index > 0; index-- {
		node = node.next
	}
	return node.value, nil
}

func (dl *DoubleLinkedList[T]) Remove(index int) (T, error) {
	if err := dl.checkRange(index); err != nil {
		return *new(T), err
	}
	var value T
	if index == 0 {
		value = dl.head.value
		node := dl.head
		dl.head = node.next
		dl.head.previous = nil
		node = nil
		dl.size--
		return value, nil
	}
	if index == dl.size-1 {
		value = dl.tail.value
		node := dl.tail
		dl.tail = node.previous
		dl.tail.next = nil
		node = nil
		dl.size--
		return value, nil
	}

	var deletedNode *DoubleNode[T]
	for deletedNode = dl.head; index > 0; index-- {
		deletedNode = deletedNode.next
	}
	deletedNode.previous.next = deletedNode.next
	deletedNode.next.previous = deletedNode.previous
	value = deletedNode.value
	deletedNode = nil
	dl.size--
	return value, nil
}

func (dl *DoubleLinkedList[T]) checkRange(index int) error {
	if index < 0 || index > dl.size-1 {
		return ErrIndexOutOfRange
	}
	return nil
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}
