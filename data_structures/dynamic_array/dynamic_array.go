package dynamic_array

import "errors"

type DynamicArray[T any] struct {
	length   uint64
	capacity uint64
	arr      []T
}

var ErrIndexOutOfRange = errors.New("index out of range")

func (da *DynamicArray[T]) Remove(n uint64) error {
	if err := da.checkIndex(n); err != nil {
		return err
	}

	copy(da.arr[n:], da.arr[n+1:da.length])
	da.arr[da.length-1] = *new(T)
	da.length--
	return nil
}

func (da *DynamicArray[T]) All() []T {
	arr := make([]T, da.length)
	for i := uint64(0); i < da.length; i++ {
		arr[i] = da.arr[i]
	}

	return arr
}

func (da *DynamicArray[T]) Get(n uint64) (T, error) {
	if err := da.checkIndex(n); err != nil {
		return *new(T), err
	}

	return da.arr[n], nil
}

func (da *DynamicArray[T]) Put(n uint64, element T) error {
	if err := da.checkIndex(n); err != nil {
		return err
	}

	da.arr[n] = element
	return nil
}

func (da *DynamicArray[T]) Add(element T) *DynamicArray[T] {
	if da.length >= da.capacity {
		da.expand()
	}

	da.arr[da.length] = element
	da.length++
	return da
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.Length() > 0
}

func (da *DynamicArray[T]) Capacity() uint64 {
	return da.capacity
}

func (da *DynamicArray[T]) Length() uint64 {
	return da.length
}

func (da *DynamicArray[T]) expand() {
	da.capacity = da.capacity << 1
	expanded := make([]T, da.capacity)
	copy(expanded, da.arr)
	da.arr = expanded
}

func (da *DynamicArray[T]) checkIndex(n uint64) error {
	if n > da.length-1 {
		return ErrIndexOutOfRange
	}
	return nil
}

func NewDynamicArray[T any](capacity uint64) *DynamicArray[T] {
	if capacity < 1 {
		panic("capacity less than 1")
	}
	da := DynamicArray[T]{capacity: capacity}
	da.arr = make([]T, capacity)

	return &da
}
