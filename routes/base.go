package routes

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	TemplateMake(w, "views/pages/index.html", "")
}

func TemplateMake(w http.ResponseWriter, path string, data any) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("Template Parsing Error: %v", err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Template Execution Error: %v", err)
	}
}
