package cocktail_sort

import "algoritms_and_structures/sorts/bubble_sort"

func CocktailSort[T any](src []T, fn bubble_sort.Compare[T]) []T {
	if len(src) < 2 {
		return src
	}
	swap := func(i, j int) {
		src[i], src[j] = src[j], src[i]
	}
	left := 0
	right := len(src) - 1
	for left <= right {
		for i := left; i < right; i++ {
			if fn(src[i], src[i+1]) {
				swap(i, i+1)
			}
		}
		left++
		for j := right; j > left; j-- {
			if fn(src[j-1], src[j]) {
				swap(j-1, j)
			}
		}
		right--
	}
	return src
}
