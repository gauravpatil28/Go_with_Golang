package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("How to make GET req")

	//Getwebreq()
	//PostJSONwebreq()
	PostFormwebreq()
}

func Checkerr(err error) {
	panic(err)
}

func Getwebreq() {
	const myurl = "http.//localhost:8000/get"

	response, err := http.Get(myurl)

	// The above command gets the url and stores it in the response url

	Checkerr(err)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Status code is: ", response.StatusCode)
	fmt.Println("Length code is: ", response.ContentLength)

	/*Ata ikde don Methods yetat:-

	I :- ki normally content la string madhe wrap karun print karun dya mhanje req msg bhetun jail.*/

	// Asa

	fmt.Println("Message recieved is :", string(content))

	/*II:- ani dusra mhanje ki apan strings package import karu shakto ani tya package madhe apan
	  ek var banavto ani strings package chya use karun ek builder banavto ani tya variable madhe sagda
	  raw data as usual store rahto mag te aplya var asta ki tya data la kuthlya form madhe vapraych
	  we can normally use it to print the string format or make it print the bytes and many more.*/

	var responsestring strings.Builder

	bytecount, _ := responsestring.Write(content)

	fmt.Println(bytecount)

	fmt.Println(responsestring.String())

}

func PostJSONwebreq() {

	const muyurl = "http://localhost:8000/post"

	// fake json payload(data)

	requestdata := strings.NewReader(`
		{
			"Company" : "Developing Welfare",
			"Age" : 20,
			"Founder" : "Gaurav Suhas Patil"
		}
	`)

	/*
		This is how you can post the data you want to post to a web site by using the imp  ``  you can do
			all of this in the json format all the text that is wrapped in the given `` and written in json
			can be posted to a site or url you want.
			json format:
			It should be written in the following format
			{
				"":""
			}
	*/

	response, err := http.Post(muyurl, "application/json", requestdata)

	/*
		Post madhe 3 goshti astat:
		 1) url link nahi tar url declare kelela
		 2) Content type kay aahe ta hya case madhe mi server madhe it is application/json mhanun te
		 3) ani the thing you want to upload/POST
	*/

	Checkerr(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PostFormwebreq() {
	const myurl = "http://localhost:8000/postform"

	//Form data

	data := url.Values{}

	data.Add("Firstname", "Gaurav")
	data.Add("Lastname", "Patil")
	data.Add("Email", "gaurav@go.dev")

	response, err := http.PostForm(myurl, data)

	Checkerr(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	Checkerr(err)

	fmt.Println(string(content))

}
