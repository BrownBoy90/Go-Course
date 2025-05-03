package main

import "fmt"

func main() {

    // for i:=1; i<=5; i++{
    //     fmt.Println(i) 
    // }

    // numbers := []int{1,2,3,4,5,6}
    // for index, value := range numbers{
    //     fmt.Printf("Index: %d, Value:%d\n", index, value)
    // }

    // %d-> Numbers, %v->value

    // for i:=1; i<=10; i++{
    //     if i%2==0 {
    //         continue
    //     }
    //     fmt.Println("Odd Number:",i)
    //     if i == 5 {
    //         break
    //     }
    // }

    // ASTERISK LAYOUT

    // rows:=5
    // // Outer Loop
    // for i:=1;i<=rows;i++ {
    //     // Inner loop for spaces before stars
    //     for j:=1;j<=rows-i;j++{
    //         fmt.Print(" ")
    //     }
    //     // inner loop for stars
    //     for k := 1; k<=2*i-1;k++ {
    //         fmt.Print("*")
    //     }
    //     fmt.Println() // Move to the next line
    // }

    // We can your for like this when we are ranging over

    // We can range over integers directly in for loop.

    for i:= range 10{
        fmt.Println(10-i)
    }


}
// Difference between Print and Println

// => fmt.Print adds a space if there are multiple arguments, but fmt.Println adds a new line after multiple statements and don't use Println for spaces and astericks
