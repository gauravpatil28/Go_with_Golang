package main //Building a API for a site which offers courses

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// -file means this thing goes into a seperate file

// Model For Course -file
type Course struct {
	CourseName  string  `json:"coursename"`
	CourseID    string  `json:"courseid"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
} //here the author is a type of author we just passed a pointer to the author

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// Fake DataBase
var Courses []Course

// Middleware ,helper  (These help you to perform some tasks i.e for eg, Encrypt the password , dont give access to DB and all)
func (c *Course) IsEmpty() bool {
	// return c.CourseName == "" && c.CourseID == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to Building API")

	// creating a router (who will listen to your requests)
	r := mux.NewRouter()

	// seeding (giving data to the database)

	Courses = append(Courses, Course{CourseID: "20", CourseName: "Go", CoursePrice: 200, Author: &Author{Name: "Gaurav", Website: "go.dev"}})
	Courses = append(Courses, Course{CourseID: "40", CourseName: "PY", CoursePrice: 600, Author: &Author{Name: "Gaurav", Website: "youtube"}})
	/* Apan varchya case madhe author cha reference pass karto karan apan tela ek pointer madhe ghetla aahe in the original struct */

	// Routing (IMP:-kuthlya route var kay hoil)
	r.HandleFunc("/", servehome).Methods("GET")            //hyacat localhost:6000/ var servehome func execute hoil
	r.HandleFunc("/courses", getallCourses).Methods("GET") //hyacat localhost:6000/courses var getAllcourses func execute hoil
	r.HandleFunc("/course/{id}", getoneCourse).Methods("GET")
	/*hyacat localhost:6000/course/{id} var getonecourse func execute hoil apna id hya karnane dila aahe karam
	apan input ghet aahot user kadun*/

	r.HandleFunc("/course", Createonecourse).Methods("POST")
	r.HandleFunc("/course/{id}", Updateonecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", Deleteonecourse).Methods("DELETE")

	// Listen to a port (where you local host will be executed)
	log.Fatal(http.ListenAndServe(":6000", r))
}

// Controllers -file

// serve home route

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to My API </h1>")) // you just send a slice of byte with the help of respnse writer
}

// We will be accepting all the values from our fake database
func getallCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the courses")
	w.Header().Set("Content-Type", "application/json")
	/* Set mhanje ki mi sangto aahe ki mi tula kuthlya type chya values det aahe hya case madhe
	mi tela content vaules det aahe ani jya json format madhe rahtil he basically sangaych asta
	ki kuthlya type chya values det aahes ani kuthlya type madhe rahtil tya values */

	json.NewEncoder(w).Encode(Courses)
	//hya method madhe mi json data encode kartoy
}

func getoneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course")
	w.Header().Set("Content-type", "application/json")

	// grabbing id from request

	params := mux.Vars(r)

	/* kahitari bug mule majha mux import nahi hot aahe ani tya karna mule he gadbad hot aahe nahitar
	varchya line madhe reader(r) read karel th ecourse id which the user is requesting ani ti store hoil
	params madhe */

	for _, course := range Courses {
		if course.CourseID == params["id"] { //We are checking for id hence we have to specify in params that it will be id
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//case where id is not found
	json.NewEncoder(w).Encode("Course id not found")
}

func Createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-type", "application/json")

	// what if the body is empty      (Mhanje user ne kahich nahi dila)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter some value")
	}

	// what if - {}		     (Mhanje user ne {} hya type madhe input dila)

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course) // Mala store nahi karaych mhanun mi underscore dila aahe

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Enter some valid json value")
		return
	}

	// generating a unique id ,string (cz we have given course id as a string)
	// append course into Courses  (adding this to our database)

	course.CourseID = strconv.Itoa(rand.Intn(100)) //Ikde course id create hoil

	Courses = append(Courses, course) // nava course aplya db madhe jail

	json.NewEncoder(w).Encode(course) // Ikde encode hoil to
}

func Updateonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-type", "application/json")

	// First Getting the ID from request
	params := mux.Vars(r) // getting the id

	// Process to follow
	// loop , id , remove , add with my ID
	for index, course := range Courses {
		if course.CourseID == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...) //slicing that id out of our slice(db)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) //Apan add karto aahe mhanun tya value la body mhadun decode kela
			course.CourseID = params["id"]
			Courses = append(Courses, course) //add that value to our slice(db)
			json.NewEncoder(w).Encode(course) // Converting to json
			return
		} else {
			json.NewEncoder(w).Encode("<h2>The given COurse id is invalid Pls check again</h2>")
		}
	}
}

func Deleteonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One course")
	w.Header().Set("Content-Type", "application/json")

	// Getting the course id which needs to be deleted

	params := mux.Vars(r)

	// loop , id , remove (course , course+1)

	for index, course := range Courses {
		if course.CourseID == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			break
		}
	}

}
