package sortAlgorithms

import (
	"searchAlgorithms"
)

const INSERTION = "insertion"

func Sort[T searchAlgorithms.Comparable[T]](algorithm string, list []T) {
	if algorithm == INSERTION {
		InsertionSort(list)
	}
}

func exchange[T searchAlgorithms.Comparable[T]](list []T, i, j int) {
	item := list[i]
	list[i] = list[j]
	list[j] = item
}

func InsertionSort[T searchAlgorithms.Comparable[T]](list []T) {
	length := len(list)
	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if list[j].CompareTo(list[j-1]) < 0 {
				exchange(list, j, j-1)
			} else {
				break
			}
		}
	}
}
