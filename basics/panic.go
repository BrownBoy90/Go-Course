package basics

import "fmt"

func main() {
	// When a function enncounters a panic.
	// It stops executing it's current activities, unwinds the stack, and then executes any deferred functions.
	// This function continues up the stack untill all functions have returned, at which point the program terminates.

	// A panic is typically used to signal an unexected error condition where the program cannot safely, used to signal an unexpected error condition where the program cannot proceed safely.

	// The syntax of  panic function is called with an optional argument of any type, which represents the  value associated with the panic.

	// SYNTAX: panic(interface{})
	// Example of a valid input
	process(-69)
	// defer will execute when the function returns a value, but will also execute even the function is panicking

	// Now, however, misuse of panic can lead to unpredictable behavior, and it should be avoided in scenarios where regular error handling suffices

}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("Input must be a non-negative number.")
		// fmt.Println("After Panic")

		// defer fmt.Println("Deferred 3")

	}
	fmt.Println("Processing Input:", input)
}
