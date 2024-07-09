package main

import "fmt"

func main() {
	fmt.Println("Welcome to If else in Go")

	number := 20
	var output string

	if number > 20 {
		fmt.Println("Greater then 20")
	} else if number < 20 {
		fmt.Println("Less than 20")
	} else {
		fmt.Println("Number is exactly 20")
	}

	if 8%2 == 0 {
		output = "It is an even number"
	} else {
		output = "it is odd number"
	}

	fmt.Println(output)

	// we can perform one more type of if else that is initialising and assigning value and then checking

	if num := 5; num < 10 {
		fmt.Println("Less than 10")
	} else if num > 10 {
		fmt.Println("More than 10")
	} else {
		fmt.Println("Exactly 10")
	}
}

