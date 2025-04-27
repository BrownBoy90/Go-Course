package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase (used for naming types in go)
	// Case Eg. CalculateArea, UserInfo, NewHTTPRequest
	// for eg. Structs, interfaces, enums

	// snake_case
	// Eg. useer_id, first_name, http_request
	// used for variable, constants and file names

	// UPPERCASE
	// use case is constants

	// mixedCase
	// Eg. javaScript, htmlDocument
	// can be used to name variables or certain identifiers, especially when dealing with external libraries

	// Naming practices

	// maintain consistncy if you are using mixedCase for variables don't use snake_case for that
	// minimize the use of abbreviation unless they are widely understood

	// Make package name short concise and lowercased, without underscore or mixedCase

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}
