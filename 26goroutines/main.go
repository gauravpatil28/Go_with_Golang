package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // Usually used as pointers

/* Jevha apan go vaparto ani techyane jo paryant to function respond nahi karat te responsibility waitgroup chi asti ki
   jo paryant te parat yet nahi program end nako hoyla techyat kahi functions astat:-
   1) Add(n):- hya add madhe apan jevdhe functions go func madhe takle aahet tyevdhe lihto ani mag jo paryant
   			   sagde respond nahi karat it waits.
   2) Wait():- jo paryant wait madhe sagde respond nahi karat it dosent allow the program to terminate.
   3) Done():- he execute zalya var program will terminate
*/

var mut sync.Mutex // Usually used as pointers

/* Jevha kuthi memory write hot aste tevha mutex chya help ne apan tya memory la lock kiva unlock karu shakto
   mhanje basically jecvha ek package write karel tevha kuthla dusra package interfere nahi karu shakat as the
   memory is locked
*/

func main() {

	// go greeter("Hello")
	// greeter("World")
	/* Jevha apan go vaparto kuthlya function chya samor mag apan tyala pathavto ani to techach call hoil
	   jevha apan tela call karu and for that only reason apan func madhe time use kela aahe ki to dar
	   3 sec nantar call hoil apan ajun padhatine tela bolavu shakto
	   pan kadhi apan time nahi dila tar fakt world print hoil karan apan tela bolvatach nahi aahot
	*/
	websitelist := []string{
		"https://go.dev",
		"https://lco.dev",
		"https://google.com",
		"https://fb.com",
		"https://instagram.com",
		"https://youtube.com",
	}

	for _, web := range websitelist {
		go getStatuscode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatuscode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("There is some error in your URL 🤪")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d is the status code for %s\n", res.StatusCode, endpoint)
	}
}
