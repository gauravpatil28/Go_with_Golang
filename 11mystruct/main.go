package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	// these does not have any kind of inheritance in golang

	gaurav := User{"Gaurav", "gaurav@go.dev", true, 19}

	fmt.Println("Details of Gaurav ", gaurav)

	//if we want everything printed with parameters we use +v

	fmt.Printf("Detalis are %+v \n", gaurav)
	fmt.Printf("Name is %v and his age is %v \n ", gaurav.Name, gaurav.Age)
}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}
