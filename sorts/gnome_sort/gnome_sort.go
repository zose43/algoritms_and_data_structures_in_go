package gnome_sort

import "algoritms_and_structures/sorts/bubble_sort"

func GnomeSort[T any](src []T, compare bubble_sort.Compare[T]) []T {
	if len(src) < 2 {
		return src
	}
	swap := func(i, j int) {
		src[i], src[j] = src[j], src[i]
	}
	i := 1
	for i < len(src) {
		if compare(src[i], src[i-1]) {
			i++
		} else {
			swap(i, i-1)
			if i > 1 {
				i--
			}
		}
	}
	return src
}

func gnomeSort[T any](src []T, compare bubble_sort.Compare[T]) []T {
	if len(src) < 2 {
		return src
	}
	swap := func(i, j int) {
		src[i], src[j] = src[j], src[i]
	}
	count := len(src) - 1
	for i := 0; i < count; i++ {
		if compare(src[i], src[i+1]) {
			swap(i, i+1)
			if i != 0 {
				for j := i; j > 0 && compare(src[j-1], src[j]); j-- {
					swap(j-1, j)
				}
			}
		}
	}
	return src
}
