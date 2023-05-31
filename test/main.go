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
