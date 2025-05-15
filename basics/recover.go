package basics

import "fmt"

func main() {

	// recover function is used to regain control of a panicking Goroutine.

	// Recover is a built in function that stops the propagation of a panic and returns the value passed to the panic call

	// The recover function is called inside the defer function so if the current goroutine is panicking, recover stops the panic and returns the value passed to the panic function it is returning.

	// Recover function is returning a value, but that value is only passed when there is a panic, so it's only passing the value passed to the panic function.

	process()
	fmt.Println("Returned from process")

}

// First we will make a function outside the main function and it will be the same process.

// Let's say we are running a process inside of our main thread, and this process has implemented a recover mechanism, using the defer and recover function

// rcover will return the error from the panicking goroutine
func process() {
	defer func() {
		if r := recover(); r!=nil {// way we implement a recovery is like this.
			fmt.Println("Recovered:",r)
		}
	}() // This defer function will return right before the process function(surrounding function is about to return)

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}

// Best Practices:-

// Always use recover with defer:- This ensures that recover is invoked in the same function where panic might ocur, allowing for controlled recovery.

// While recover can be used to catch panics, it is generally better to log or report the panic details rather than silently ignoring it, because recover is not just meant to recover your program from a panic, it is there to catch panics. It is there to report panics to you and then it's upto the developer to use that panic data to log and report the details of that panic and remedy the situation, whatever the problem is in our app, remedy that, dix that so there are no furthur panics for that reason 

// do not over use them it is not healthy for app