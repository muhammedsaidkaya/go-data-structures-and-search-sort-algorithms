package main

import (
	"custom_types"
	"fmt"
	"search"
)

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

type ID string

func (id ID) CompareTo(other ID) int {
	if id < other {
		return -1
	} else if id == other {
		return 0
	} else {
		return 1
	}
}

func main() {

	//Custom Types + Binary Search
	people := custom_types.NewArray(Person{ID: 1}, Person{ID: 2}, Person{ID: 3})
	resultIndex := people.Search(Person{ID: 3})
	fmt.Println(resultIndex)

	strArray := custom_types.NewArray(ID("a"), ID("b"), ID("c"))
	resultIndex = strArray.Search(ID("a"))
	fmt.Println(resultIndex)

	resultIndex = search.NewBinarySearch(strArray).Execute(ID("b"))
	fmt.Println(resultIndex)
}
