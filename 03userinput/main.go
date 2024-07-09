package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Rating")

	// comma ok || err err

	/*so in go it does not show where the error has occured for catching that we have the
	comma ok ,err thing in which we can take input or catch the error where it occurs
	in this we can also do it for only one thing at a time if we want input we just use
	input and  give _(underscore) ont he place of error adn vice versa */

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating : ", input)
	fmt.Printf("Type of Rating : %T ", input)

}
