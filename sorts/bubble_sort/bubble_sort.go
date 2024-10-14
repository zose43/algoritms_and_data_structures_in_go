package bubble_sort

func BubbleSortInt(src []int, isAscent bool) []int {
	if len(src) < 2 {
		return src
	}
	swap := func(i, j int) {
		src[i], src[j] = src[j], src[i]
	}
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			if isAscent {
				if src[i] > src[j] {
					swap(i, j)
				}
			} else {
				if src[i] < src[j] {
					swap(i, j)
				}
			}
		}
	}
	return src
}

type Compare[T any] func(T, T) bool

func BubbleSort[T any](src []T, fn Compare[T]) []T {
	if len(src) < 2 {
		return src
	}
	swap := func(i, j int) {
		src[i], src[j] = src[j], src[i]
	}
	for i := 0; i+1 < len(src); i++ {
		for j := 0; j+1 < len(src)-i; j++ {
			if fn(src[j], src[j+1]) {
				swap(j, j+1)
			}
		}
	}
	return src
}
