package lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleLinkedList_PushTail(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type want[T any] struct {
		wantErr bool
		size    int
		tail    T
	}
	type testCase[T any] struct {
		name string
		sl   *SingleLinkedList[T]
		args args[T]
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "added 1 node before",
			sl: func() *SingleLinkedList[int] {
				l := NewSingleList[int]()
				l.tail = newSingleNode[int](1)
				l.head = newSingleNode[int](1)
				l.size++
				return l
			}(),
			args: args[int]{value: 5},
			want: want[int]{
				tail: 5,
				size: 2,
			},
		},
		{
			name: "empty",
			sl:   NewSingleList[int](),
			args: args[int]{value: 5},
			want: want[int]{
				tail: 5,
				size: 1,
			},
		},
		{
			name: "last node isn't nil",
			sl: func() *SingleLinkedList[int] {
				l := NewSingleList[int]()
				l.tail = newSingleNode[int](1)
				l.head = newSingleNode[int](1)
				l.tail.next = newSingleNode[int](3)
				l.size = 2
				return l
			}(),
			args: args[int]{value: 5},
			want: want[int]{
				tail:    1,
				size:    2,
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sl.PushTail(tt.args.value); (err != nil) != tt.want.wantErr {
				t.Errorf("PushTail() error = %v, wantErr %v", err, tt.want.wantErr)
			}
			assert.Equalf(t, tt.sl.tail.value, tt.want.tail, "PushTail(%v) got tail", tt.args.value)
			assert.Equalf(t, tt.sl.size, tt.want.size, "PushTail(%v) got size", tt.args.value)
		})
	}
}

func TestSingleLinkedList_Map(t *testing.T) {
	type args[T any] struct {
		each OnEach[T]
	}
	type want[T any] struct {
		wantErr bool
		sl      *SingleLinkedList[T]
	}
	type testCase[T any] struct {
		name string
		sl   *SingleLinkedList[T]
		args args[T]
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "multiply",
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				_ = sl.PushTail(2)
				_ = sl.PushTail(4)
				_ = sl.PushTail(6)
				return sl
			}(),
			args: args[int]{
				each: func(value *int) error {
					*value *= *value
					return nil
				},
			},
			want: want[int]{
				sl: func() *SingleLinkedList[int] {
					sl := NewSingleList[int]()
					_ = sl.PushTail(4)
					_ = sl.PushTail(16)
					_ = sl.PushTail(36)
					return sl
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sl.Map(tt.args.each)
			assert.EqualValues(t, tt.sl, tt.want.sl, "forEach() got values")
			assert.Condition(t, func() (success bool) {
				return (nil != err) == tt.want.wantErr
			}, "forEach() got error")
		})
	}
}

func TestSingleLinkedList_PushHead(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type want[T any] struct {
		size int
		head T
	}
	type testCase[T any] struct {
		name string
		sl   *SingleLinkedList[T]
		args args[T]
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "empty",
			sl:   NewSingleList[int](),
			args: args[int]{value: 4},
			want: want[int]{
				head: 4,
				size: 1,
			},
		},
		{
			name: "not empty",
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				node := newSingleNode(2)
				sl.head = node
				sl.tail = node
				sl.size++
				return sl
			}(),
			args: args[int]{value: 4},
			want: want[int]{
				head: 4,
				size: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sl.PushHead(tt.args.value)
			assert.Equalf(t, tt.want.head, tt.sl.head.value, "pushHead(%v) head", tt.args)
			assert.Equalf(t, tt.want.size, tt.sl.size, "pushHead(%v) size", tt.args)
		})
	}
}

