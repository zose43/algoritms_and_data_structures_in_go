package set

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(value T) bool {
	oldLen := len(*s)
	(*s)[value] = struct{}{}
	return oldLen == len(*s)
}

func (s *Set[T]) IsEmpty() bool {
	return len(*s) > 0
}

func (s *Set[T]) Size() int {
	return len(*s)
}

func (s *Set[T]) RemoveAll() {
	*s = Set[T]{}
}

func (s *Set[T]) Contains(value T) bool {
	if _, ok := (*s)[value]; ok {
		return true
	}
	return false
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	diffSet := NewSet[T]()
	for value := range *s {
		if !other.Contains(value) {
			diffSet.Add(value)
		}
	}
	return diffSet
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	diffSymSet := NewSet[T]()
	for value := range *s {
		if !other.Contains(value) {
			diffSymSet.Add(value)
		}
	}
	for value := range *other {
		if !s.Contains(value) {
			diffSymSet.Add(value)
		}
	}
	return diffSymSet
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersectSet := NewSet[T]()
	if s.Size() < other.Size() {
		for value := range *s {
			if other.Contains(value) {
				intersectSet.Add(value)
			}
		}
	} else {
		for value := range *other {
			if s.Contains(value) {
				intersectSet.Add(value)
			}
		}
	}
	return intersectSet
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := NewSet[T]()
	for value := range *s {
		unionSet.Add(value)
	}
	for value := range *other {
		unionSet.Add(value)
	}
	return unionSet
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.Size() < other.Size() {
		return false
	}
	for value := range *s {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func NewSet[T comparable](values ...T) *Set[T] {
	s := make(Set[T])
	for _, value := range values {
		s.Add(value)
	}
	return &s
}
