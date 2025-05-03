package basics

import "fmt"

func main() {
// The switch statements provides a concise way to evaluate multiple possible conditions against a single expression

// Used in cases where you had 25 to 30 conditions
	// switch statement in go is (switch case default) (fallthrough) if you want to check all the options in the switch its a rarecase so use fallthrough in that case. Only the case right after fall through will be executed 
	// switch expression {
	// case value1:
		// code
		// fallthrough
	// case value2:
		// code
	// case value3:
		// code
	// case value4:
		// code
	// default:
		// Code to be executed if expression does not match any value
	// }

	// switch statement in go is (switch case break default)
	// switch expression {
	// case value1:
		// code
		// break;
	// case value2:
		// code
	// case value3:
		// code
	// case value4:
		// code
	// default:
		// Code to be executed if expression does not match any value
	// }

	// number:=15
	// switch{
	// case number<10:
	// 	fmt.Println("Number is less than 15")
	// } // case can have expressions also, cases separated by commas are or sense.


	// Switch Case for Type Assertions 
	checkType(10)
	checkType(3.4)
	checkType("Hello")
	checkType('c')
	checkType(true) 

}

func checkType(x interface{}) /* This interface() means that x can take any data Type*/{
	switch x.(type) { // We cannot use fallthrough in type switch
	case int:
		fmt.Println("It's and integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("It's Unknown Type")
	}

}