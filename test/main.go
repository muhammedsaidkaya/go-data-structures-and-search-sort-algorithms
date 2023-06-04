package main

import (
	"custom_types"
	"fmt"
	"searchAlgorithms"
	"sortAlgorithms"
)

func main() {

	peopleList := []Person{
		{ID: 2}, {ID: 1}, {ID: 2}, {ID: 4}, {ID: 3},
	}

	//Custom Type
	peopleSet := custom_types.NewSet(peopleList...)
	fmt.Println(peopleSet)

	//Custom Type Sorting
	peopleSet.Sort()
	fmt.Println(peopleSet)

	//Custom Type Searching
	resultIndex := peopleSet.Search(Person{ID: 3})
	fmt.Println(resultIndex)

	//---------------
	//Binary Search
	resultIndex = searchAlgorithms.NewBinarySearch(peopleList).Execute(Person{ID: 3})
	fmt.Println(resultIndex)

	resultIndex = searchAlgorithms.BinarySearch[Person]{Items: peopleList}.Execute(Person{ID: 3})
	fmt.Println(resultIndex)

	//---------------
	//Insertion Sort
	sortAlgorithms.Sort(sortAlgorithms.INSERTION, peopleList)
	fmt.Println(peopleList)
	sortAlgorithms.Shuffle(peopleList)

	//Selection Sort
	sortAlgorithms.Sort(sortAlgorithms.SELECTION, peopleList)
	fmt.Println(peopleList)
	sortAlgorithms.Shuffle(peopleList)

	//Merge Sort
	sortAlgorithms.Sort(sortAlgorithms.MERGE, peopleList)
	fmt.Println(peopleList)
	sortAlgorithms.Shuffle(peopleList)

	//Quick Sort
	sortAlgorithms.Sort(sortAlgorithms.QUICK, peopleList)
	fmt.Println(peopleList)

	//BST
	bst := custom_types.BST[Person]{
		Root: nil,
	}
	bst.Put(Person{ID: 3})
	bst.Put(Person{ID: 2})
	bst.Put(Person{ID: 4})
	fmt.Println(bst.InOrder(bst.Root, []Person{}))
	result := bst.Get(Person{ID: 2})
	if result != nil {
		fmt.Println(result.Value)
	}
	bst.Delete(Person{ID: 2})
	result = bst.Get(Person{ID: 2})
	fmt.Println(result)

	queue := custom_types.NewQueue[Person]()
	fmt.Println(queue)
	queue.Enqueue(Person{ID: 1})
	queue.Enqueue(Person{ID: 2})
	fmt.Println(queue)
	item, _ := queue.Dequeue()
	fmt.Println(item)
	fmt.Println(queue)

	stack := custom_types.NewStack[Person](2)
	fmt.Println(stack)
	stack.Push(Person{ID: 1})
	fmt.Println(stack)
	stack.Push(Person{ID: 2})
	fmt.Println(stack)
	err := stack.Push(Person{ID: 3})
	fmt.Println(err)
	value, _ := stack.Pop()
	fmt.Println(value)
	fmt.Println(stack)
	fmt.Println(stack.Tail)
	stack.Push(Person{ID: 3})
	fmt.Println(stack)

	fmt.Println("\nLINKED LIST")
	ll := custom_types.NewLinkedList[Person](&custom_types.LinkedListOptions{
		SortOptions:         &custom_types.SortOptions{Sorted: true, Desc: false},
		MultipleOccurrences: true,
		Reversed:            true,
	}).Data(peopleList)
	//ll := custom_types.NewSortedLinkedList[Person]()
	fmt.Println(ll.IsCircular())
	fmt.Println(ll)
	fmt.Println(ll.Length())
	ll.Remove(Person{ID: 4})
	ll.Reverse()
	ll.RemoveByIndex(2)
	fmt.Println(ll)
	item, _ = ll.GetByIndex(1)
	fmt.Println(item)
	fmt.Println(ll.GetSortedList())

	fmt.Println(ll.GetLastNthItem(1))

	ll.Insert(Person{ID: 3})
	ll.Insert(Person{ID: 5})
	ll.Insert(Person{ID: 6})
	ll.Insert(Person{ID: 6})
	fmt.Println(ll)
	item, _ = ll.GetMiddleItem()
	fmt.Println(item)
	fmt.Println(ll.Count(Person{ID: 6}))
	ll.RemoveDuplicates()
	fmt.Println(ll)
	ll.Shuffle()
	ll.Sort()
	fmt.Println(ll)

	lt := custom_types.NewSortedLinkedList[Person]().Data([]Person{
		{ID: 2},
		{ID: 5},
	})
	fmt.Println(lt)
	fmt.Println(ll.Intersection(lt))
}

type Person struct {
	ID int
}

func (_p Person) CompareTo(other Person) int {
	if _p.ID < other.ID {
		return -1
	} else if _p.ID == other.ID {
		return 0
	} else {
		return 1
	}
}
