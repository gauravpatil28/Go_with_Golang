package main

import (
	"fmt"
	"log"
	"net/http"

	"githib.com/gauravpatil28/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()

	fmt.Println("Server is getting Started ...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at the port 4000..")
}
