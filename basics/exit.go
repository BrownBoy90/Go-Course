package basics

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred statement ")
	// os.exit is a function that terminates the program immidiately with the given status code.

	// It is useful for the situations where you need to halt the execution of the program completely, without defering any functions or performing any cleanup operations, that means that the exit will be done in a hastily fashion, without doing any cleanup, or without. When OS.exit is called, it stops the program execution immidiately.
	// the function takes an integer argument which represents the status code returned to the operating system, conventionally a status code of 0 indicates successful completeion, while any non-zero status code

	fmt.Println("Starting the main function.")

	// Exit with a status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of the main function")

}

// It is used immediately to stop the execution completely without furthur processing.

// We can use os.exit when a specific condition is met, that requires the program to stop the execution completely without furthur processing.

// We also use os.exit to differentiate between different exit statuses to indicate success, failure, or specific error conditions to the calling environment or system.

// Best Practices:-
// Avoid deferred action, also ensure that all necessary cleanup operations are performed explicitly before calling os.Exit()

// Be careful about status codes, 0 is for success and non-zero is for error conditions, is for specific error conditions to communicate the output of the program execution to the calling environment. 

// avoid abusive use
// It should be used sparingly and only when truly necessary, such as in unrecoverable error scenarios or when the program must stop immidiately.