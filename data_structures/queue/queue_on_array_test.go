package queue

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type testCase[T any] struct {
		name string
		args uint
		want *Queue[T]
	}
	tests := []testCase[int]{
		{
			name: "fixed size constant",
			want: &Queue[int]{fixedSize: defaultQueueCapacity, data: make([]int, 0)},
		},
		{
			name: "fixed size custom",
			args: 4,
			want: &Queue[int]{fixedSize: 4, data: make([]int, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[int](tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type testCase[T any] struct {
		name       string
		q          *Queue[T]
		args       int
		wantErr    bool
		checkQueue *Queue[T]
	}
	tests := []testCase[int]{
		{
			name:       "enqueue with error",
			args:       2,
			wantErr:    true,
			q:          &Queue[int]{fixedSize: 3, data: make([]int, 3)},
			checkQueue: &Queue[int]{fixedSize: 3, data: make([]int, 3)},
		},
		{
			name:       "enqueue when empty",
			args:       2,
			wantErr:    false,
			q:          &Queue[int]{fixedSize: 3, data: make([]int, 0)},
			checkQueue: &Queue[int]{fixedSize: 3, data: []int{2}},
		},
		{
			name:       "enqueue when not empty",
			args:       2,
			wantErr:    false,
			q:          &Queue[int]{fixedSize: 5, data: []int{4, 6, 8, 9}},
			checkQueue: &Queue[int]{fixedSize: 5, data: []int{4, 6, 8, 9, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.Enqueue(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.checkQueue, tt.q)
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type testCase[T any] struct {
		name       string
		q          *Queue[T]
		want       T
		isExist    bool
		checkQueue *Queue[T]
	}
	tests := []testCase[int]{
		{
			name:       "dequeue with empty",
			q:          &Queue[int]{fixedSize: 5, data: make([]int, 0)},
			want:       0,
			isExist:    false,
			checkQueue: &Queue[int]{fixedSize: 5, data: make([]int, 0)},
		},
		{
			name:       "dequeue when not empty",
			q:          &Queue[int]{fixedSize: 5, data: []int{2, 5, 8, 9, 12}},
			want:       2,
			isExist:    true,
			checkQueue: &Queue[int]{fixedSize: 5, data: []int{5, 8, 9, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.q.Dequeue()
			assert.Equalf(t, tt.want, got, "Dequeue()")
			assert.Equalf(t, tt.isExist, ok, "Dequeue()")
			assert.Equal(t, tt.checkQueue, tt.q)
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type testCase[T any] struct {
		name       string
		q          *Queue[T]
		want       T
		isExist    bool
		checkQueue *Queue[T]
	}
	tests := []testCase[int]{
		{
			name:       "dequeue with empty",
			q:          &Queue[int]{fixedSize: 5, data: make([]int, 0)},
			want:       0,
			isExist:    false,
			checkQueue: &Queue[int]{fixedSize: 5, data: make([]int, 0)},
		},
		{
			name:       "dequeue when not empty",
			q:          &Queue[int]{fixedSize: 5, data: []int{2, 5, 8, 9, 12}},
			want:       2,
			isExist:    true,
			checkQueue: &Queue[int]{fixedSize: 5, data: []int{2, 5, 8, 9, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.q.Peek()
			assert.Equalf(t, tt.want, got, "Peek()")
			assert.Equalf(t, tt.isExist, ok, "Peek()")
			assert.Equal(t, tt.checkQueue, tt.q)
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    *Queue[T]
	}
	tests := []testCase[int]{
		{
			name: "clear when empty",
			q:    &Queue[int]{fixedSize: 5, data: make([]int, 0)},
		},
		{
			name: "clear when not empty",
			q:    &Queue[int]{fixedSize: 5, data: []int{2, 5, 8, 9, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Clear()
			assert.Empty(t, tt.q.data)
		})
	}
}
