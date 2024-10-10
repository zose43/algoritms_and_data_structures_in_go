package queue

import (
	"container/list"
	"fmt"
)

type listQueueElement[T any] struct {
	value T
}

type BasedOnListQueue[T any] struct {
	fixedSize uint
	list      *list.List
}

func NewBasedOnListQueue[T any](fixedSize uint) *BasedOnListQueue[T] {
	ql := BasedOnListQueue[T]{
		fixedSize: defaultQueueCapacity,
		list:      list.New(),
	}
	if fixedSize > 0 {
		ql.fixedSize = fixedSize
	}
	return &ql
}

func (ql *BasedOnListQueue[T]) Size() int {
	return ql.list.Len()
}

func (ql *BasedOnListQueue[T]) IsEmpty() bool {
	return ql.list.Len() < 1
}

func (ql *BasedOnListQueue[T]) Enqueue(value T) error {
	if uint(ql.list.Len()) >= ql.fixedSize {
		return ErrQueueIsOverflow
	}
	el := listQueueElement[T]{value: value}
	ql.list.PushBack(&el)
	return nil
}

func (ql *BasedOnListQueue[T]) Dequeue() (T, bool) {
	if ql.IsEmpty() {
		return *new(T), false
	}
	el := ql.list.Front()
	if el == nil {
		return *new(T), false
	}
	ql.list.Remove(el)
	return el.Value.(*listQueueElement[T]).value, true
}

func (ql *BasedOnListQueue[T]) Peek() (T, bool) {
	if ql.IsEmpty() {
		return *new(T), false
	}
	el := ql.list.Front()
	if el == nil {
		return *new(T), false
	}
	return el.Value.(*listQueueElement[T]).value, true
}

func (ql *BasedOnListQueue[T]) Print() {
	for el := ql.list.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(*listQueueElement[T]).value)
	}
}
