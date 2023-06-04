package custom_types

import (
	"bytes"
	"errors"
	"fmt"
	"searchAlgorithms"
	"sortAlgorithms"
)

type SortOptions struct {
	Sorted bool
	Desc   bool
}

type LinkedListOptions struct {
	SortOptions         *SortOptions
	Reversed            bool
	MultipleOccurrences bool
}

type LinkedList[T searchAlgorithms.Comparable[T]] struct {
	Head    *Item[T]
	Options *LinkedListOptions
}

func (linkedList LinkedList[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	var p *Item[T] = linkedList.Head
	for p != nil {
		buffer.WriteString(fmt.Sprintf("%v", p.Value))
		if p.Next != nil {
			buffer.WriteString(",")
		}
		p = p.Next
	}
	buffer.WriteString(")")
	return buffer.String()
}

func (linkedList *LinkedList[T]) Reversed() *LinkedList[T] {
	linkedList.Options.Reversed = true
	return linkedList
}

func (linkedList *LinkedList[T]) Sorted() *LinkedList[T] {
	linkedList.Options.SortOptions.Sorted = true
	return linkedList
}

func (linkedList *LinkedList[T]) Data(values []T) *LinkedList[T] {
	for _, value := range values {
		if err := linkedList.Insert(value); err != nil {
			panic(err)
		}
	}
	return linkedList
}

func NewSortedLinkedList[T searchAlgorithms.Comparable[T]]() *LinkedList[T] {
	return NewLinkedList[T](&LinkedListOptions{
		SortOptions: &SortOptions{},
	}).Sorted()
}

func NewLinkedList[T searchAlgorithms.Comparable[T]](options *LinkedListOptions) *LinkedList[T] {
	return &LinkedList[T]{
		Head:    nil,
		Options: options,
	}
}

func (linkedList LinkedList[T]) IsCircular() bool {
	if linkedList.Head == nil {
		return true
	}
	var p *Item[T] = linkedList.Head.Next

	for p != nil && p != linkedList.Head {
		p = p.Next
	}
	return p == linkedList.Head
}

func (linkedList *LinkedList[T]) Sort() {
	linkedList.Options.SortOptions.Sorted = true
	list := linkedList.GetSortedList()
	linkedList.RemoveAll()
	linkedList.Data(list)
}

func (linkedList LinkedList[T]) Intersection(other *LinkedList[T]) (*LinkedList[T], error) {
	if !linkedList.Options.SortOptions.Sorted || !other.Options.SortOptions.Sorted {
		return nil, errors.New("Sort options must be set...")
	}
	return &LinkedList[T]{
		Head: linkedList.intersectionHelper(linkedList.Head, other.Head),
	}, nil
}

func (linkedList LinkedList[T]) intersectionHelper(first, second *Item[T]) *Item[T] {
	if first == nil || second == nil {
		return nil
	}
	if first.Value.CompareTo(second.Value) < 0 {
		return linkedList.intersectionHelper(first.Next, second)
	}

	if first.Value.CompareTo(second.Value) > 0 {
		return linkedList.intersectionHelper(first, second.Next)
	}
	return &Item[T]{
		Value: first.Value,
		Next:  linkedList.intersectionHelper(first.Next, second.Next),
	}
}

func (linkedList LinkedList[T]) Count(value T) uint64 {
	if linkedList.Head == nil {
		return 0
	}
	var p *Item[T] = linkedList.Head
	var counter uint64
	for p != nil {
		if p.Value.CompareTo(value) == 0 {
			counter++
		}
		p = p.Next
	}
	return counter
}

func (linkedList LinkedList[T]) GetMiddleItem() (any, error) {
	if linkedList.Head == nil {
		return nil, errors.New("Linkedlist is empty...")
	}
	var p, q *Item[T] = linkedList.Head, linkedList.Head
	for p != nil {
		if q.Next != nil && q.Next.Next != nil {
			q = q.Next.Next
		} else {
			break
		}
		p = p.Next
	}
	return p.Value, nil
}

func (linkedList *LinkedList[T]) Insert(value T) error {
	var p, q, r *Item[T] = linkedList.Head, nil, nil
	//Sorted gez
	if linkedList.Options.SortOptions.Sorted {
		for p != nil && (linkedList.Options.Reversed && p.Value.CompareTo(value) > 0 || !linkedList.Options.Reversed && p.Value.CompareTo(value) < 0) {
			q = p
			p = p.Next
		}
	} else {
		for p != nil && p.Value.CompareTo(value) != 0 {
			q = p
			p = p.Next
		}
	}
	//Aynı item
	if !linkedList.Options.MultipleOccurrences {
		if p != nil && p.Value.CompareTo(value) == 0 {
			return errors.New("There is already item")
		}
	}
	r = &Item[T]{
		Value: value,
		Next:  nil,
	}
	//En başta
	if q == nil {
		linkedList.Head = r
		r.Next = p
	} else {
		//Ortada veya sonda ise
		q.Next = r
		r.Next = p
	}
	return nil
}

func (linkedList *LinkedList[T]) Shuffle() error {
	if linkedList.Head == nil {
		return errors.New("Linkedlist is empty...")
	}
	list := linkedList.GetSortedList()
	sortAlgorithms.Shuffle(list)
	linkedList.RemoveAll()
	linkedList.Options.SortOptions.Sorted = false
	linkedList.Options.Reversed = false
	linkedList.Data(list)
	return nil
}

func (linkedList LinkedList[T]) Length() uint64 {
	if linkedList.Head == nil {
		return 0
	}
	var p *Item[T] = linkedList.Head
	var length uint64
	for p != nil {
		p = p.Next
		length++
	}
	return length
}

func (linkedList *LinkedList[T]) Reverse() {
	var prev, curr, next *Item[T] = nil, linkedList.Head, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	linkedList.Head = prev
	linkedList.Options.Reversed = !linkedList.Options.Reversed
}

func (linkedList LinkedList[T]) Get(value T) (any, error) {
	if linkedList.Head == nil {
		return nil, errors.New("There is no item...")
	}
	var p *Item[T] = linkedList.Head
	for p != nil {
		if p.Value.CompareTo(value) == 0 {
			break
		}
		p = p.Next
	}
	if p == nil {
		return nil, errors.New("Item not found...")
	}
	return p.Value, nil
}

func (linkedList LinkedList[T]) GetLastNthItem(value uint64) (any, error) {
	if linkedList.Head == nil {
		return nil, errors.New("Linkedlist is empty...")
	}
	var counter uint64
	var p, q *Item[T] = linkedList.Head, linkedList.Head
	for p != nil {
		if counter == value {
			break
		}
		p = p.Next
		counter++
	}
	if p == nil {
		return nil, errors.New("Index out of length...")
	}
	for p.Next != nil {
		q = q.Next
		p = p.Next
	}
	return q.Value, nil
}

func (linkedList LinkedList[T]) GetSortedList() []T {
	var list []T
	if linkedList.Head == nil {
		return list
	}
	var p *Item[T] = linkedList.Head
	for p != nil {
		list = append(list, p.Value)
		p = p.Next
	}
	sortAlgorithms.MergeSort(list)
	return list
}

func (linkedList LinkedList[T]) GetByIndex(value uint64) (any, error) {
	if linkedList.Head == nil {
		return nil, errors.New("Linkedlist is empty...")
	}
	var p *Item[T] = linkedList.Head
	var counter uint64
	for p != nil {
		if counter == value {
			break
		}
		p = p.Next
		counter++
	}
	if p == nil {
		return nil, errors.New("Index out of length...")
	}
	return p.Value, nil
}

func (linkedList *LinkedList[T]) Remove(value T) error {
	var p, q *Item[T] = linkedList.Head, nil
	if p == nil {
		return errors.New("Linkedlist is empty...")
	}
	//Sorted gez
	for p != nil {
		if p.Value.CompareTo(value) == 0 {
			if linkedList.Options.MultipleOccurrences {
				if q == nil {
					linkedList.Head = p.Next
				} else {
					q.Next = p.Next
				}
				p = p.Next
				continue
			} else {
				break
			}
		}
		q = p
		p = p.Next
	}
	//Eğer item yoksa
	if p == nil {
		return errors.New("Item not found...")
	}
	//En başta ise
	if q == nil {
		linkedList.Head = p.Next
	} else {
		//Ortada veya sonda ise
		q.Next = p.Next
	}
	return nil
}

func (linkedList *LinkedList[T]) RemoveAll() error {
	if linkedList.Head == nil {
		return errors.New("Linked is empty")
	}
	linkedList.Head = nil
	return nil
}

func (linkedList *LinkedList[T]) RemoveDuplicates() error {
	if !linkedList.Options.SortOptions.Sorted {
		return errors.New("sorted option must be set to remove duplicated items")
	}
	var p *Item[T] = linkedList.Head
	for p != nil {
		if p.Next != nil && p.Value.CompareTo(p.Next.Value) == 0 {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return nil
}

func (linkedList *LinkedList[T]) RemoveByIndex(value uint64) error {
	var p, q *Item[T] = linkedList.Head, nil
	if p == nil {
		return errors.New("Linkedlist is empty...")
	}
	var counter uint64
	//Sorted gez
	for p != nil {
		if value == counter {
			break
		}
		q = p
		p = p.Next
		counter++
	}
	//Eğer item yoksa
	if p == nil {
		return errors.New("Index out of length...")
	}
	//Index 0 ise
	if q == nil {
		linkedList.Head = p.Next
	} else {
		//Ortada veya sonda ise
		q.Next = p.Next
	}
	return nil
}
