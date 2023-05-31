package custom_types

import "search"

type Array[T search.Comparable] []T

func NewArray[T search.Comparable](values ...T) Array[T] {
	var array Array[T]
	for _, value := range values {
		array = append(array, value)
	}
	return array
}

func (array Array[T]) Search(key T) int {
	return search.BinarySearch[T]{Items: array}.FindIndex(key)
}
