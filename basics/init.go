package basics

import "fmt"

func init() {
	fmt.Println("Initializing package1 ...")
}
func init() {
	fmt.Println("Initializing package2 ...")
}
func init() {
	fmt.Println("Initializing package3 ...")
}
func init() {
	fmt.Println("Initializing package4 ...")
}

func main() {

	// The init function is a special functiont hat can be declared in any package. It is used to perform initialization tasks for the package before it is used.

	// It has no parameters and no return values. It's declared like any other function, but with the name init.

	// Go executes init functions automatically when the package is initialized. This happens before the main function is executed. 

	// The init function always gets executed before the main function, and it occurs exactly once per package, even if the package is imported multiple files, it will happen only once.

	// Within a single package, go executes init functions in the order in which they are declared.

	// If there are multiple init functions, they execute sequentially following their textual order in the package file.

	// Usage of init functions:-
	// It is used for tasks such as initializing variables, performing setup operations, registering components or configurations, and initializing state required for the package to funciton correctly.

	fmt.Println("Inside the main function")


}

// Best Practices:-
	// Avoid Side effects that might affect the behaviour of other packages our cause unexpected behaviour during package initialization, also ensure that the order of init function execution is predictable and does not lead to dependencies, that could cause initialization failures.

	// Also document the purpose and side effects of init functions, especially in larger packages or libraries to aid understanding and maintainance.