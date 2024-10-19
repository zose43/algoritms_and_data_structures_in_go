package tree

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TreeExampleElement struct {
	data any
	key  int
}

func (e *TreeExampleElement) Key() int {
	return e.key
}

func TestBinaryTree_Insert(t *testing.T) {
	type args[T Rib] struct {
		Value T
	}
	type testCase[T Rib] struct {
		name   string
		tree   *BinaryTree[T]
		args   args[T]
		result *BinaryTree[T]
	}
	tests := []testCase[*TreeExampleElement]{
		{
			name:   "Insert when empty",
			args:   args[*TreeExampleElement]{Value: &TreeExampleElement{key: 50}},
			tree:   &BinaryTree[*TreeExampleElement]{},
			result: fetchFilledBinaryTree([]*TreeExampleElement{{key: 50}}),
		},
		{
			name: "Insert when full on the right leaf",
			args: args[*TreeExampleElement]{Value: &TreeExampleElement{key: 80}},
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 57},
				{key: 54},
				{key: 25},
				{key: 5},
				{key: 15},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 57},
				{key: 54},
				{key: 25},
				{key: 5},
				{key: 15},
				{key: 80},
			}),
		},
		{
			name: "Insert when full on the left leaf",
			args: args[*TreeExampleElement]{Value: &TreeExampleElement{key: 10}},
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 57},
				{key: 54},
				{key: 25},
				{key: 5},
				{key: 15},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 57},
				{key: 54},
				{key: 25},
				{key: 5},
				{key: 15},
				{key: 10},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tree.Insert(tt.args.Value)
			assert.Equal(t, tt.result, tt.tree)
		})
	}
}

func fetchFilledBinaryTree(src []*TreeExampleElement) *BinaryTree[*TreeExampleElement] {
	tr := NewBinaryTree[*TreeExampleElement]()
	if len(src) == 0 {
		return tr
	}
	tr.root = newNode(src[0])
	if len(src) == 1 {
		return tr
	}
	for _, element := range src[1:] {
		var isLeft bool
		n := tr.root
		parent := n
		for n != nil {
			parent = n
			if n.key() > element.Key() {
				isLeft = true
				n = n.left
			} else {
				isLeft = false
				n = n.right
			}
			if n == nil {
				if isLeft {
					parent.left = newNode(element)
				} else {
					parent.right = newNode(element)
				}
			}
		}
	}
	return tr
}

func TestBinaryTree_Find(t *testing.T) {
	type args struct {
		key int
	}
	type testCase[T Rib] struct {
		name  string
		tree  *BinaryTree[T]
		args  args
		want  T
		exist bool
	}
	tests := []testCase[*TreeExampleElement]{
		{
			name:  "Find when empty",
			tree:  fetchFilledBinaryTree([]*TreeExampleElement{}),
			args:  args{key: 50},
			exist: false,
			want:  nil,
		},
		{
			name: "don't find the element",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 40},
				{key: 60},
				{key: 100},
				{key: 10},
				{key: 15},
				{key: 25},
			}),
			args:  args{key: 50},
			exist: false,
			want:  nil,
		},
		{
			name: "find when element is root",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 40},
				{key: 60},
				{key: 100},
				{key: 10},
				{key: 15},
				{key: 25},
			}),
			args:  args{key: 55},
			exist: true,
			want:  &TreeExampleElement{key: 55},
		},
		{
			name: "find when element is right leaf",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 40},
				{key: 60},
				{key: 100},
				{key: 10},
				{key: 15},
				{key: 25},
			}),
			args:  args{key: 100},
			exist: true,
			want:  &TreeExampleElement{key: 100},
		},
		{
			name: "find when element is middle element",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 40},
				{key: 60},
				{key: 100},
				{key: 10},
				{key: 15},
				{key: 25},
			}),
			args:  args{key: 40},
			exist: true,
			want:  &TreeExampleElement{key: 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, isExist := tt.tree.Find(tt.args.key)
			assert.Equalf(t, tt.want, got, "Find(%v)", tt.args.key)
			assert.Equalf(t, tt.exist, isExist, "Find(%v)", tt.args.key)
		})
	}
}

func TestBinaryTree_Minimum(t *testing.T) {
	type testCase[T Rib] struct {
		name  string
		tree  *BinaryTree[T]
		want  T
		exist bool
	}
	tests := []testCase[*TreeExampleElement]{
		{
			name:  "Minimum when empty",
			tree:  fetchFilledBinaryTree([]*TreeExampleElement{}),
			exist: false,
			want:  nil,
		},
		{
			name: "Minimum when left node is empty",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 100},
			}),
			exist: true,
			want:  &TreeExampleElement{key: 50},
		},
		{
			name: "Minimum when tree is full",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 100},
				{key: 30},
				{key: 15},
			}),
			exist: true,
			want:  &TreeExampleElement{key: 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, isExist := tt.tree.Minimum()
			assert.Equalf(t, tt.want, got, "Minimum()")
			assert.Equalf(t, tt.exist, isExist, "Minimum()")
		})
	}
}

