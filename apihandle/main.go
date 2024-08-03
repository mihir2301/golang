package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"Author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

}

func serveHome(w http.ResponseWriter, r http.Request) {
	w.Write([]byte("<h1>Welcome</h1>"))
}

func GetallCourses(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("NO course found")
	return
}
func CreateOneCorse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data present")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses= append(courses,course)
	json.NewEncoder(w).Encode(course) 
	return
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","appliction/json")
	params:=mux.Vars(r)
	for index,course:=range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses,course)
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
}
