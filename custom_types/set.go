package custom_types

import (
	"bytes"
	"fmt"
	"searchAlgorithms"
	"sortAlgorithms"
)

type Set[T searchAlgorithms.Comparable[T]] []T

func (set Set[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	for index, value := range set {
		buffer.WriteString(fmt.Sprintf("%v", value))
		if index != len(set)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func NewSet[T searchAlgorithms.Comparable[T]](values ...T) Set[T] {
	var set Set[T]
	for _, value := range values {
		if !set.Contains(value) {
			set = append(set, value)
		}
	}
	return set
}

func (set Set[T]) Search(key T) int {
	return searchAlgorithms.BinarySearch[T]{Items: set}.Execute(key)
}

func (set Set[T]) Sort() {
	sortAlgorithms.InsertionSort(set)
}

func (set Set[T]) Contains(searchItem T) bool {
	for _, value := range set {
		if value.CompareTo(searchItem) == 0 {
			return true
		}
	}
	return false
}
