package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://learncodeonline.in"

/*apan constant madhe ani globallyich declare karto url karan te accesebale asta ani ek error hoyche chances kami astat*/

func main() {
	fmt.Println("Welcome to Handling the webrequests")

	//getting the webreq
	response, err := http.Get(url)

	/*apan ikde comma err vaparto karan handling web request can cause many errors and if any occurs we can easily
	understand or catch or rectify it hence it is advisable to check for errors all the time to get correct output*/

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //THIS IS YOUR RESPONSIBILITY TO CLOSE THE CONNECTION(RESPONSE).
	//ani defer bcz it should be closed at the end.

	//Checking the type of response recieved

	fmt.Printf("The response recieved is of the type : %T", response)

	showdata, err := io.ReadAll(response.Body)

	/*IO will help us with majority of reading things in this and we want to read the body only hence BODY you
	can also go for HEADER STATUS and more*/

	if err != nil {
		panic(err)
	}

	//  fmt.Println(showdata)    Just uncomment and check what happens if not made string

	content := string(showdata)

	/*we need to convert to string else we will see the data in numbers*/

	fmt.Println(content)
}
