package gnome_sort

import (
	"math/rand"
	"testing"
	"time"
)

func compareFn(i, j int) bool {
	return i > j
}

func fetchFixedSlice(n int) []int {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return gen.Perm(n)
}

func gnomeSortCustomXInt(b *testing.B, x int) {
	s := fetchFixedSlice(x)
	b.ResetTimer()
	gnomeSort(s, compareFn)
}

func gnomeSortXInt(b *testing.B, x int) {
	s := fetchFixedSlice(x)
	b.ResetTimer()
	GnomeSort(s, compareFn)
}

func BenchmarkGnomeSortCustom1000(b *testing.B) {
	gnomeSortCustomXInt(b, 1000)
}

func BenchmarkGnomeSortCustom10000(b *testing.B) {
	gnomeSortCustomXInt(b, 10000)
}

func BenchmarkGnomeSortCustom100000(b *testing.B) {
	gnomeSortCustomXInt(b, 100000)
}

func BenchmarkGnomeSort1000(b *testing.B) {
	gnomeSortXInt(b, 1000)
}

func BenchmarkGnomeSort10000(b *testing.B) {
	gnomeSortXInt(b, 10000)
}

func BenchmarkGnomeSort100000(b *testing.B) {
	gnomeSortXInt(b, 100000)
}
