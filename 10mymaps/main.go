package main

import "fmt"

func main() {
	fmt.Println("Maps In Golang")

	//Maps declaration

	languages := make(map[string]string)

	//giving values to map

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages are : ", languages)
	fmt.Println("JS is short form for : ", languages["JS"])

	//deleting a value from map

	delete(languages, "RB")

	fmt.Println(languages)

	// Loops in Golang

	for key, value := range languages {
		fmt.Printf("For key %v , The value is %v \n", key, value)
	}
	// %v is used for giving out the value
	// Range takes values from the languages or just defines where the values are meant to be coming from
}
