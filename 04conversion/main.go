package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza app")
	fmt.Println("Pls rate the app")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating, ", input)

	/* We wanted to see of the conversion would show error or not so this time
	we have declared the error variable with the input variable*/

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	/*strconv is the func which we use to convert the given string in the data type you want
	and string trimspace is use because a string has some values ahead of it or trailing it
	and we dont want those values hence the trimspace func is used     IMP:- Strings is a very useful func in GOLANG
	and we cant add 1 directly to the input because it is having its type as string and a string and a int
	cannot be directly added to one another hence we need to convert it before we add 1 to it */

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your Rating is, ", newRating+1)
	}
}
