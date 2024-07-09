package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitlist [4]string //Declaring a array
	fruitlist[0] = "Apple"  //Assigning the values to the array
	fruitlist[2] = "Bannana"
	fruitlist[3] = "Srawberry"

	fmt.Println("Fruitlist is :", fruitlist)                 //Printing the array
	fmt.Println("Lenght of fruit list is :", len(fruitlist)) //can find the length of an array

	/*as we can se that we just had given 3 inpputs in the array but it has a size of 4
	then the 2nd space where we didnt give any input can be seen to have a
	greater space between the 1st and the 3rd element of the array*/

	var veglist = []string{"Bhendi", "Tomato", "Potato"}
	fmt.Println("Vegetable list is : ", veglist)
	fmt.Println("Lenght of veg list is :", len(veglist))
}
