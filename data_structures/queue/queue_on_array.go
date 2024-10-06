package queue

import (
	"errors"
	"fmt"
)

const defaultQueueCapacity = 20

type Queue[T any] struct {
	fixedSize uint
	data      []T
}

func NewQueue[T any](fixedSize uint) *Queue[T] {
	q := Queue[T]{
		data:      []T{},
		fixedSize: fixedSize,
	}
	if fixedSize == 0 {
		q.fixedSize = uint(defaultQueueCapacity)
	}
	return &q
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() < 1
}

var ErrQueueIsOverflow = errors.New("queue size overflow")

func (q *Queue[T]) Enqueue(value T) error {
	if uint(len(q.data)) >= q.fixedSize {
		return ErrQueueIsOverflow
	}
	q.data = append(q.data, value)
	return nil
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	return q.data[0], true
}

func (q *Queue[T]) Clear() {
	q.data = []T{}
}

func (q *Queue[T]) Print() {
	for _, value := range q.data {
		fmt.Println(value)
	}
}
