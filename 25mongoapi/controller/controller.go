package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"githib.com/gauravpatil28/mongoapi/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// the following will be used to connect us to the database we just created
const connectionstring = "mongodb+srv://Trial:28032003@cluster0.zlcysqi.mongodb.net/?retryWrites=true&w=majority"

// Giving our database a NAme
const dbName = "Netflix"

// Giving a name to a collection
const colName = "Watchlist"

// MOST IMPORTANT (here you take a reference of the mongodb collection)
var collection *mongo.Collection

// Connecting with mongodb
func init() {
	// init mhanje initialization that will only be executed only one when the program is compiled the first time

	// Client option
	clientoption := options.Client().ApplyURI(connectionstring)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientoption) //Bascially context helps you to connect to things which are not installed in your pc

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo Connection Secured Successfully")

	collection = client.Database(dbName).Collection(colName) //Once connection is secured naming the DB and naming the collection

	// Collection refernce
	fmt.Println("Collection instance/reference is ready")

	// fmt.Println(collection)

	/* the client is the thing which will help us connect to the DB or the Collection or whatever we need and that will be called
	whenever a connection is required */
}

/* Apan jo action plan vaparnar aahe to asa aahe ki apan ek different helpers banavnar je ki aplyala vaues gheun yetil check kartil
   ani baki sagda sort karun aplyala values detil ani ek dusre helperes rahtil MONGOdb helpers je ki tya values tyenchya kadun
   gheun DB var update kartil */

// MONGODB helpers -file (hyana eka seperate file madhe taku shakto)

// Serve Home
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kay chalu aahe yaarr "))
}

// Insert 1 record

func insertoneMovie(movie models.Netflix) {
	// hya varchya line madhe apan aplya models madhun Netflix struct la vaparla

	inserted, err := collection.InsertOne(context.Background(), movie)

	//Jevha kadhi apan DB sobat interact karto aplyala context vaprava lagto

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in DB with id : ", inserted.InsertedID) //inserted id is created with the help of insertOne func
}

// Update 1 record

func updateoneMovie(movieID string) {

	id, err := primitive.ObjectIDFromHex(movieID)
	//the line of code will convert the given movie id in the form which can be read by mongodb

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	// mongo db takes id as _id and this will find it that id

	update := bson.M{"$set": bson.M{"watched": true}}
	// $ stands for flag in mongo and we are filtering the value and marking it as watched in it

	result, err := collection.UpdateOne(context.Background(), filter, update)
	//collection makes our job easier by its func first it will filter and then update and store in result

	fmt.Println("Modified count : ", result.ModifiedCount)
}

// Delete 1 record

func deleteoneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deletecount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with the delete count : ", deletecount)
}

// Delete all Records from Mongo DB

func deleteAllrecords() int64 {
	// 2 padhtine lihu shaktos ek tar filter sobat nahitar filter chya values saral filter chya jagevar pass karun

	// filter := bson.D{{}}       `` Ani ha bson madhe D ani M don goshti use hotat D mhanje ordered ani M mhanje unordered
	deletecount, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of Movies Deleted : ", deletecount.DeletedCount)
	return deletecount.DeletedCount //hya dogahi padhatine tumhi count return kiva show karu shaktat
}

// Get all Movies

func getallMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M // declaring this to store the movies in it

	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie) //decode
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
		//if decoded add to movies
	}

	defer cursor.Close(context.Background())

	return movies
}

// ACTUAL Controllers

// Get all movies

func GetmyMovies(w http.ResponseWriter, r *http.Request) {
	// we have given this func first alphabet as capital bcz we want to export it
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "GET")

	allmovies := getallMovies()

	json.NewEncoder(w).Encode(allmovies)
}

// Create a movie

func CreateaMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") // Only post method allowed for this func

	var movie models.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertoneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// Mark movie as watched

func MarkasWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateoneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete a Movie

func DeleteaMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteoneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all Movies

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteall := deleteAllrecords()

	json.NewEncoder(w).Encode(deleteall)
}
