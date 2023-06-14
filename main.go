package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template // variable type pointer/template from package Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml")) // all file in template folder
}

func main() {
	http.HandleFunc("/", index) // anything that coms in to main rout will going to run function index
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil) // that is going to ListenAndServe on port 8080
}

func index(w http.ResponseWriter, r *http.Request) { // it needs a ResponseWriter and it needs a pointer to Request
	// io.WriteString(w, "hello Aram")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
