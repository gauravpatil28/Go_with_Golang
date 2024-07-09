package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	// var fruitlist = []string{}    default syntax for slices
	// slices means a array with n number of values

	var fruitlist = []string{"Apple", "Kiwi", "Mango"}
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "Strawberry", "Banana", "Grapes") // Elements can be added into slices like this using append
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:4])

	/* Basically [1;4] means that just display from the 1st element till the 3rd element
	because the last limit is non-inclusive we can also just give one value i.e [1:] or [:4] in this manner */

	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	/* this is another way of declaring a slice we just declare the type of slice and
	how many elemnets it would be handling */

	highScores[0] = 631
	highScores[1] = 126
	highScores[2] = 238
	highScores[3] = 467
	//highScores[4] = 379

	//values assinged equal to the size given , it will show error if we print the last commented line

	highScores = append(highScores, 343, 756, 543)

	//but by using the append comand we can give n numbers and it will allocate the memory accordingly

	fmt.Println(highScores)

	// the following command will sort all the int data types into ascending order

	sort.Ints(highScores)

	fmt.Println(highScores)

	// the following command will give boolean ans if it is sorted or not

	fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a index value (any value) from a given slice

	var courses = []string{"Python", "Golang", "C", "C++", "Java", "SQL", "HTML"}

	fmt.Println(courses)

	var index = 5

	courses = append(courses[:index], courses[index+1:]...)

	/*what we have done in the above command is that we first told to print values till the index
	(where index is non-inclusive) then we added a second argument which told to print from index+1
	so that the  value we need to be deleted will be removed*/
	/* AND we used the 3 dots(...) at the end because append is designed to take 1 argument
	when we used the 3 dots it allows and runs both of our arguments*/

	fmt.Println(courses)
}
