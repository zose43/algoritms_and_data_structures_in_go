package hash_table

import (
	"github.com/stretchr/testify/assert"
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
