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

type Hasher[K comparable] interface {
	makeHash(key K) (uint64, error)
}

type hashMaker[K comparable] struct {
	capacity uint64
}

func (h *hashMaker[K]) makeHash(key K) (uint64, error) {
	hs := fnv.New64()
	if _, err := hs.Write([]byte(fmt.Sprintf("%v", key))); err != nil {
		return 0, err
	}
	hashVal := hs.Sum64()
	hashVal = (h.capacity - 1) & (hashVal ^ (hashVal >> 16))
	return hashVal, nil
}

type HashTable[K comparable, T any] struct {
	size     uint64
	capacity uint64
	table    []*node[K, T]
	hasher   Hasher[K]
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

func (t *HashTable[K, T]) IsEmpty() bool {
	return t.size < 1
}

var ErrElementIsEmptyByKey = errors.New("cannot find element by key")

func (t *HashTable[K, T]) Get(key K) (T, error) {
	hashIndex, err := t.hasher.makeHash(key)
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
	hashIndex, err := t.hasher.makeHash(key)
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
	hashIndex, err := t.hasher.makeHash(key)
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

type HashTableFunc[T any] func(val T)

func (t *HashTable[K, T]) Foreach(tFunc HashTableFunc[T]) {
	if t.IsEmpty() {
		return
	}
	for _, el := range t.table {
		if el != nil {
			for v := el; v != nil; v = v.next {
				tFunc(v.value)
			}
		}
	}
}

func (t *HashTable[K, T]) Clear() {
	t.size = 0
	for i, el := range t.table {
		if el != nil {
			t.table[i] = nil
		}
	}
	fmt.Println("clear")
}

func (t *HashTable[K, T]) Remove(key K) error {
	hashIndex, err := t.hasher.makeHash(key)
	if err != nil {
		return err
	}
	if t.table[hashIndex] == nil {
		return fmt.Errorf("%s %v", ErrElementIsEmptyByKey, key)
	}
	current := t.table[hashIndex]
	if current.key == key {
		t.table[hashIndex] = current.next
		current = nil
		t.size--
		return nil
	}
	for el := current; el.next != nil; el = el.next {
		rmEl := el.next
		if rmEl.key == key {
			el.next = rmEl.next
			rmEl = nil
			t.size--
			return nil
		}
	}
	return fmt.Errorf("%s %v", ErrElementIsEmptyByKey, key)
}

func (t *HashTable[K, T]) Keys() []K {
	var keys []K
	if t.IsEmpty() {
		return keys
	}
	for _, v := range t.table {
		if v != nil {
			for el := v; el != nil; el = el.next {
				keys = append(keys, el.key)
			}
		}
	}
	return keys
}

func (t *HashTable[K, T]) Values() []T {
	var values []T
	if t.IsEmpty() {
		return values
	}
	for _, v := range t.table {
		if v != nil {
			for el := v; el != nil; el = el.next {
				values = append(values, el.value)
			}
		}
	}
	return values
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
		hasher:   &hashMaker[K]{capacity: defaultCapacity},
	}
}

var ErrCapacityIsEmpty = errors.New("size cannot be 0")

func NewHashTableWithCapacity[K comparable, T any](capacity uint64) (*HashTable[K, T], error) {
	if capacity < 1 {
		return new(HashTable[K, T]), ErrCapacityIsEmpty
	}
	t := make([]*node[K, T], capacity)
	ht := &HashTable[K, T]{
		capacity: capacity,
		table:    t,
		hasher:   &hashMaker[K]{capacity: capacity},
	}

	return ht, nil
}
