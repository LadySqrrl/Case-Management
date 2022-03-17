package main

import (
	"html/template"
	"net/http"

	"github.com/LadySqrrl/Case-Management/reports"
)

var testTemplate *template.Template

//ViewData shows widget data
type ViewData struct {
	Reevals []reports.Reeval
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("go-html/index.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := reports.ReevalsByMonth()
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
