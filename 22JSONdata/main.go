package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to JSON data")

	// EncodeJSON()
	DecodeJSONfromweb()
}

type Courses struct {
	Name     string `json:"Coursename"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"Tags,omitempty"`
}

//Tags chya pudhe je aahe mhanje slice declare karat aahes tu which will contain the data in the format as string
/*je mi pudhe lihla aahe `` hya madhe te tya name ne print hoil in the json file ani omitempty mhanje jikde tags nastil
kiva nil lihla asel tar tikde kahich print nahi honar ani Password pan nahi print honar karan tyachya pudhe - dila aahe mhanje
Dont print it */
func EncodeJSON() {

	//creating obejct of the struct

	lcoCourses := []Courses{
		{"Python", 299, "Freecourses.in", "Easypass123", []string{"language", "py"}},
		{"Java", 699, "Freecourses.in", "gdfsu39", []string{"app-dev", "jv"}},
		{"Reactjs", 199, "Freecourses.in", "asihdgf", nil}, //tu Nil deu shkato if any place you dont feel to give input
	}

	//Package this data as JSON

	// JSONdata , err := json.Marshal(lcoCourses)

	// if err != nil{
	// panic(err)
	// }

	/*ikde don type chya goshti astat ek asta normal marshal sarkha jasa var decalare kela aahe ani ek asta ki tu
	Marshal madhe ajun ek type ne declare karu shakto ani that will help you to show more neatly presented output.*/

	newJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	/*hya method madhe kay hota ki pahile tar tu to var deto which you want to give to convert to json
	ani mag ek prefix value deychi aste ji to dar \t  nantar show karel ji ki hya case madhe kahich nahi aahe
	ani mag last la mi \t dila aahe jyachyane ki values madhe spacing yeil*/

	if err != nil {
		panic(err)
	}

	fmt.Println(string(newJson))

}

func DecodeJSONfromweb() {

	/*Tar kahli tu jo data recieve kartos to eka byte format madhe asto so tu ek slice banavto
	  jyachat tu to data recieve karshil ani to json data detos*/

	recieveddata := []byte(`
	{
		"Coursename": "Java",
		"Price": 699,
		"Website": "Freecourses.in",
		"Tags": ["app-dev","jv"]		
	}		
	`) //Leaving comma at the end of a given data makes it invalid so be sure that there is no comma after the end of last line.
	//In this case Tags is the last line.

	var lcocourse Courses

	/*var je apan var decalre kela aahe to aahe aplya struct cha type cha var aahe ani apan jo data recieve kela aahe
	aplyala to tya struct madhe takaycha aahe so for that usecase we use the given method of declaring a var*/

	checkifvalid := json.Valid(recieveddata)

	//Just checking if the data recieved from the web is of correct type or not (it will return the value in BOOL)

	if checkifvalid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(recieveddata, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON NOT VALID")
	}
	/*var apan techa refrence pass kela aahe karan aplyala original madhe change karaych aahe aplyala techi copy pnako aahe*/

	// Some cases you just need to add data to the key value pair

	var myOnlinedata map[string]interface{}
	json.Unmarshal(recieveddata, &myOnlinedata)
	fmt.Printf("%#v\n", myOnlinedata)

	/*Apna hya dusrya case madhe fakt key values var focus karto ani apan tya karna sathi apan ek map banavto which is type of
	  Strnig ani mag tyachat interface pass karto mahanje apan jo data recieve karu tyachat array pan rahu shakto
	  mag to ignore nahi honar ani nat recieve hotil.*/

	for k, v := range myOnlinedata {
		fmt.Printf("Key is %v and the data is %v and Type is %v \n", k, v, v)
	}

}
