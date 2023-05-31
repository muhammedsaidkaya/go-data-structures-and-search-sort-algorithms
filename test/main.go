package main

import (
	"custom_types"
	"fmt"
	"search"
)

func main() {

	//Binary Search
	items1 := []int{1, 2, 3, 4, 5, 6}
	bs := search.BinarySearch[int]{Items: items1}
	resultIndex := bs.FindIndex(2)
	fmt.Println(resultIndex)

	items2 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	as := search.BinarySearch[float32]{Items: items2}
	resultIndex = as.FindIndex(3.0)
	fmt.Println(resultIndex)

	items3 := []string{"a", "b", "c", "d", "e", "f"}
	ds := search.BinarySearch[string]{Items: items3}
	resultIndex = ds.FindIndex("e")
	fmt.Println(resultIndex)

	items4 := []int{1, 2, 3, 4, 5, 6}
	resultIndex = search.NewBinarySearch[int](items4).FindIndex(2)
	fmt.Println(resultIndex)

	//Custom Types
	intArray := custom_types.NewArray(1, 2, 3)
	resultIndex = intArray.Search(3)
	fmt.Println(resultIndex)

	strArray := custom_types.NewArray("a", "b", "c")
	resultIndex = strArray.Search("a")
	fmt.Println(resultIndex)

}
