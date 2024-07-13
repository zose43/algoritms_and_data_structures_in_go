package hash_table

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type people struct {
	name string
	age  int
}

func TestHashTable_Get(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name    string
		th      *HashTable[K, T]
		args    K
		want    T
		wantErr bool
	}
	tests := []testCase[int, *people]{
		{
			name: "got error when empty",
			th: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return t
			}(),
			args:    1,
			want:    (*people)(nil),
			wantErr: true,
		},
		{
			name: "got error when hash failed",
			th: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				h := MockHasher[int]{}
				h.On("makeHash", mock.AnythingOfType("int")).
					Return(uint64(0), errors.New("test err"))
				t.hasher = &h
				return t
			}(),
			args:    1,
			want:    (*people)(nil),
			wantErr: true,
		},
		{
			name: "got error when cannot find by key",
			th: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				t.table[132] = &node[int, *people]{value: &people{name: "test"}}
				return t
			}(),
			args:    1,
			want:    (*people)(nil),
			wantErr: true,
		},
		{
			name: "got single element",
			th: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				t.table[132] = &node[int, *people]{value: &people{name: "test"}}
				return t
			}(),
			args: 0,
			want: &people{name: "test"},
		},
		{
			name: "got throw multiple elements",
			th: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				t.table[132] = &node[int, *people]{
					value: &people{name: "test"},
					key:   0,
					next: &node[int, *people]{
						key:   2,
						value: &people{name: "test1"},
					},
				}
				return t
			}(),
			args: 2,
			want: &people{name: "test1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.th.Get(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, got, tt.want, "Get(%v)", tt.args)
		})
	}
}

func TestHashTable_makeHash(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name    string
		ht      *HashTable[K, T]
		args    K
		want    uint64
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[int, *people]{
		{
			name:    "got value till size",
			want:    uint64(132),
			wantErr: assert.NoError,
			ht: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return t
			}(),
		},
		{
			name:    "got value more than size",
			args:    500,
			want:    uint64(149),
			wantErr: assert.NoError,
			ht: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return t
			}(),
		},
		{
			name:    "got value == size",
			args:    150,
			want:    uint64(5),
			wantErr: assert.NoError,
			ht: func() *HashTable[int, *people] {
				t, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return t
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ht.hasher.makeHash(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("makeHash(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "makeHash(%v)", tt.args)
			assert.Less(t, got, tt.ht.capacity, "makeHash(%v)", tt.args)
		})
	}
}

func TestHashTable_Put(t *testing.T) {
	type args[K comparable, T any] struct {
		key K
		val T
	}
	type want[T any] struct {
		wantErr assert.ErrorAssertionFunc
		size    uint64
		val     T
	}
	type testCase[K comparable, T any] struct {
		name  string
		ht    *HashTable[K, T]
		args  args[K, T]
		want  want[T]
		check bool
	}
	tests := []testCase[int, *people]{
		{
			name: "got error by hasher",
			args: args[int, *people]{key: 2, val: &people{name: "test", age: 1}},
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				h := MockHasher[int]{}
				h.On("makeHash", mock.AnythingOfType("int")).
					Return(uint64(0), errors.New("test err"))
				mt.hasher = &h
				return mt
			}(),
			check: false,
			want: want[*people]{
				wantErr: assert.Error,
			},
		},
		{
			name: "put without collision",
			args: args[int, *people]{key: 2, val: &people{name: "test", age: 1}},
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return mt
			}(),
			check: true,
			want: want[*people]{
				val:     &people{name: "test", age: 1},
				wantErr: assert.NoError,
				size:    1,
			},
		},
		{
			name: "put with collision",
			args: args[int, *people]{key: 2, val: &people{name: "test", age: 2}},
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = mt.newNode(0, &people{name: "test", age: 0})
				mt.size++
				return mt
			}(),
			check: true,
			want: want[*people]{
				val:     &people{name: "test", age: 2},
				wantErr: assert.NoError,
				size:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want.wantErr(t, tt.ht.Put(tt.args.key, tt.args.val), fmt.Sprintf("Put(%v, %v)", tt.args.key, tt.args.val))
			if tt.check {
				assert.Equalf(t, tt.want.size, tt.ht.size, "Put(%v, %v) size", tt.args.key, tt.args.val)
				i, _ := tt.ht.hasher.makeHash(tt.args.key)
				assert.Equalf(t, *tt.want.val, *iterateThrowCollision(tt.args.key, tt.ht.table[i]), "Put(%v, %v) value", tt.args.key, tt.args.val)
			}
		})
	}
}

func iterateThrowCollision(key int, n *node[int, *people]) *people {
	for cn := n; cn != nil; cn = cn.next {
		if cn.key == key {
			return cn.value
		}
	}
	return new(people)
}

func TestHashTable_Contains(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name    string
		ht      *HashTable[K, T]
		args    K
		want    bool
		wantErr assert.ErrorAssertionFunc
		check   bool
	}
	tests := []testCase[int, *people]{
		{
			name:    "got error by hasher",
			args:    1,
			want:    false,
			wantErr: assert.Error,
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				h := MockHasher[int]{}
				h.On("makeHash", mock.AnythingOfType("int")).
					Return(uint64(0), errors.New("test err"))
				mt.hasher = &h
				return mt
			}(),
		},
		{
			name:    "empty table",
			args:    1,
			want:    false,
			wantErr: assert.NoError,
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return mt
			}(),
			check: true,
		},
		{
			name:    "doesn't contain",
			args:    1,
			want:    false,
			wantErr: assert.NoError,
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 2, value: &people{name: "test", age: 2}}
				return mt
			}(),
			check: true,
		},
		{
			name:    "contains",
			args:    0,
			want:    true,
			wantErr: assert.NoError,
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{
					key:   2,
					value: &people{name: "test", age: 2},
					next:  &node[int, *people]{key: 0, value: &people{age: 0, name: "test"}},
				}
				return mt
			}(),
			check: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ht.Contains(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("Contains(%v)", tt.args)) {
				return
			}
			if tt.check {
				assert.Equalf(t, tt.want, got, "Contains(%v)", tt.args)
			}
		})
	}
}

