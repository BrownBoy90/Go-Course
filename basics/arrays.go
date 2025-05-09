package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int // Here, we are declaring a blank array, so it will be initialized with 0

	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	fruits := [4]string{"Apples", "Banana", "Oranges", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third Element:", fruits[2])

	orignalArray := [3]int{1, 2, 3}
	copiedArray := orignalArray

	copiedArray[0] = 100

	fmt.Println(copiedArray)
	fmt.Println(orignalArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for index, value := range numbers { // Instead of index and value we can use for i, v
		// If we want to discard any of the returns provided buy range we can use _ for e.g.

		// This _ in golang is known as blank Identifier

		// for _, v:= range numbers {
		// 		fmt.Println("Value:", value)
		// }
		fmt.Println("Index:", index, "Value:", value)
	}

	// a, _ := someFunction()
	// fmt.Println(a)
	// fmt.Println(b)

	// Underscore usecase
	b := 2
	_ = b // Now this b variable is assigned to _ which means that the program will run without the use of b

	// Coomparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	var matrix = [3][3]int = [3][3]int{}
}

func someFunction() (int, int) { // This function accepts nothing and returns 2 integers that's why there are two parenthesis
	return 1, 2

}
