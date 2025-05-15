package main

import "fmt"

func main() {

	// CLosures are a powerful concept that allows functions to capture and manipulate variables that are defined outside of their body

	// Definition:- A closure is a function value that references variables from outside its body. The function may access and assign to the captured variables, and these variables persist as long as the closure itself is referenced.

	// Closures work with lexical scoping, meaning they capture variables from their surrounding context where they are defined. This allows closures to access variables even after the outer function has finished execution. Closures leverage this by allowinng functions to capture and manipulate their surrounding state.

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
