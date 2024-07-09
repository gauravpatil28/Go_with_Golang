package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study")
	presentime := time.Now()
	//fmt.Println(presentime)

	fmt.Println(presentime.Format("02-01-2006  15:04:03  Monday"))
	/* we can format using .Format and very imp that the given DATE:-02-01-2006 Is the only date you can use
	to define the format because it is what it is and the time should be given using the same numbers
	and if u want to see the day you should give monday*/

	/*If you want to declare the date in the program*/

	// 	newdate := time.Date(2003, time.March, 28, 07, 30, 23, 0, time.Local)
	// 	fmt.Println(newdate.Format("02-01-2006 15:04:03 Monday"))
}
