package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type courseInfo struct {
	Title string `json:"Title"`
}

var courses map[string]courseInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!\n")
}

func allcourses(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all courses\n")
	kv := r.URL.Query()
	kv_string := fmt.Sprintf("Query kev/value: %v", kv)

	fmt.Println(kv_string)

	for k, v := range kv {
		kv_string2 := fmt.Sprintf("key: %v, val: %v", k, v)
		fmt.Println(kv_string2)

	}

	if val, ok := kv["country"]; ok {
		fmt.Println(val)
		fmt.Println(val[0])
		fmt.Println(ok)
	}

	//returns all the courses in JSON
	json.NewEncoder(w).Encode(courses)
}

func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Detail for course "+params["courseid"]+"\n")

	fmt.Fprintf(w, r.Method+"\n")

	if r.Method == "GET" {
		if _, ok := courses[params["courseid"]]; ok {
			json.NewEncoder(w).Encode(courses[params["courseid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found\n"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := courses[params["courseid"]]; ok {
			delete(courses, params["courseid"])
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found\n"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST - create new class
		if r.Method == "POST" {

			// read the string sent
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				//convert json to object
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course information in JSON format\n"))
					return
				}

				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"] + "\n"))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Duplicate course ID\n"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format\n"))
			}

		}

		// PUT - update or create class
		if r.Method == "PUT" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				//convert json to object
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course information in JSON format\n"))
					return
				}

				// Check if course exists, add only if course does not exist
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"] + "\n"))
				} else {
					// update course
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusNoContent)
					w.Write([]byte("201 - Course updated: " + params["courseid"] + "\n"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format\n"))
			}

		}
	}
}

func cond_test() {
	if h, k := two_test(); h && k {
		fmt.Println("here")
	}
}

func two_test() (bool, bool) {
	return true, true
}

func main() {
	//cond_test()

	courses = make(map[string]courseInfo)

	var newCourse courseInfo
	newCourse.Title = "Advanced Programming"
	courses["IOS301"] = newCourse
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)

	router.HandleFunc("/api/v1/courses", allcourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
