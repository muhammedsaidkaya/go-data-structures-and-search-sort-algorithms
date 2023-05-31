package main

import (
	"custom_types"
	"fmt"
)

type Person struct {
	ID int
}

func (_p Person) CompareTo(other Person) int {
	if _p.ID < other.ID {
		return 1
	} else if _p.ID == other.ID {
		return 0
	} else {
		return -1
	}
}

type ID string

func (id ID) CompareTo(other ID) int {
	if id < other {
		return 1
	} else if id == other {
		return 0
	} else {
		return -1
	}
}

func main() {

	//Binary Search
	//items1 := []int{1, 2, 3, 4, 5, 6}
	//bs := search.BinarySearch[int]{Items: items1}
	//resultIndex := bs.FindIndex(2)
	//fmt.Println(resultIndex)
	//
	//items2 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	//as := search.BinarySearch[float32]{Items: items2}
	//resultIndex = as.FindIndex(3.0)
	//fmt.Println(resultIndex)
	//
	//items3 := []string{"a", "b", "c", "d", "e", "f"}
	//ds := search.BinarySearch[string]{Items: items3}
	//resultIndex = ds.FindIndex("e")
	//fmt.Println(resultIndex)
	//
	//items4 := []int{1, 2, 3, 4, 5, 6}
	//resultIndex = search.NewBinarySearch[int](items4).FindIndex(2)
	//fmt.Println(resultIndex)

	//Custom Types
	people := custom_types.NewArray(Person{ID: 1}, Person{ID: 2}, Person{ID: 3})
	resultIndex := people.Search(Person{ID: 2})
	fmt.Println(resultIndex)

	strArray := custom_types.NewArray(ID("a"), ID("b"), ID("c"))
	resultIndex = strArray.Search("a")
	fmt.Println(resultIndex)

}
