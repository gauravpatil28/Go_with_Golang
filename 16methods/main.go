package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")

	/*Basically methods classes chya functions sarkhe work kartat. tula kadhi struct baghitla tar sagdi info bhetat hoti
	Pan kadhi methods vaparle tar je hava tech bhetta */

	gaurav := User{"Gaurav", 20, "gaurav@go.dev", true}

	fmt.Println(gaurav)

	/*Hyala print karayche pan diff types aahet tu Printf Vapru shakto for required details, ani kadhi proper mapping sobat hava asel
	tar to %+v takun kadhu shakto*/

	//  fmt.Printf("The Info of Gaurav Is %+v :", gaurav)

	fmt.Println("Age of Gaurav is = ", gaurav.Age)

	gaurav.GetStatus()
	gaurav.FAge()
}

/*He khali apla normal struct define kela aahe*/
type User struct {
	Name           string
	Age            int
	Email          string
	ActivityStatus bool
}

// Methods

/* Ha format aahe hyacha func mag bracket madhe struct ani ek object(je ki class che objects represent kartat) mag naav ani empty () */

func (u User) GetStatus() {
	fmt.Println("The Status of User is :", u.ActivityStatus)
}

/* hyachat tu values la change karun present karu shakto jya original data madhe change nahi hot te temprorily
present hotat change houn*/

func (u User) FAge() {
	u.Age = 22
	fmt.Println("The age is :", u.Age)
}
