package main

/*
	HA BASICALLY AAHE KI BACKEND KASA BANTA
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Modules")

	Greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("GET") //WE just need the GET method in this case

	log.Fatal(http.ListenAndServe(":3000", r))

	/* Ikde Apan ek server host kela aahe je ki aplyala servehome navacha func cha web page deil
	   ani apan hya varchya command la respect deun sangitla aahe ki mala 3000 port number var host karaych aahe \
	   ani majha router r aahe, He aplyala err pan deu shakto ani for that apan comma ok syntax vapru shakto
	   pan hya case madhe apan log navach package vaparla aahe which will log the value if something goes wrong
	*/

}

func Greeter() {
	fmt.Println("Hello there mod users")
}

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME TO THE REAL SIDE OF CODING ✌️</h1>"))

}

/*
  Hyachat r mhanje reader which reads the value anyone is sending and w is the writer by the help of which
  you send the response to the query which you have with the help of the reader(r)
*/

/*
   kahi goshti astat termianl badal kadhi apn "go build ." he use kela tar hya case madhe aplya code ne given
   goshti build hotat ani mag apan "go list all" vapru shakto je ki aplyala dakhavtil all the packages used
   ani more classified hava asel tar apna "go list -m all" vapru shakto je ki aplyala dakhavtil tech packages je are used
   we have few more "go list -m -versions github.com/gorilla/mux" which will lis down all the versions of gorilla mux
   pubilicaly available. we also have "go mod why ...."in the dots you can ask why we are dependent on any of the package,
   "go mod graph" will pull out all the graphs that what is dependent on what
   We have couple more:-
   "go mod edit -go 1.19.4" :- This edits the version of go in the mod file
   "go mod edit -module .." :- This edits the module name of the the mod file
   "go mod vendor"			:- This brings everything from the local system not from the web
	but if we still run the command "go run main.go" it will bring everything from web
	we need to run "go run -mod=vendor main.go"  it will search the vendor first
*/
