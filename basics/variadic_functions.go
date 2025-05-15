package basics

import "fmt"

func main() {

// Variadic functions are functions which can accept variable amount of parameters

// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ...type3 ) returnType{
	//function body
	// }

	// fmt.Println("Summ of 1, 2, 3:",sum(1,2,3))
	// statement, ```(statement, total)

	// Variadic parameters must be the last parameter in the function signature.

	// If you have a slice of values that you want to pass to a variadic function, you can do so by appending ellipsis the three dots to the slice when calling the function, and this syntax unpacks the slice into individual arguments.

	// So to unpack the slice we use ellipsis which destructures structures the slice into individual elements into individual values.

//--------Let's See that in Action------
	numbers:= []int{1,2,3,4,5,9}

	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence:",sequence3,"Total:", total3)

}
  
func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for  _, v:= range nums {
		total += v
	}
	return sequence, total
}