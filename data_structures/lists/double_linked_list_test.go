package lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLinkedList_PushHead(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name    string
		dl      *DoubleLinkedList[T]
		args    args[T]
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[int]{
		{
			name:    "push than empty",
			dl:      fetchDoubleLinkedList(0),
			args:    args[int]{value: 0},
			wantErr: assert.NoError,
		},
		{
			name:    "push than not empty",
			dl:      fetchDoubleLinkedList(4),
			args:    args[int]{value: 0},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.dl.PushHead(tt.args.value), fmt.Sprintf("PushHead(%v)", tt.args.value))
			assert.Equal(t, tt.args.value, tt.dl.head.value)
		})
	}
}

func fetchDoubleLinkedList(counter int) *DoubleLinkedList[int] {
	dl := NewDoubleLinkedList[int]()
	if counter < 1 {
		return dl
	}
	head := NewDoubleNode(1)
	dl.head = head
	dl.tail = head
	dl.size++
	for i := 1; i < counter; i++ {
		node := dl.head
		for k := 1; k < i; k++ {
			node = node.next
		}
		newNode := NewDoubleNode(i << 1)
		newNode.previous = node
		node.next = newNode
		dl.tail = newNode
		dl.size++
	}

	return dl
}

func TestDoubleLinkedList_PushTail(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name    string
		dl      *DoubleLinkedList[T]
		args    args[T]
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[int]{
		{
			name:    "push than empty",
			dl:      fetchDoubleLinkedList(0),
			args:    args[int]{value: 0},
			wantErr: assert.NoError,
		},
		{
			name:    "push than not empty",
			dl:      fetchDoubleLinkedList(4),
			args:    args[int]{value: 0},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.dl.PushTail(tt.args.value), fmt.Sprintf("PushTail(%v)", tt.args.value))
			assert.Equal(t, tt.dl.tail.value, tt.args.value)
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	type args[T any] struct {
		index int
		value T
	}
	type testCase[T any] struct {
		name       string
		dl         *DoubleLinkedList[T]
		args       args[T]
		wantErr    assert.ErrorAssertionFunc
		checkIndex int
	}
	tests := []testCase[int]{
		{
			name: "got error because of empty dl",
			dl:   fetchDoubleLinkedList(0),
			args: args[int]{
				value: 0,
				index: 0,
			},
			wantErr:    assert.Error,
			checkIndex: 0,
		},
		{
			name: "push head",
			dl:   fetchDoubleLinkedList(4),
			args: args[int]{
				value: 10,
				index: 0,
			},
			wantErr:    assert.NoError,
			checkIndex: 0,
		},
		{
			name: "push tail",
			dl:   fetchDoubleLinkedList(4),
			args: args[int]{
				value: 10,
				index: 3,
			},
			wantErr:    assert.NoError,
			checkIndex: 4,
		},
		{
			name: "push middle",
			dl:   fetchDoubleLinkedList(4),
			args: args[int]{
				value: 10,
				index: 1,
			},
			wantErr:    assert.NoError,
			checkIndex: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.dl.Insert(tt.args.index, tt.args.value), fmt.Sprintf("Insert(%v, %v)", tt.args.index, tt.args.value))
			assert.Equal(t, fetchValueByIndex[int](tt.checkIndex, tt.dl), tt.args.value)
		})
	}
}

func fetchValueByIndex[T any](index int, dl *DoubleLinkedList[T]) T {
	if dl.size < 1 {
		return *new(T)
	}
	node := dl.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

func TestDoubleLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	type want struct {
		wantErr assert.ErrorAssertionFunc
	}
	type testCase[T any] struct {
		name       string
		dl         *DoubleLinkedList[T]
		args       args
		want       want
		checkIndex int
	}
	tests := []testCase[int]{
		{
			name: "got error",
			args: args{index: 0},
			dl:   fetchDoubleLinkedList(0),
			want: want{
				wantErr: assert.Error,
			},
		},
		{
			name: "get head",
			args: args{index: 0},
			dl:   fetchDoubleLinkedList(4),
			want: want{
				wantErr: assert.NoError,
			},
			checkIndex: 0,
		},
		{
			name: "get tail",
			args: args{index: 3},
			dl:   fetchDoubleLinkedList(4),
			want: want{
				wantErr: assert.NoError,
			},
			checkIndex: 3,
		},
		{
			name: "get middle",
			args: args{index: 2},
			dl:   fetchDoubleLinkedList(4),
			want: want{
				wantErr: assert.NoError,
			},
			checkIndex: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dl.Get(tt.args.index)
			if !tt.want.wantErr(t, err, fmt.Sprintf("Get(%v)", tt.args.index)) {
				return
			}
			assert.Equalf(t, fetchValueByIndex[int](tt.checkIndex, tt.dl), got, "Get(%v)", tt.checkIndex)
		})
	}
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	type args struct {
		index int
	}
	type want[T any] struct {
		value       T
		wantErr     assert.ErrorAssertionFunc
		removeValue T
	}
	type testCase[T any] struct {
		name       string
		dl         *DoubleLinkedList[T]
		args       args
		want       want[int]
		checkIndex int
	}
	tests := []testCase[int]{
		{
			name: "got error",
			dl:   fetchDoubleLinkedList(0),
			args: args{index: 0},
			want: want[int]{
				wantErr: assert.Error,
			},
		},
		{
			name: "remove head",
			dl:   fetchDoubleLinkedList(4),
			args: args{index: 0},
			want: want[int]{
				removeValue: 1,
				value:       2,
				wantErr:     assert.NoError,
			},
		},
		{
			name: "remove tail",
			dl:   fetchDoubleLinkedList(4),
			args: args{index: 3},
			want: want[int]{
				removeValue: 6,
				value:       4,
				wantErr:     assert.NoError,
			},
			checkIndex: 2,
		},
		{
			name: "remove middle",
			dl:   fetchDoubleLinkedList(4),
			args: args{index: 1},
			want: want[int]{
				removeValue: 2,
				value:       4,
				wantErr:     assert.NoError,
			},
			checkIndex: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dl.Remove(tt.args.index)
			if !tt.want.wantErr(t, err, fmt.Sprintf("Remove(%v)", tt.args.index)) {
				return
			}
			assert.Equalf(t, tt.want.removeValue, got, "Remove(%v)", tt.args.index)
			assert.Equalf(t, fetchValueByIndex[int](tt.checkIndex, tt.dl), tt.want.value, "Remove(%v) got new by index", tt.checkIndex)
		})
	}
}
