package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/process", processor)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		// if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Hello Full Cycle</h1>")
	d := struct {
		Title string
	}{
		Title: "Processamento",
	}

	tpl.ExecuteTemplate(w, "index.gohtml", d)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		Title string
		First string
		Last  string
	}{
		Title: "Processamento",
		First: fname,
		Last:  lname,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
