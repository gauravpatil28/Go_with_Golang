package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	mydefer()
}

// world , One ,Two  					main madhla printing sequence
// 0,1,2,3,4						    for loop chi printing sequence
//Hello 43210 Two One world				final printing sequence

func mydefer() {
	for i := 0; i < 5; i++ {
		rdefe print(i)
	}
}

/*So tar "DEFER" kay karta eka statement kiva line la code chya last madhe execute karto
AT THE EXTREME LAST JUST BEFORE LAST CURLY BRACKET
ani te eka stack sarkha kaam karta FILO(first in last out) kiva LIFO
hyacha neat working dakhavla aahe var tya 3 statements ne
mhanje basically to pahile main madhe HELLO execute karel ani mag func karan tya madhe defer nahi aahe
Ani function statement chya defer adhi execute hotil karan ti statement execution adhi hoil*/
