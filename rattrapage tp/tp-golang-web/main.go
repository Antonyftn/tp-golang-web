package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var viewCounter int
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type Student struct {
	Name  string
	Age   int
	Sex   string
	Image string
}

var students = []Student{
	{"Alice", 20, "Female", "/static/female.png"},
	{"Bob", 21, "Male", "/static/male.png"},
}

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/counter", counterHandler)
	http.HandleFunc("/form", formHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "students.html", students)
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	viewCounter++
	tmpl.ExecuteTemplate(w, "counter.html", viewCounter)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		sex := r.FormValue("sex")
		image := "/static/male.png"
		if sex == "Female" {
			image = "/static/female.png"
		}
		students = append(students, Student{name, age, sex, image})
	}
	tmpl.ExecuteTemplate(w, "form.html", nil)
}
