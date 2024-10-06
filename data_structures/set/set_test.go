package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_Add(t *testing.T) {
	type testCase[T int] struct {
		name     string
		s        *Set[int]
		arg      int
		want     bool
		checkSet *Set[int]
	}
	tests := []testCase[int]{
		{
			name:     "add when empty",
			s:        &Set[int]{},
			arg:      5,
			want:     true,
			checkSet: &Set[int]{5: struct{}{}},
		},
		{
			name:     "add when exist",
			s:        &Set[int]{3: struct{}{}, 5: struct{}{}},
			arg:      5,
			want:     false,
			checkSet: &Set[int]{3: struct{}{}, 5: struct{}{}},
		},
		{
			name:     "add when not empty",
			s:        &Set[int]{3: struct{}{}, 7: struct{}{}},
			arg:      5,
			want:     true,
			checkSet: &Set[int]{3: struct{}{}, 7: struct{}{}, 5: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Add(tt.arg); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
				assert.ElementsMatch(t, keys(tt.checkSet), keys(tt.s))
			}
		})
	}
}

func TestSet_RemoveAll(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		s        *Set[T]
		checkSet *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "remove when empty",
			s:    &Set[int]{},
		},
		{
			name: "remove when full",
			s:    &Set[int]{3: struct{}{}, 5: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.RemoveAll()
			assert.Empty(t, keys(tt.s))
		})
	}
}

func keys(s *Set[int]) []int {
	var collection []int
	for v := range *s {
		collection = append(collection, v)
	}
	return collection
}

func TestNewSet(t *testing.T) {
	type args[T comparable] struct {
		values []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "new set",
			args: args[int]{
				values: []int{1, 2, 3},
			},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSet(tt.args.values...), "NewSet(%v)", tt.args.values)
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "contains when empty",
			s:    &Set[int]{},
			args: args[int]{value: 2},
			want: false,
		},
		{
			name: "contains when full",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{value: 2},
			want: true,
		},
		{
			name: "not contains when full",
			s:    &Set[int]{1: struct{}{}, 4: struct{}{}, 3: struct{}{}},
			args: args[int]{value: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Contains(tt.args.value), "Contains(%v)", tt.args.value)
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		args *Set[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "difference when empty source",
			s:    &Set[int]{},
			args: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{},
		},
		{
			name: "difference when empty diff",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "success difference",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{1: struct{}{}},
		},
		{
			name: "all elements are different",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Difference(tt.args), "Difference(%v)", tt.args)
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		args *Set[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "difference when empty source",
			s:    &Set[int]{},
			args: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "difference when empty diff",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "success difference",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{1: struct{}{}, 4: struct{}{}},
		},
		{
			name: "all elements are different",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.SymmetricDifference(tt.args), "SymmetricDifference(%v)", tt.args)
		})
	}
}

func TestSet_Intersect(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		args *Set[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "difference when empty source",
			s:    &Set[int]{},
			args: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{},
		},
		{
			name: "difference when empty diff",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{},
			want: &Set[int]{},
		},
		{
			name: "success intersect",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "all elements are different",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
			want: &Set[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Intersect(tt.args), "Intersect(%v)", tt.args)
		})
	}
}

func TestSet_Union(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		args *Set[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "union when empty source",
			s:    &Set[int]{},
			args: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "union when empty diff",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "union with doubles",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}},
		},
		{
			name: "union when all elements are different",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 6: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Union(tt.args), "Union(%v)", tt.args)
		})
	}
}

func TestSet_IsSubset(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		args *Set[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "subset when empty source",
			s:    &Set[int]{},
			args: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "subset when empty diff",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{},
		},
		{
			name: "no subset",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{4: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "union when all elements are different",
			s:    &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.IsSubset(tt.args), "IsSubset(%v)", tt.args)
		})
	}
}
