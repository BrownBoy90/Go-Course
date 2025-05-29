package main

import "fmt"

func main() {
	var Q[4]map[string]string
	Q[0] = make(map[string] string)

	Q[0]["Q"] = "Which keyword is used to define a function in Go?"
	Q[0]["O1"] = "🔴def"
	Q[0]["O2"] = "🔴function"
	Q[0]["O3"] = "🔴func"
	Q[0]["O4"] = "🔴fun"

	Q[1] = make(map[string] string)

	Q[1]["Q"] = "Which of these companies created the Go programming language?"
	Q[1]["O1"] = "🔴Apple"
	Q[1]["O2"] = "🔴Google"
	Q[1]["O3"] = "🔴X"
	Q[1]["O4"] = "🔴Meta"

	Q[2] = make(map[string] string)

	Q[2]["Q"] = "Which language is known for its significant indentation?"
	Q[2]["O1"] = "🔴Python"
	Q[2]["O2"] = "🔴Java"
	Q[2]["O3"] = "🔴C++"
	Q[2]["O4"] = "🔴Swift"

	Q[3] = make(map[string] string)

	Q[3]["Q"] = "Which company developed the Android operating system originally?"
	Q[3]["O1"] = "🔴Google"
	Q[3]["O2"] = "🔴Android Inc."
	Q[3]["O3"] = "🔴Nokia"
	Q[3]["O4"] = "🔴Microsoft"
	fmt.Println("")
	fmt.Println("Hello, User!")
	fmt.Println("Welcome to Command line Quiz App!🥳🥳")
	fmt.Println("Answer the following Questions!")
	fmt.Println("You will be given a question along with the 4 options, you need to choose the most appropriate option:-")

	// I will make two variants one using slices and other using map

	// Map
	correct_options := []string{"O3", "O2", "O1", "O2"}
	var answers_opted[4] string
	for i,_ := range Q {
		fmt.Println(Q[i]["Q"])
		fmt.Println(Q[i]["O1"])
		fmt.Println(Q[i]["O2"])
		fmt.Println(Q[i]["O3"])
		fmt.Println(Q[i]["O4"])
		fmt.Printf("Enter the correct option number(write 'O1' for option 1): ")
		fmt.Scanln(&answers_opted[i])
		fmt.Println("")
	}
	// fmt.Println(answers_opted[0])
	// fmt.Println(answers_opted[1])
	// fmt.Println(answers_opted[2])
	// fmt.Println(answers_opted[3])

	fmt.Printf("You got %d out of 4", give_marks(correct_options, answers_opted))
	fmt.Println("")
}

func give_marks(correct_options []string, answers_opted [4]string) int {
	var marks int = 0
	for i:=0;i<4;i++{
		if correct_options[i] == answers_opted[i]{
			marks++
		}
	}

	return marks
}