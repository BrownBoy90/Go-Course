package basics

import (
	"fmt"
	"slices"
)

func main() {
	// var sliceName[]ElementType

	// var numbers []int // it declares a numbers variable and it's going to contain a slice, a series of numbers
	// 		// Here an underlying array is not connected

	// // Slice means:- it is kind of array but does not have a fixed length

	// var numbers1 = []int // here an uunderlying array is connected which has 0 size

	// var numbers2 = []int {9,8,7}  

	//SLICES - References to underlying arrays, they do not store any data themseleves but provide a window into the array's elements, slices can grow and shrink dynamically

	// function len to check the length of the slice

	// function cap to check the capacity of the slice, it will check the number of elements in the underlying array, starting from the slices first element.

	// To initialize slice using make
	// slice := make([]int,5)			// Make stateent is only useful when we already know the size of the underlying array

		// Here What we are doing is we are initiating a slice with a length of five, a capacity of five

	a := [5]int{1,2,3,4,5}
	slice1:=a[2:4]
	// fmt.Println(cap(slice1))
	// fmt.Println(slice1)

	slice1 = append(slice1, 6,7)
	fmt.Println("Slice1:",slice1)

	// Copying a slice
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("SliceCopy:",sliceCopy)

	// var nilSlice []int

	for i, v := range slice1{
		fmt.Println(i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}

	twoD := make([][]int, 3)
	for i := 0; i<3; i++ {
		innerLen := i +1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i+j
		}
	}

	fmt.Println(twoD)

}