package basics 

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source:= rand.NewSource(time.Now().UnixNano()) // We are initializing a new souce with a known seed with a seed which is the current time.

	// And Now we are going to generate a random New Source

	random:=rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) +1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		// fmt.Scanln(guess) // A copy of guess is made and then that copy of guess is given to Scanlon and the orignal variable guess is not updated.

		fmt.Scanln(&guess) // As soon as we add this & sign that becomes the adddress of this variable.

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess<target {
			fmt.Println("too low! Try guessing a higher number ")
		} else {
			fmt.Println("Too high! Try Guessing a lower number.")
		}

		

	}

}