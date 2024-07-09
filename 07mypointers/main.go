package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int //Default pointer value
	fmt.Println("Default Value of pointer is ", ptr)

	myNumber := 23
	var ptr1 = &myNumber //Refrence to the address where pointer is pointing

	//this shows the address of the value where the value is stored
	fmt.Println("Value of Pointer is ", ptr1)

	//this shows the value of the address where the pointer is pointing
	fmt.Println("Value of pointer is ", *ptr1)

	*ptr1 = *ptr1 + 2 //this happens because the pointer is pointing to the actual value of number not any copy of it
	fmt.Println("The new Value is ", myNumber)
}
