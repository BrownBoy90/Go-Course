package basics

import "fmt"

func main() {
	// In go defer is a mechanism that allows you to postpone the execution of a function untill the surrounding function returns, either because the surrounding function executed a return statement reached the end of its function body or because the corresponding go routine is panicking.

	// Go routines are functions which run in the background, which are running concurrently in the background, and they are not part of the main thread.

	// any thing which is a go routine is thrown to the back so that it finshes off its work, not in the main thread, not blocking the main thread but in the background, and then comes back and joins the main thread once it's finished.  

	// defer statement is followed by a function or method call.
	// The function call is evaluated immediately, but the execution is deferred untill the surrounding function returns.

	// we can make a defer function at the beginning, but it will only execute once the surrounding function returns.

	// defer function is always part of the another function.

	// surrounding function means the function that encloses the defer function.

//-------Let's See that in action-------



	// So anything with defer, any statement or any function which has defer keyword as its prefix will be deferred till the end of that function.

	// we can also have multiple deferred statements in a function, and they will be executed in a last in first out order when the function returns.

	process(10)

	// When we are using defer arguments to deferred functions are evaluated immediately when the defer statement is encountered.

	// So just because defer dunction is getting executed at the end, doesn't mean that it is getting evaluated at the end

	// defer function will be evaluated immediately as sooon as it is encountered.

	// This is IMPORTANT to note, especially when the values of arguments might change by the time the deferred function executes.

//-------PRACTICAL USECASES OF DEFER----
	// defer is comonly used to ensure that resources like files or database connections are closed after they are opened, and when using Mutexes, which we will get to later on in our course

	// When using Mutexes to synchronize Goroutines, defer can be used to ensure that a mutex is unlocked even if a function panics.

	// Defer functions are also useful for logging and tracing entry and exit points of functions.

	// The best practices of using defer is to keep deferred actions short. Deferred should be short and simple to avoid unexpected behaviour and to keep the functions's logic clear. Using defer in LOOPS or NESTED FUNCTIONS can lead to subtle bugs if not handled carefully. So be very very careful.

	// It is like try catch finally. and it is like finally. Defer and finally both do cleanup actions. SO there are certain clean up activities that we delegate to defer.

	// defer ensures that the critical actions are performed in a predictable manner regardless of how functions exit which is especially important when maintaining code reliability and readability.

	// However, it should be used judiciously and with an understanding of its implications on program execution flow.
}

func _() {
	// Instead of just making a function we can defer anything.
	defer fmt.Println("Deferred Statement executed") // This gets executed after the normal print statement
	fmt.Println("Normal Execution statement")
}

func process(i int) {
	// Instead of just making a function we can defer anything.

	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("First deferred Statement executed") // This gets executed after the normal print statement
	defer fmt.Println("Second deferred Statement executed")
	defer fmt.Println("Third deferred Statement executed")
	i++
	fmt.Println("Normal Execution statement")
	fmt.Println("Value if i:",i)
}	// Here the third deferred statement was executed first and then the second and the first. So the one that was deferred earliest was executed at the last

