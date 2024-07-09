package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")

	Hello()

	basicadd := adder(3, 8)

	/* We have to use comma ok operator bcz we have two type
	of values that we are returning we can also use another
	variable declared for that */

	// complexadd , message := badd(45,6,6,6)

	complexadd, _ := badd(4, 56, 46, 64, 545, 5)

	fmt.Println("Addition of 2 number is :", basicadd)

	fmt.Println("Addition of n numbers is :", complexadd)

}

func Hello() {
	fmt.Println("Basic declaration")
}

// when we give parameter to func and expext to return values we must also define the return type i.e int in this condition
func adder(val1 int, val2 int) int {
	return val1 + val2
}

/* For the foll func i have given the limit as infinite i can just add values
and it will keep giving me result for that we use the 3 dots and i am returning two data types a
total and a message thats why i have to give 2 return data types at the declaration */

func badd(val ...int) (int, string) {
	total := 0

	for _, sum := range val {
		total += sum
	}

	return total, "String return type"
}
