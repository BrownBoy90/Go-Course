// In go when we are declaring package the package should match the folder in which this package exist

// So the package should be according to the folder that the file is in except the main package.

// Any other package other than main should be inside the folder with the same name.

package basics

import "fmt"

var middleName = "Cane" // global scope is only limited to the package outside this package we cannot access it. 
// If we are declaring a package level variable(global) we cannot use gofer notation ':='

func main() {
	var age int

	// Type of the variable is optional if we are initializing the variable
	var name string = "John"
	var name1 = "Jane"
  
	// Go has provided a convinient way of initialising variables

	count := 10 // variable name := value which you want to initialize, this method of initialisation works only inside functions, you have to use htis notation locally

	lastname := "Smith" // herre var or string is not allowed

	// Default Values
	// Numeric Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions and structs: nil

	// variables in go have blocked scope just like C++


}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}

// here the compilor will stop when it has found a single mistake

// Variables live only in their scope, so their lifetime is within the scopes, when scope is created variables get initialised when the scope ends variable finishes