func TestSingleLinkedList_Insert(t *testing.T) {
	type args[T any] struct {
		index int
		value T
	}
	type want[T any] struct {
		wantErr bool
		sl      *SingleLinkedList[T]
	}
	type testCase[T any] struct {
		name string
		sl   *SingleLinkedList[T]
		args args[T]
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "insert to head",
			sl:   NewSingleList[int](),
			args: args[int]{
				value: 10,
				index: 0,
			},
			want: want[int]{
				sl: func() *SingleLinkedList[int] {
					sl := NewSingleList[int]()
					node := newSingleNode(10)
					sl.head = node
					sl.tail = node
					sl.size++
					return sl
				}(),
			},
		},
		{
			name: "insert to tail",
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				node := newSingleNode(2)
				sl.head = node
				sl.tail = node
				sl.size++
				return sl
			}(),
			args: args[int]{
				value: 10,
				index: 1,
			},
			want: want[int]{
				sl: func() *SingleLinkedList[int] {
					sl := NewSingleList[int]()
					node := newSingleNode(2)
					node.next = newSingleNode(10)
					sl.head = node
					sl.tail = node.next
					sl.size = 2
					return sl
				}(),
			},
		},
		{
			name: "got err",
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				node := newSingleNode(2)
				sl.head = node
				sl.tail = node
				sl.size++
				return sl
			}(),
			args: args[int]{
				value: 10,
				index: 2,
			},
			want: want[int]{
				sl: func() *SingleLinkedList[int] {
					sl := NewSingleList[int]()
					node := newSingleNode(2)
					sl.head = node
					sl.tail = node
					sl.size++
					return sl
				}(),
				wantErr: true,
			},
		},
		{
			name: "insert to middle",
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				node := newSingleNode(2)
				tailNode := newSingleNode(15)
				node.next = newSingleNode(5)
				node.next.next = tailNode
				sl.head = node
				sl.tail = tailNode
				sl.size = 3
				return sl
			}(),
			args: args[int]{
				value: 3,
				index: 1,
			},
			want: want[int]{
				sl: func() *SingleLinkedList[int] {
					sl := NewSingleList[int]()
					node := newSingleNode(2)
					tailNode := newSingleNode(15)
					node.next = newSingleNode(3)
					node.next.next = newSingleNode(5)
					node.next.next.next = tailNode
					sl.head = node
					sl.tail = tailNode
					sl.size = 4
					return sl
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sl.Insert(tt.args.index, tt.args.value)
			assert.Conditionf(t, func() (success bool) {
				return (err != nil) == tt.want.wantErr
			}, "insert(%d,%v) got error", tt.args.index, tt.args.value)
			assert.Equalf(t, tt.want.sl, tt.sl, "insert(%d,%v)", tt.args.index, tt.args.value)
		})
	}
}

func TestSingleLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	type want[T any] struct {
		wantErr bool
		value   T
	}
	type testCase[T any] struct {
		name string
		sl   *SingleLinkedList[T]
		args args
		want want[T]
	}
	tests := []testCase[int]{
		{
			name: "got error",
			args: args{index: 2},
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				node := newSingleNode(2)
				lastNode := newSingleNode(4)
				node.next = lastNode
				sl.head = node
				sl.tail = lastNode
				sl.size = 2
				return sl
			}(),
			want: want[int]{
				wantErr: true,
				value:   0,
			},
		},
		{
			name: "got value",
			args: args{index: 2},
			sl: func() *SingleLinkedList[int] {
				sl := NewSingleList[int]()
				node := newSingleNode(2)
				lastNode := newSingleNode(8)
				node.next = newSingleNode(4)
				node.next.next = newSingleNode(6)
				node.next.next.next = lastNode
				sl.head = node
				sl.tail = lastNode
				sl.size = 4
				return sl
			}(),
			want: want[int]{
				value: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sl.Get(tt.args.index)
			if (err != nil) != tt.want.wantErr {
				t.Errorf("Get(%v)", tt.args.index)
			}
			assert.Equalf(t, tt.want.value, got, "Get(%v)", tt.args.index)
		})
	}
}

func TestSingleLinkedList_Delete(t *testing.T) {
	type args struct {
		index int
	}
	type want[T any] struct {
		wantErr assert.ErrorAssertionFunc
		value   T
		size    int
	}
	type testCase[T any] struct {
		name  string
		args  args
		want  want[T]
		check bool
	}
	tests := []testCase[int]{
		{
			name: "got err",
			args: args{index: 7},
			want: want[int]{
				wantErr: assert.Error,
			},
		},
		{
			name: "delete head",
			args: args{},
			want: want[int]{
				wantErr: assert.NoError,
				size:    5,
				value:   1,
			},
			check: true,
		},
		{
			name: "delete tail",
			args: args{index: 5},
			want: want[int]{
				wantErr: assert.NoError,
				size:    5,
				value:   10,
			},
			check: true,
		},
		{
			name: "delete middle value",
			args: args{index: 3},
			want: want[int]{
				wantErr: assert.NoError,
				size:    5,
				value:   6,
			},
			check: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := fetchSL()
			got, gotErr := sl.Delete(tt.args.index)
			tt.want.wantErr(t, gotErr, fmt.Sprintf("Delete(%v)", tt.args.index))
			if tt.check {
				assert.Equalf(t, got, tt.want.value, fmt.Sprintf("Delete(%v) check value", tt.args.index))
				assert.Equalf(t, sl.size, tt.want.size, fmt.Sprintf("Delete(%v) check size", tt.args.index))
			}
		})
	}
}

func fetchSL() *SingleLinkedList[int] {
	sl := NewSingleList[int]()
	head := newSingleNode(1)
	current := head
	all := 6
	for i := 1; i < all; i++ {
		current.next = newSingleNode(i << 1)
		current = current.next
	}
	sl.head = head
	sl.tail = current
	sl.size = all
	return sl
}