func TestHashTable_Remove(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name    string
		ht      *HashTable[K, T]
		args    K
		wantErr assert.ErrorAssertionFunc
		check   bool
	}
	tests := []testCase[int, *people]{
		{
			name: "got error by hasher",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				h := MockHasher[int]{}
				h.On("makeHash", mock.AnythingOfType("int")).
					Return(uint64(0), errors.New("test err"))
				mt.hasher = &h
				return mt
			}(),
			wantErr: assert.Error,
		},
		{
			name: "got error when empty value",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return mt
			}(),
			wantErr: assert.Error,
		},
		{
			name: "got error when not find",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 0, value: &people{name: "test", age: 0}}
				return mt
			}(),
			args:    2,
			wantErr: assert.Error,
		},
		{
			name: "remove without collision",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 44, value: &people{name: "test", age: 44}}
				return mt
			}(),
			args:    44,
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "remove with collision in tail",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 44, value: &people{name: "test", age: 44}}
				mt.resolvePutCollision(8, &people{name: "test", age: 8}, 132)
				mt.resolvePutCollision(2, &people{name: "test", age: 2}, 132)
				mt.resolvePutCollision(0, &people{name: "test", age: 0}, 132)
				return mt
			}(),
			args:    44,
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "remove with collision in head",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 44, value: &people{name: "test", age: 44}}
				mt.resolvePutCollision(8, &people{name: "test", age: 8}, 132)
				mt.resolvePutCollision(2, &people{name: "test", age: 2}, 132)
				mt.resolvePutCollision(0, &people{name: "test", age: 0}, 132)
				return mt
			}(),
			args:    0,
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "remove with collision in middle",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 44, value: &people{name: "test", age: 44}}
				mt.resolvePutCollision(8, &people{name: "test", age: 8}, 132)
				mt.resolvePutCollision(2, &people{name: "test", age: 2}, 132)
				mt.resolvePutCollision(0, &people{name: "test", age: 0}, 132)
				return mt
			}(),
			args:    2,
			wantErr: assert.NoError,
			check:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.ht.Remove(tt.args), fmt.Sprintf("Remove(%v)", tt.args))
			if tt.check {
				i, _ := tt.ht.hasher.makeHash(tt.args)
				assert.Equal(t, new(people), iterateThrowCollision(tt.args, tt.ht.table[i]))
			}
		})
	}
}

func TestHashTable_Clear(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name string
		ht   *HashTable[K, T]
	}
	tests := []testCase[int, *people]{
		{
			name: "empty",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return mt
			}(),
		},
		{
			name: "not empty",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{
					key:   8,
					value: &people{name: "test", age: 8},
					next:  &node[int, *people]{key: 44, value: &people{name: "test", age: 44}},
				}
				mt.table[0] = &node[int, *people]{
					key:   99,
					value: &people{name: "test", age: 99},
				}
				mt.size = 3
				return mt
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ht.Clear()
			assert.Zero(t, tt.ht.size)
			assert.Condition(t, func() (success bool) {
				for _, n := range tt.ht.table {
					if n != nil {
						return false
					}
				}
				return true
			}, "after Clear() ht contains not empty values")
		})
	}
}

func TestHashTable_Keys(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name string
		ht   *HashTable[K, T]
		want []K
	}
	tests := []testCase[int, *people]{
		{
			name: "when empty",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return mt
			}(),
			want: ([]int)(nil),
		},
		{
			name: "got keys",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 44, value: &people{name: "test", age: 44}}
				mt.resolvePutCollision(8, &people{name: "test", age: 8}, 132)
				mt.resolvePutCollision(2, &people{name: "test", age: 2}, 132)
				mt.resolvePutCollision(0, &people{name: "test", age: 0}, 132)
				mt.size = 4
				return mt
			}(),
			want: []int{8, 44, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, tt.ht.Keys(), "Keys()")
		})
	}
}

func TestHashTable_Values(t *testing.T) {
	type testCase[K comparable, T any] struct {
		name string
		ht   *HashTable[K, T]
		want []T
	}
	tests := []testCase[int, *people]{
		{
			name: "when empty",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				return mt
			}(),
			want: ([]*people)(nil),
		},
		{
			name: "got keys",
			ht: func() *HashTable[int, *people] {
				mt, _ := NewHashTableWithCapacity[int, *people](uint64(150))
				mt.table[132] = &node[int, *people]{key: 44, value: &people{name: "test", age: 44}}
				mt.resolvePutCollision(8, &people{name: "test", age: 8}, 132)
				mt.resolvePutCollision(2, &people{name: "test", age: 2}, 132)
				mt.resolvePutCollision(0, &people{name: "test", age: 0}, 132)
				mt.size = 4
				return mt
			}(),
			want: []*people{
				{name: "test", age: 44},
				{name: "test", age: 8},
				{name: "test", age: 2},
				{name: "test", age: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, tt.ht.Values(), "Values()")
		})
	}
}
