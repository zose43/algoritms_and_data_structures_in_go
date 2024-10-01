package stack

import (
	"algoritms_and_structures/data_structures/lists"
	"fmt"
)

type LinkedListStack[T any] struct {
	fixedSize uint
	list      *lists.SingleLinkedList[T]
}

func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *LinkedListStack[T]) Size() int {
	return s.list.Size()
}

func (s *LinkedListStack[T]) Push(value T) error {
	if uint(s.Size()) >= s.fixedSize {
		return ErrStackOverflow
	}
	s.list.PushHead(value)
	return nil
}

func (s *LinkedListStack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}
	v, err := s.list.Delete(0)
	if err != nil {
		return *new(T), false
	}
	return v, true
}

func (s *LinkedListStack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}
	v, err := s.list.Get(0)
	if err != nil {
		return *new(T), false
	}
	return v, true
}

func (s *LinkedListStack[T]) Print() {
	_ = s.list.Map(func(value *T) error {
		fmt.Println(*value)
		return nil
	})
}

func NewLinkedListStack[T any](size uint) (*LinkedListStack[T], error) {
	if size < 1 {
		return nil, ErrStackInit
	}
	return &LinkedListStack[T]{
		fixedSize: size,
		list:      lists.NewSingleList[T](),
	}, nil
}
