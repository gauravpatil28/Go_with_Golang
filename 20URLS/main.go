package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentID=6385fskjdf"

/*basically a URL consist of few parts which you need to remeber:-1) Scheme (https)
															 	 2) Host (lco.dev)
															 	 ३) Port Number (in this case 3000)
															    ४) Path (/learn)
 After Path all that is there that is called as the Query	  	५) Rawquery (coursename)
 & is used as a comma in the url  */

func main() {
	fmt.Println("URL's in GOlang")

	fmt.Println(myurl)

	//Parsing a URL means traversing a URL accessing elemnts or different parts of URL

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryparameter := result.Query() //Here we are making the queries seperated OR can access them one by one
	fmt.Printf("The type of Query are : %T \n", queryparameter)

	fmt.Println(queryparameter["coursename"])

	for _, val := range queryparameter {
		fmt.Println("The Query Param are :", val)
	}

	/*Bascially kadhi tula url banvaych aahe tar asa banavshil Ani ha hyachat "&" hi sign khup easy
	aahe miss karayla tar te lakshat thevayla lagel ani mag systematically tula ek ek karun goshti deyla lagel
	AS GIVEN ABOVE PARTS OF URL
	ani ha te ek goshta lakshat thev parts baddal path,query ani baki sarkha */

	anotherurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Gaurav",
	}

	/*Need to convert the url into string to get it working and thats why we need to convert it to a string*/
	newurl := anotherurl.String()

	fmt.Println(newurl)

}
