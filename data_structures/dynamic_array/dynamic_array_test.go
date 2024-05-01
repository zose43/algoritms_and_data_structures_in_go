package dynamic_array

import (
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
