package comb_sort

import "algoritms_and_structures/sorts/bubble_sort"

func CombSort[T any](src []T, fn bubble_sort.Compare[T]) []T {
	if len(src) < 2 {
		return src
	}
	swap := func(i, j int) {
		src[i], src[j] = src[j], src[i]
	}
	const factor = 1.25
	step := float64(len(src) - 1)
	for step >= 1 {
		for i := 0; float64(i)+step < float64(len(src)); i++ {
			if fn(src[i], src[i+int(step)]) {
				swap(i, i+int(step))
			}
		}
		step /= factor
	}
	return src
}
