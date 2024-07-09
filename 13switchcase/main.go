package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Switch case in GOlang")

	//We gonna check this out with a game of dice
	//for that we need to generate a random number

	rand.Seed(time.Now().UnixNano())

	dicevalue := rand.Intn(6) + 1

	// // we added 1 because the end limit is non inclusive

	fmt.Println(dicevalue)

	switch dicevalue {
	case 1:

		fmt.Println("Dice value is 1 you can open")

	case 2:

		fmt.Println("DIce value is 2 you can move 2 spots")

	case 3:

		fmt.Println("DIce value is 3 you can move 3 spots")

		fallthrough // this is used to print the current case and the case following it

	case 4:

		fmt.Println("DIce value is 4 you can move 4 spots")

	case 5:

		fmt.Println("DIce value is 5 you can move 5 spots")

	case 6:

		fmt.Println("DIce value is 6 you can move 6 spots and Roll again")

	default:
		fmt.Println("How can dice exceed 6 ???")
	}

	// There is one more advantage of switch case in Go that we can add multiple cases in one case no need to make differnet cases for it

	/*
		switch days{
		case "Saturday" , "Sunday":
			fmt.Println("Weekend hai Enjoy Karo")
		}
	*/

	// You can also give conditions inside the switch also

	/*
		switch{
		case dice value == 1:
			fmt.Println("You can move one line ahead")
		}
	*/

}
