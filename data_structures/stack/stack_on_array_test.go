package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	type args struct {
		fixedSize uint
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    *Stack[T]
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "got error",
			wantErr: true,
		},
		{
			name:    "empty stack",
			args:    args{fixedSize: 4},
			wantErr: false,
			want:    &Stack[int]{fixedSize: uint(4), elements: make([]int, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStack[int](tt.args.fixedSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStack_Clear(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    *Stack[T]
		want *Stack[T]
	}
	tests := []testCase[int]{
		{
			name: "empty stack",
			s:    &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
			want: &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
		},
		{
			name: "full stack",
			s:    &Stack[int]{fixedSize: 3, elements: []int{1, 2, 3}},
			want: &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		s       *Stack[T]
		want    T
		isExist bool
	}
	tests := []testCase[int]{
		{
			name:    "empty stack",
			s:       &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
			want:    0,
			isExist: false,
		},
		{
			name:    "full stack",
			s:       &Stack[int]{fixedSize: 3, elements: []int{1, 2, 3}},
			want:    1,
			isExist: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.s.Peek()
			assert.Equalf(t, tt.want, got, "Peek()")
			assert.Equalf(t, tt.isExist, ok, "Peek()")
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type testCase[T any] struct {
		name    string
		s       *Stack[T]
		want    T
		isExist bool
		src     *Stack[T]
	}
	tests := []testCase[int]{
		{
			name: "empty stack",
			s:    &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
			src:  &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
			want: 0,
		},
		{
			name:    "full stack",
			s:       &Stack[int]{fixedSize: 3, elements: []int{1, 2, 3}},
			src:     &Stack[int]{fixedSize: 3, elements: []int{2, 3}},
			want:    1,
			isExist: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.s.Pop()
			assert.Equalf(t, tt.want, got, "Pop()")
			assert.Equalf(t, tt.isExist, ok, "Pop()")
			assert.Equalf(t, tt.src, tt.s, "Pop()")
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		value int
	}
	type testCase[T any] struct {
		name    string
		s       *Stack[T]
		src     *Stack[T]
		args    args
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[int]{
		{
			name:    "got error",
			s:       &Stack[int]{fixedSize: 3, elements: []int{1, 2, 3}},
			src:     &Stack[int]{fixedSize: 3, elements: []int{1, 2, 3}},
			wantErr: assert.Error,
			args:    args{value: 4},
		},
		{
			name:    "empty stack",
			s:       &Stack[int]{fixedSize: 3, elements: make([]int, 0)},
			src:     &Stack[int]{fixedSize: 3, elements: []int{1}},
			wantErr: assert.NoError,
			args:    args{value: 1},
		},
		{
			name:    "full stack",
			s:       &Stack[int]{fixedSize: 3, elements: []int{1, 2}},
			src:     &Stack[int]{fixedSize: 3, elements: []int{3, 1, 2}},
			wantErr: assert.NoError,
			args:    args{value: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.s.Push(tt.args.value), fmt.Sprintf("Push(%v)", tt.args.value))
			assert.Equal(t, tt.src, tt.s)
		})
	}
}
