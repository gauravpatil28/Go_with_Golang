package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

// Type 1: Where you can define type and pass it to contain any type of value you want to assign
// type Num interface {
// 	int | int16 | float64 | float32
// }

// func Add[T Num](a T, b T) T {
// 	return a + b
// }

// Type 2: Here we have this package constraints Where the ordered function has all of the Datatypes defined
// func Add[T constraints.Ordered](a T, b T) T {
// 	return a + b
// }

// Type 3: We use a custom type and to use that we need to use tilda(~) to make it work
// type UserId int
// func Add[T ~int | float64](a T, b T) T {
// 	return a + b
// }
// func main() {
// 	a := UserId(3)
// 	b := UserId(4)
// 	result := Add(a, b)
// 	fmt.Println(result)
// }

func Mapvalues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newVaulues []T
	for _, v := range values {
		newVaulue := mapFunc(v)
		newVaulues = append(newVaulues, newVaulue)
	}
	return newVaulues
}

type CustomMap[T comparable, V constraints.Ordered] map[T]V // comparable means anything which can be compared

func main() {
	result := Mapvalues([]int{1, 2, 3}, func(i int) int {
		return i * 3
	})
	log.Println(result)

	res := User[string]{
		Id:   4,
		Name: "Gaurav",
		Data: "String Data",
	}
	fmt.Println(res)
}
