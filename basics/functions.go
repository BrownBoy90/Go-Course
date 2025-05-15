// functions are fundamental building b locks in go, encapsulating reusable code blocks that can be invoked multiple times with diferent inputs.

package basics

import "fmt"

func main() {

	// func <name>(parameters list) returnType {
		// code block
		// return value
	// }

	// A function name should be a valid identifier and should follow ghost naming conventions.

	// If it is a public function it should start with an uppercase letter, and if it's a private function then it needs to start with a lower case. 

	// fmt.Println() // We can see that this is a public function that's why it starts with Capital P. If there was some function that was confined to use within the fmt package it should have been started with small letter.

	// If return statement is omitted the function return default zero values for their types.

	// sum := add(1,2)
	// fmt.Println(sum)
	// We can call the function in Println as well

	// Another type of functions are anonymous functions we can call them closures or literals

	// greet := func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }

	// greet()

	// operation := add

	// result := operation(3,5)

	// fmt.Println(result)

	// This demonstrates that we can use function as types and functions in go can be assigned to variables, passed as arrguments to other functipns and returned from functions, making them a first class citizen or a first class object

//---------FIRST CLASS CITIZENS--------
	// It refers  to entities that have no restrictions on their use throughout the language and can be treated uniformly throughout the language.When an entity is a first class citizen it means that you can perform a wide range of operations on it, just as you would with basic data types like integers or strings. And these operations would typically include passing as arguments returning from functions, assigning to variables ot storing in data structures

	result := applyOperation(5, 3, add)
	fmt.Println("5+3 =",result)

	// returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6*2 =", multiplyBy2(6))
}

func add(a, b int) int  { // If there are multiple return types you have to include them in parenthesis separated by commas
	return a + b
}

// Function thaat takes a function as an argument
func applyOperation(x int, y int, operation func(int, int)int) int {
	return operation(x,y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func (x int) int {
		return x * factor
	}
}