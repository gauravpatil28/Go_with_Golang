package router

import (
	"githib.com/gauravpatil28/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetmyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateaMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkasWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteaMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllMovie).Methods("DELETE")

	return router
}

/*
One very interesting thing about routes is that if you give a route like  /home  so it will go check the home route and perform the function
but it you give it as /home/  then the last forward slash can be very helpful because of that slash if you give anything ahead of the
route eg. /home/huaieghfiaj  then it will just consider the home and perform it. will ignore the values if not matching and if in the
first home route if we only gave   /home/  then it would have thrown a error. and if there is a route   /home/go    and in the url
u give  /home/go  then it will access that route only   but if you give   /home/skjdbvskh   then it will only access the home route.

*/
