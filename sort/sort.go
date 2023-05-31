package sortAlgorithms

import (
	"math/rand"
	"searchAlgorithms"
)

const INSERTION = "insertion"
const SELECTION = "selection"
const MERGE = "merge"
const QUICK = "quick"

func Sort[T searchAlgorithms.Comparable[T]](algorithm string, list []T) {
	if algorithm == INSERTION {
		InsertionSort(list)
	} else if algorithm == SELECTION {
		SelectionSort(list)
	} else if algorithm == MERGE {
		MergeSort(list)
	} else if algorithm == QUICK {
		QuickSort(list)
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

func SelectionSort[T searchAlgorithms.Comparable[T]](list []T) {
	length := len(list)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if list[j].CompareTo(list[min]) < 0 {
				min = j
			}
		}
		exchange(list, i, min)
	}
}

func MergeSort[T searchAlgorithms.Comparable[T]](list []T) {
	if len(list) > 1 {
		mid := len(list) / 2
		left := make([]T, mid)
		right := make([]T, len(list)-mid)
		for i := 0; i < mid; i++ {
			left[i] = list[i]
		}
		for i := mid; i < len(list); i++ {
			right[i-mid] = list[i]
		}
		MergeSort(left)
		MergeSort(right)
		mergeSortHelper(list, left, right, mid, len(list)-mid)
	}
}

func mergeSortHelper[T searchAlgorithms.Comparable[T]](list, left, right []T, leftIndex, rightIndex int) {
	var i, j, k int
	for i < leftIndex || j < rightIndex {
		if i >= leftIndex {
			list[k] = right[j]
			j++
		} else if j >= rightIndex {
			list[k] = left[i]
			i++
		} else if left[i].CompareTo(right[j]) < 0 {
			list[k] = left[i]
			i++
		} else {
			list[k] = right[j]
			j++
		}
		k++
	}
}

func QuickSort[T searchAlgorithms.Comparable[T]](list []T) {
	Shuffle(list)
	quickSortHelper(list, 0, len(list)-1)
}

func Shuffle[T searchAlgorithms.Comparable[T]](list []T) {
	for i := range list {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

func quickSortHelper[T searchAlgorithms.Comparable[T]](list []T, low, high int) {
	if low < high {
		pivot := quickSortHelperPartition(list, low, high)
		quickSortHelper(list, low, pivot-1)
		quickSortHelper(list, pivot+1, high)
	}
}

func quickSortHelperPartition[T searchAlgorithms.Comparable[T]](list []T, low, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j].CompareTo(pivot) < 0 {
			i++
			exchange(list, i, j)
		}
	}
	exchange(list, i+1, high)
	return i + 1
}
