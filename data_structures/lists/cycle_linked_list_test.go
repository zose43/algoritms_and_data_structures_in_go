package lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycleLinkedList_Add(t *testing.T) {
	type want[T any] struct {
		size int
		val  T
	}
	type testCase[T any] struct {
		name string
		cl   *CycleLinkedList[T]
		args T
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "add when empty",
			cl:   fetchCycledList(0),
			args: 1,
			want: want[int]{val: 1, size: 1},
		},
		{
			name: "add",
			cl:   fetchCycledList(3),
			args: 8,
			want: want[int]{val: 8, size: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cl.Add(tt.args)
			assert.Equalf(t, tt.cl.head.value, tt.want.val, "add(%d) value", tt.args)
			assert.Equalf(t, tt.cl.size, tt.want.size, "add(%d) size", tt.args)
		})
	}
}

func fetchCycledList(count int) *CycleLinkedList[int] {
	cl := NewCycleLinkedList[int]()
	if count < 1 {
		return cl
	}
	node := newCycledLinkedNode[int](1)
	cl.head = node
	cl.head.next = node
	cl.head.prev = node
	cl.size++
	for i := 1; count >= i; i++ {
		val := i << 1
		node = newCycledLinkedNode[int](val)
		node.next = cl.head
		node.prev = cl.head.prev
		cl.head.prev.next = node
		cl.head.prev = node
		cl.head = node
		cl.size++
	}

	return cl
}

func TestCycleLinkedList_Value(t *testing.T) {
	type testCase[T any] struct {
		name    string
		cl      *CycleLinkedList[T]
		want    T
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[int]{
		{
			name:    "got err",
			cl:      fetchCycledList(0),
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "got value",
			cl:      fetchCycledList(3),
			want:    6,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cl.Value()
			if !tt.wantErr(t, err, fmt.Sprintf("Value()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "Value()")
		})
	}
}

func TestCycleLinkedList_Rotate(t *testing.T) {
	type want[T any] struct {
		val T
	}
	type testCase[T any] struct {
		name string
		cl   *CycleLinkedList[T]
		want want[T]
		arg  T
	}
	tests := []testCase[int]{
		{
			name: "rotate next 1",
			cl:   fetchCycledList(3),
			want: want[int]{val: 4},
			arg:  1,
		},
		{
			name: "rotate next 3",
			cl:   fetchCycledList(3),
			want: want[int]{val: 1},
			arg:  3,
		},
		{
			name: "rotate next 4",
			cl:   fetchCycledList(3),
			want: want[int]{val: 6},
			arg:  4,
		},
		{
			name: "rotate prev 1",
			cl:   fetchCycledList(3),
			want: want[int]{val: 1},
			arg:  -1,
		},
		{
			name: "rotate prev 3",
			cl:   fetchCycledList(3),
			want: want[int]{val: 4},
			arg:  -3,
		},
		{
			name: "rotate prev 4",
			cl:   fetchCycledList(3),
			want: want[int]{val: 6},
			arg:  -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cl.Rotate(tt.arg)
			assert.Equalf(t, tt.cl.head.value, tt.want.val, "Rotatet(%d)", tt.arg)
		})
	}
}

func TestCycleLinkedList_Remove(t *testing.T) {
	type want[T any] struct {
		val    T
		size   int
		result bool
	}
	type testCase[T any] struct {
		name  string
		cl    *CycleLinkedList[T]
		want  want[T]
		check bool
	}
	tests := []testCase[int]{
		{
			name: "remove when empty",
			cl:   fetchCycledList(0),
		},
		{
			name: "remove",
			cl:   fetchCycledList(4),
			want: want[int]{
				val:    6,
				size:   4,
				result: true,
			},
			check: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want.result, tt.cl.Remove(), "Remove()")
			if tt.check {
				assert.Equalf(t, tt.want.size, tt.cl.size, "Remove() size")
				assert.Equalf(t, tt.want.val, tt.cl.head.value, "Remove() value")
			}
		})
	}
}
