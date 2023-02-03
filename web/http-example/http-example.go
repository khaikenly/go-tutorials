package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Struct fields must start with upper case letter (exported) for the JSON package to see their value.
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func (student *Student) NewStudent(id int, name string, age int, sex string) Student {
	return Student{
		Id:   id,
		Name: name,
		Age:  age,
		Sex:  sex,
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page!")
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to contact page!")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to about page!")
}

func Students(w http.ResponseWriter, r *http.Request) {
	students := make([]Student, 0)

	for i := 0; i < 5; i++ {
		s := Student{}
		students = append(students, s.NewStudent(i, "khainp", 23+i, "male"))
	}

	json.NewEncoder(w).Encode(students)
}

func main() {
	fmt.Println("Http Example")

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contact", ContactPage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/api/students", Students)

	log.Fatal(http.ListenAndServe(":80", nil))
}
