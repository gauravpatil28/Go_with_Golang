package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	Id   int
	Name string
	Data T
}

func Trial() {
	u := User[string]{
		Id:   4,
		Name: "Gaurav",
		Data: "String Data",
	}
	fmt.Println(u)
}

// func main() {
// 	u := User[int]{
// 		Id:   4,
// 		Name: "Gaurav",
// 		Data: 5,
// 	}
// 	fmt.Println(u)
// }
