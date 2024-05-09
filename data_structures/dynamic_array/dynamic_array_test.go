package dynamic_array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicArray_Remove(t *testing.T) {
	arr := []int{15, 9, 2, 7}
	type want[T any] struct {
		err    bool
		arr    []T
		length uint64
	}
	type testCase[T any] struct {
		name string
		da   *DynamicArray[T]
		num  uint64
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "remove first element",
			da:   fetchDA[int](5, arr),
			num:  0,
			want: want[int]{arr: []int{9, 2, 7}, length: 3},
		},
		{
			name: "remove last element",
			da:   fetchDA[int](3, arr),
			num:  3,
			want: want[int]{arr: []int{15, 9, 2}, length: 3},
		},
		{
			name: "remove middle element",
			da:   fetchDA[int](5, arr),
			num:  1,
			want: want[int]{arr: []int{15, 2, 7}, length: 3},
		},
		{
			name: "get out of range",
			da:   fetchDA[int](5, arr),
			want: want[int]{arr: arr, length: 4, err: true},
			num:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.da.Remove(tt.num)
			assert.Equal(t, err != nil, tt.want.err)
			assert.Equal(t, tt.want.arr, tt.da.All())
			assert.Equal(t, tt.want.length, tt.da.Length())
		})
	}
}

func fetchDA[T any](capacity uint64, elements []T) *DynamicArray[T] {
	da := NewDynamicArray[T](capacity)
	for _, element := range elements {
		da.Add(element)
	}

	return da
}

func TestDynamicArray_IsEmpty(t *testing.T) {
	type Cat struct {
		name string
		age  int
	}
	type testCase[T Cat] struct {
		name string
		da   *DynamicArray[T]
		want bool
	}
	tests := []testCase[Cat]{
		{
			name: "not empty",
			da:   &DynamicArray[Cat]{length: uint64(1)},
			want: true,
		},
		{
			name: "empty",
			da:   &DynamicArray[Cat]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.da.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestDynamicArray_expand(t *testing.T) {
	type testCase[T int] struct {
		name string
		da   *DynamicArray[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "n=1",
			da:   NewDynamicArray[int](uint64(1)),
			want: 2,
		},
		{
			name: "n=3",
			da:   NewDynamicArray[int](uint64(3)),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.da.expand()
			assert.Equal(t, tt.want, cap(tt.da.arr), "arr n")
			assert.Equal(t, tt.want, int(tt.da.capacity), "da n")
		})
	}
}

func TestDynamicArray_checkIndex(t *testing.T) {
	type args struct {
		n uint64
	}
	type testCase[T int] struct {
		name string
		da   *DynamicArray[T]
		args args
		want error
	}
	tests := []testCase[int]{
		{
			name: "n = length",
			da:   &DynamicArray[int]{arr: []int{1, 2, 0}, length: uint64(3), capacity: uint64(3)},
			args: args{n: 2},
			want: nil,
		},
		{
			name: "n > length",
			da:   &DynamicArray[int]{arr: []int{1, 2, 0}, length: uint64(3), capacity: uint64(3)},
			args: args{n: 3},
			want: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.da.checkIndex(tt.args.n), tt.want, fmt.Sprintf("checkIndex(%v)", tt.args.n))
		})
	}
}

func TestNewDynamicArray(t *testing.T) {
	type args struct {
		n uint64
	}
	type want[T int] struct {
		da    *DynamicArray[T]
		panic bool
	}
	type testCase[T int] struct {
		name string
		args args
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "got panic",
			want: want[int]{da: nil, panic: true},
		},
		{
			name: "n = 5",
			args: args{n: 5},
			want: want[int]{da: &DynamicArray[int]{capacity: uint64(5), arr: make([]int, 5)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want.panic {
				assert.PanicsWithValue(t, "capacity less than 1", func() {
					NewDynamicArray[int](tt.args.n)
				}, "NewDynamicArray(%v) panic", tt.args.n)
			} else {
				da := NewDynamicArray[int](tt.args.n)
				assert.Equalf(t, tt.want.da, da, "NewDynamicArray(%v)", tt.args.n)
			}
		})
	}
}

func TestDynamicArray_Add(t *testing.T) {
	type args struct {
		element []int
	}
	type want struct {
		l uint64
		c uint64
	}
	type testCase[T int] struct {
		name string
		args args
		want want
	}
	tests := []testCase[int]{
		{
			name: "add 1 element",
			args: args{element: []int{5}},
			want: want{l: 1, c: 1},
		},
		{
			name: "add 3 elements",
			args: args{element: []int{5, 7, 9}},
			want: want{l: 3, c: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := NewDynamicArray[int](1)
			for _, e := range tt.args.element {
				da.Add(e)
			}
			assert.Equalf(t, tt.want.c, da.Capacity(), "Add(%v) capacity", tt.args.element)
			assert.Equalf(t, tt.want.l, da.Length(), "Add(%v) length", tt.args.element)
		})
	}
}

func TestDynamicArray_All(t *testing.T) {
	type testCase[T int] struct {
		name string
		da   *DynamicArray[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty",
			da:   NewDynamicArray[int](uint64(1)),
			want: []int{},
		},
		{
			name: "3 elements",
			da: NewDynamicArray[int](uint64(1)).
				Add(1).
				Add(2).
				Add(3),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.da.All(), "All()")
		})
	}
}

func TestDynamicArray_Get(t *testing.T) {
	type args struct {
		n uint64
	}
	type want[T any] struct {
		element T
		err     error
	}
	type testCase[T any] struct {
		name string
		da   *DynamicArray[T]
		args args
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "got error",
			da:   NewDynamicArray[int](uint64(3)).Add(1),
			args: args{n: 3},
			want: want[int]{
				element: 0,
				err:     ErrIndexOutOfRange,
			},
		},
		{
			name: "got 2",
			da: NewDynamicArray[int](uint64(2)).
				Add(1).
				Add(4).
				Add(2),
			args: args{n: 2},
			want: want[int]{
				element: 2,
				err:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.da.Get(tt.args.n)
			assert.Equalf(t, got, tt.want.element, "Get(%v) element", tt.args.n)
			assert.Equalf(t, err, tt.want.err, "Get(%v) error", tt.args.n)
		})
	}
}

func TestDynamicArray_Put(t *testing.T) {
	type args[T any] struct {
		element T
		n       uint64
	}
	type want[T any] struct {
		err     error
		element T
	}
	type testCase[T any] struct {
		name string
		da   *DynamicArray[T]
		args args[T]
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "got error",
			da:   NewDynamicArray[int](uint64(3)).Add(1),
			args: args[int]{n: 3},
			want: want[int]{
				err: ErrIndexOutOfRange,
			},
		},
		{
			name: "got 6",
			da: NewDynamicArray[int](uint64(2)).
				Add(1).
				Add(4).
				Add(2),
			args: args[int]{n: 1, element: 6},
			want: want[int]{
				element: 6,
				err:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.da.Put(tt.args.n, tt.args.element), tt.want.err, fmt.Sprintf("Put(%v, %v)", tt.args.n, tt.args.element))
			got, _ := tt.da.Get(tt.args.n)
			assert.Equal(t, got, tt.args.element, fmt.Sprintf("Put(%v, %v) got element", tt.args.n, tt.args.element))
		})
	}
}
