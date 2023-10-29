package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/torbatti/nim-griz/models"
	"github.com/torbatti/nim-griz/utils"
)

var App *utils.App

func TemplateMaker(w http.ResponseWriter, path string, data any) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("Template Parsing Error: %v", err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Template Execution Error: %v", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	// Years
	var game_ys []models.Game
	var years []string
	App.Db.Distinct("year").Select("year").Find(&game_ys)
	for _, v := range game_ys {
		if v.Year != 0 {
			years = append(years, strconv.Itoa(int(v.Year)))
		}
	}

	// PlatForms
	var game_ps []models.Game
	var platforms []string
	App.Db.Distinct("platform").Select("platform").Find(&game_ps)
	for _, v := range game_ps {
		if v.Platform != "" {
			platforms = append(platforms, v.Platform)
		}
	}

	// Data
	data := struct {
		Title     string
		Years     []string
		Platforms []string
	}{
		Title:     "Nim Griz",
		Years:     years,
		Platforms: platforms,
	}

	TemplateMaker(w, "views/pages/index.html", data)
}