func TestBinaryTree_Maximum(t *testing.T) {
	type testCase[T Rib] struct {
		name  string
		tree  *BinaryTree[T]
		want  T
		exist bool
	}
	tests := []testCase[*TreeExampleElement]{
		{
			name:  "Maximum when empty",
			tree:  fetchFilledBinaryTree([]*TreeExampleElement{}),
			exist: false,
			want:  nil,
		},
		{
			name: "Maximum when only root",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
			}),
			exist: true,
			want:  &TreeExampleElement{key: 50},
		},
		{
			name: "Maximum when tree is full",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 75},
				{key: 100},
				{key: 30},
				{key: 15},
			}),
			exist: true,
			want:  &TreeExampleElement{key: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, isExist := tt.tree.Maximum()
			assert.Equalf(t, tt.want, got, "Maximum()")
			assert.Equalf(t, tt.exist, isExist, "Maximum()")
		})
	}
}

func TestBinaryTree_Remove(t *testing.T) {
	type args struct {
		key int
	}
	type testCase[T Rib] struct {
		name    string
		tree    *BinaryTree[T]
		args    args
		wantErr assert.ErrorAssertionFunc
		result  *BinaryTree[T]
		check   bool
	}
	tests := []testCase[*TreeExampleElement]{
		{
			name:    "Remove when empty",
			tree:    fetchFilledBinaryTree([]*TreeExampleElement{}),
			args:    args{key: 50},
			wantErr: assert.NoError,
			check:   false,
		},
		{
			name: "Remove when key does not exist",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 70},
				{key: 30},
				{key: 27},
				{key: 84},
				{key: 105},
			}),
			args:    args{key: 50},
			wantErr: assert.Error,
			check:   false,
		},
		{
			name: "Remove when only root",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
			}),
			result:  fetchFilledBinaryTree([]*TreeExampleElement{}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when when only right node",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
			}),
			result:  fetchFilledBinaryTree([]*TreeExampleElement{{key: 60}}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when when only right node",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
			}),
			result:  fetchFilledBinaryTree([]*TreeExampleElement{{key: 60}}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when when only left node",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 75},
				{key: 25},
				{key: 10},
				{key: 15},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 75},
				{key: 25},
				{key: 10},
				{key: 15},
			}),
			args:    args{key: 60},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when root has both nodes",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 10},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 10},
				{key: 60},
			}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when node has both nodes and there lefts != node left",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 20},
				{key: 13},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 20},
				{key: 60},
				{key: 10},
				{key: 5},
				{key: 15},
				{key: 12},
				{key: 13},
			}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when node has both nodes and there lefts != node left and not root",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 20},
				{key: 13},
				{key: 100},
				{key: 14},
				{key: 11},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 10},
				{key: 5},
				{key: 14},
				{key: 12},
				{key: 20},
				{key: 13},
				{key: 100},
				{key: 11},
			}),
			args:    args{key: 15},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when node has both nodes and there lefts == node left",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 15},
				{key: 60},
				{key: 10},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
			}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when node has both nodes and there rights == node right",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 55},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
				{key: 57},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 57},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
			}),
			args:    args{key: 55},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when node has both nodes and there rights != node right",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 40},
				{key: 80},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
				{key: 57},
				{key: 50},
				{key: 45},
				{key: 65},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 45},
				{key: 80},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
				{key: 57},
				{key: 50},
				{key: 65},
			}),
			args:    args{key: 40},
			wantErr: assert.NoError,
			check:   true,
		},
		{
			name: "Remove when node has both nodes and there rights != node right and not root",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 40},
				{key: 80},
				{key: 60},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
				{key: 57},
				{key: 50},
				{key: 45},
				{key: 65},
				{key: 62},
			}),
			result: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 40},
				{key: 80},
				{key: 62},
				{key: 10},
				{key: 15},
				{key: 5},
				{key: 12},
				{key: 13},
				{key: 70},
				{key: 57},
				{key: 50},
				{key: 45},
				{key: 65},
			}),
			args:    args{key: 60},
			wantErr: assert.NoError,
			check:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.tree.Remove(tt.args.key), fmt.Sprintf("Remove(%v)", tt.args.key))
			if tt.check {
				assert.Equal(t, tt.result, tt.tree)
			}
		})
	}
}

func TestBinaryTree_SymmetricTraversal(t *testing.T) {
	type testCase[T Rib] struct {
		name   string
		tree   *BinaryTree[T]
		result string
	}
	tests := []testCase[*TreeExampleElement]{
		{
			name: "empty tree",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{}),
		},
		{
			name: "valid symmetric traversal",
			tree: fetchFilledBinaryTree([]*TreeExampleElement{
				{key: 50},
				{key: 60},
				{key: 55},
				{key: 75},
				{key: 65},
				{key: 90},
				{key: 100},
				{key: 25},
				{key: 15},
				{key: 20},
				{key: 5},
				{key: 40},
			}),
			result: "5 15 20 25 40 50 55 60 65 75 90 100",
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		t.Run(tt.name, func(t *testing.T) {
			tt.tree.SymmetricTraversal(func(element *TreeExampleElement) {
				_, _ = buf.WriteString(fmt.Sprintf("%d ", element.Key()))
			})
			assert.Equal(t, tt.result, strings.TrimRight(buf.String(), " "))
		})
	}
}
