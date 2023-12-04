package utils

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T])
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func (set Set[T]) Add(values ...T) {
	for _, value := range values {
		set[value] = struct{}{}
	}
}

func (set Set[T]) Remove(value T) {
	delete(set, value)
}

func (set Set[T]) Contains(value T) bool {
	_, ok := set[value]
	return ok
}

func (set Set[T]) Values() []T {
	res := make([]T, 0, len(set))
	for value := range set {
		res = append(res, value)
	}
	return res
}

func (set Set[T]) IsEmpty() bool {
	return len(set) == 0
}

func (set Set[T]) Len() int {
	return len(set)
}

func (set Set[T]) Clear() {
	clear(set)
}

func (s Set[E]) String() string {
	return fmt.Sprintf("%v", s.Values())
}

func (set Set[T]) Equal(other Set[T]) bool {
	if set.Len() != other.Len() {
		return false
	}
	for value := range set {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (set Set[T]) Union(other Set[T]) Set[T] {
	res := NewSet(set.Values()...)
	res.Add(other.Values()...)
	return res
}

func (set Set[T]) Intersect(other Set[T]) Set[T] {
	if set.Len() > other.Len() {
		other.Intersect(set)
	}
	res := NewSet[T]()
	for value := range set {
		if other.Contains(value) {
			res.Add(value)
		}
	}
	return res
}
