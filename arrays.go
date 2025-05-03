package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers[5] int // Here, we are declaring a blank array, so it will be initialized with 0

	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	fruits := [4]string{"Apples","Banana","Oranges","Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third Element:", fruits[2])

	orignalArray := [3]int{1,2,3}
	copiedArray := orignalArray

	copiedArray[0] = 100

	fmt.Println(copiedArray)
	fmt.Println(orignalArray)

	for i := 0; i<len(numbers); i++ {
		fmt.Println("Element at index,",i,':', numbers[i])
	}
}