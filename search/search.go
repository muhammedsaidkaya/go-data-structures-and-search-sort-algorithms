package search

type Search interface {
	Execute() int
}

type Comparable interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type BinarySearch[T Comparable] struct {
	Items []T
}

func NewBinarySearch[T Comparable](items []T) BinarySearch[T] {
	return BinarySearch[T]{Items: items}
}

func (_b BinarySearch[T]) FindIndex(key T) int {
	return _b.helper(key, 0, len(_b.Items)-1)
}

func (_b BinarySearch[T]) helper(key T, lo, hi int) int {
	if lo <= hi {
		mid := (lo + hi) / 2
		if key < _b.Items[mid] {
			return _b.helper(key, lo, mid-1)
		} else if key > _b.Items[mid] {
			return _b.helper(key, mid+1, hi)
		} else {
			return mid
		}
	}
	return -1
}
