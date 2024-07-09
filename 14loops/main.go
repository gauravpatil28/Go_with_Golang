package main

import "fmt"

// There is no while loop in GOLANG

func main() {
	fmt.Println("Loops in GoLang")

	//declaring a slice
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// initialising a for loop

	// CASE 1
	// 	for d := 0; d < len(days); d++ {
	// 		fmt.Println(days[d])
	// 	}                                                   This is just commented it works fine

	// CASE 2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// CASE 3 (i)  (using comma ok syntax)

	// for index, day := range days {
	// 	fmt.Printf("Index is %v , and value is %v \n", index, day)
	// }

	// // CASE 3 (ii)  (using comma ok syntax)

	// for _, day := range days {
	// 	fmt.Printf("Value is %v \n ", day)
	// }

	value := 1

	for value < 10 {

		if value == 5 {
			goto jump
		}

		fmt.Println("Value is ", value)
		value++
	}

	// goto declaration

jump:

	fmt.Println("Welcome to the goto World 😁")
}
