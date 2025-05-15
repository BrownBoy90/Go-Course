// Range Keyword

package basics

import "fmt"

func main() {

	message := "Hello World"

	for i, v := range message {
		// fmt.Println(i,v) // v will print the unicode value
		fmt.Printf("Index: %d, Runes: %c\n", i,v) // alphabets are called Runes in   Go same as chars in C
	}

	// Range Keyword operates on a copy of the data structure it iterates over. Therefore, modifying index or value inside the loop does not affect the orignal data structure.

	//If we are iterating over a map or an array, or a slice, or even this string, when we modify any value inside the loop, the modification will not affect the orignal value.

// ----------RANGE BEHAVIOUR-----------
	// for arrays, slices and strings range iterates in order from first element to the last.

	// for maps, range iterates over key value pairs, but in an unspecified order, and for channels range

	// for chanels range iterates untill the chanel is closed.

	// If we do not need anything from range for eg one of two returns, practice using underscores. Using underscores prevents memory leaks by allowing go's garbage collectors to reclaim the memory.


}