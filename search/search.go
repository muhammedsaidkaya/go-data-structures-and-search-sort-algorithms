package search

type Comparable[T any] interface {
	CompareTo(other T) int
}

type BinarySearch[T Comparable[T]] struct {
	Items []T
}

func NewBinarySearch[T Comparable[T]](items []T) BinarySearch[T] {
	return BinarySearch[T]{Items: items}
}

func (_b BinarySearch[T]) Execute(key T) int {
	return _b.helper(key, 0, len(_b.Items)-1)
}

func (_b BinarySearch[T]) helper(key T, lo, hi int) int {
	if lo <= hi {
		mid := (lo + hi) / 2
		cmp := _b.Items[mid].CompareTo(key)
		if cmp < 0 {
			return _b.helper(key, lo, mid-1)
		} else if cmp > 0 {
			return _b.helper(key, mid+1, hi)
		} else {
			return mid
		}
	}
	return -1
}
