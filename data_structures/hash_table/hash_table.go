package hash_table

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const defaultCapacity uint64 = 10 << 10

type node[K comparable, T any] struct {
	key   K
	value T
	next  *node[K, T]
}

type HashTable[K comparable, T any] struct {
	size     uint64
	capacity uint64
	table    []*node[K, T]
}

func (t *HashTable[K, T]) newNode(key K, value T) *node[K, T] {
	return &node[K, T]{key: key, value: value}
}

func (t *HashTable[K, T]) Capacity() uint64 {
	return t.capacity
}

func (t *HashTable[K, T]) Size() uint64 {
	return t.size
}

func (t *HashTable[K, T]) makeHash(key K) (uint64, error) {
	h := fnv.New64()
	if _, err := h.Write([]byte(fmt.Sprintf("%v", key))); err != nil {
		return 0, err
	}
	hashVal := h.Sum64()
	hashVal = (t.capacity - 1) & (hashVal ^ (hashVal >> 16))
	return hashVal, nil
}

var ErrElementIsEmptyByKey = errors.New("cannot find element by key")

func (t *HashTable[K, T]) Get(key K) (T, error) {
	hashIndex, err := t.makeHash(key)
	if err != nil {
		return *new(T), errors.Join(err, ErrElementIsEmptyByKey)
	}
	if t.table[hashIndex] == nil {
		return *new(T), fmt.Errorf("%v %v", ErrElementIsEmptyByKey, key)
	}
	for el := t.table[hashIndex]; el != nil; el = el.next {
		if el.key == key {
			return el.value, nil
		}
	}
	return *new(T), fmt.Errorf("%v %v", ErrElementIsEmptyByKey, key)
}

func (t *HashTable[K, T]) Contains(key K) (bool, error) {
	hashIndex, err := t.makeHash(key)
	if err != nil {
		return false, errors.Join(err, ErrElementIsEmptyByKey)
	}
	if t.table[hashIndex] == nil {
		return false, nil
	}
	for el := t.table[hashIndex]; el != nil; el = el.next {
		if el.key == key {
			return true, nil
		}
	}
	return false, nil
}

func (t *HashTable[K, T]) Put(key K, val T) error {
	hashIndex, err := t.makeHash(key)
	if err != nil {
		return err
	}
	if t.table[hashIndex] == nil {
		t.table[hashIndex] = t.newNode(key, val)
		t.size++
		return nil
	}
	for el := t.table[hashIndex]; el != nil; el = el.next {
		if el.key == key {
			el.value = val
			return nil
		}
	}
	t.resolvePutCollision(key, val, hashIndex)
	t.size++
	return nil
}

func (t *HashTable[K, T]) resolvePutCollision(key K, val T, index uint64) {
	// add element to head
	t.table[index] = &node[K, T]{
		key:   key,
		value: val,
		next:  t.table[index],
	}
}

func NewHashTable[K comparable, T any]() *HashTable[K, T] {
	t := make([]*node[K, T], defaultCapacity)
	return &HashTable[K, T]{
		capacity: defaultCapacity,
		table:    t,
	}
}

var ErrCapacityIsEmpty = errors.New("capacity cannot be 0")

func NewHashTableWithCapacity[K comparable, T any](capacity uint64) (*HashTable[K, T], error) {
	if capacity < 1 {
		return new(HashTable[K, T]), ErrCapacityIsEmpty
	}
	t := make([]*node[K, T], capacity)
	ht := &HashTable[K, T]{
		capacity: capacity,
		table:    t,
	}

	return ht, nil
}
