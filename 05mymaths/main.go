package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	// var num1 int = 4
	// var num2 float32 = 6.3

	// fmt.Println("Sum is ",num1 + num2)

	// Random Number from math/rand

	//rand.Seed(time.Now().UnixNano())   This is a code

	/* We use the above command becasue the time always changes t=hence we are more likely to get
	any random number the next time we run the code without this the system will always give you
	a certain specific value*/

	//fmt.Println("Random number you want is : ", rand.Intn(10))  this is a code

	/*rand intn prints random number below or equal to the number you have specified in the brakets
	you will only recieve 1 if you dont give the command prior to that which will make the numbers
	random*/

	// Random number from crypto

	newrandomnum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("Random number", newrandomnum)
}
