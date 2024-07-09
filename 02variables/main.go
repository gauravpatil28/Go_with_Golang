package main

import "fmt"

const Login = "Access" /*we have given L of login as capital letter so it is a public
variable so any one can access that*/

func main() {
	var name string = "Gaurav"
	fmt.Println(name)
	fmt.Printf("The username is of the type: %T \n", name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The username is of the type: %T \n", isLoggedIn)

	var value int = 748
	fmt.Println(value)
	fmt.Printf("The username is of the type: %T \n", value)

	var valfloat float32 = 748.72537426
	fmt.Println(valfloat) //float 32 is used to display only 4 digits after decimal
	fmt.Printf("The username is of the type: %T \n", valfloat)

	var valfloatb float64 = 748.7253742684656
	fmt.Println(valfloatb) //float 64 is used to display only 8 digits after decimal
	fmt.Printf("The username is of the type: %T \n", valfloatb)

	//default values and aliases

	var anothervar int
	fmt.Println(anothervar) // When you don't initialise the values of int it will take default value as 0
	fmt.Printf("The username is of the type: %T \n", anothervar)

	var anotherval string
	fmt.Println(anotherval) // When you don't initialise the values of string it will take default string as blank
	fmt.Printf("The username is of the type: %T \n", anotherval)

	//implict type
	// here you have not declared the variable type but the lexer determines and hence it does not throw error
	var website = "youtube.com"
	fmt.Println(website)

	//no var  style

	/*we do not use here the code word "var" then also it runs because we have to use the Walurus operator(:=)
	and it will run without any error, this walurus operator can only be used inside any function means it
	cannot be declared globally*/

	noofuser := 4000
	fmt.Println(noofuser)

	//accesing const

	fmt.Println(Login)
	fmt.Printf("The username is of the type: %T \n", Login)

	str := "IIII"
	leng := len(str)
	for i := 0; i < leng; i++ {
		fmt.Println(string(str[i]))
	}
}
