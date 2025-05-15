// In go functions can declare multiple return values by listing them in parenthesis after the parameter list

// Each return must have a distinct type, so here's how we do that.

package basics

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2,...){
	// code Block
	// return returnValue1, returnValue2, ...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	// Benefits of multiple return values:-
	// => Error Handling

	result, err := compare(3,2)
	if err!= nil{
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	
}

func divide(a, b int) (quotient int, remainder int) { // named re  turn
	quotient = a / b
	remainder = a % b
	return /*quotient, remainder*/
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil // If we do not have an error to send we will send nil
	} else if b>a {
		return "b is greater than a", nil // similarly if we have an error, but we do not have the correct return value that fuction was supposed to pass on then we will use nil in place of string and an error in place of the error
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
