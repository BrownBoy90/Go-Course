package basics

import "fmt"

func main() {
	// i:=1
	// for i<=5{
	// 	fmt.Println("Iteration:",i )
	// 	i++
	// }

	sum:=0
	for {
		sum+=10
		fmt.Println("Sum:", sum)
		if sum>=50 {
			break
		}
		// ++,-- works in go
	}
}