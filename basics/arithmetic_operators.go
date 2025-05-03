package basics

import (
	"fmt"
	"math"
)

func main() {
	// Variable Declaration
	var a, b int = 10, 3

	result := a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Substraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22.0 / 7
	fmt.Println(p)

	// Overflow with signed integer
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)
	
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64

	fmt.Println(smallFloat)
}

// Be mindful of potentially overflow and underflow issues, especially with large numbers

// Overflow:- It ocurs when the result of a computation exceeds the maximum value that can be stored in a given datatype (after overflow unsigned goes to 0 and signed int goes to negative minimum)

// In float point numbers underflow can lead to loss of precision or significant digits in calculation involving very small values.

// operator precidence:-
// Parenthesis
// Multiplication*, Division/, remainder%
// Addition+, Substraction-

// In Go, when you divide two integers the result is  truncated integer

// to make the result a float when dividing atleast one of a, b should be a float
