package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	fixedSize uint
	elements  []T
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) > 1
}

var ErrStackOverflow = errors.New("stack overflow")

func (s *Stack[T]) Push(value T) error {
	if uint(s.Size()) >= s.fixedSize {
		return ErrStackOverflow
	}
	s.elements = append(s.elements, value)
	return nil
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}
	el := s.elements[0]
	s.elements = s.elements[1:]
	return el, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}
	return s.elements[0], true
}

func (s *Stack[T]) Clear() {
	s.elements = []T{}
}

func (s *Stack[T]) Print() {
	for _, el := range s.elements {
		fmt.Printf("%v\n", el)
	}
}

var ErrStackInit = errors.New("stack init error, size must be greater than zero")

func NewStack[T any](fixedSize uint) (*Stack[T], error) {
	if fixedSize < 1 {
		return nil, ErrStackInit
	}
	return &Stack[T]{fixedSize: fixedSize}, nil
}
