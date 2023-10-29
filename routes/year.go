package routes

import (
	"net/http"
	"strconv"

	"github.com/torbatti/nim-griz/models"
)

func Years(w http.ResponseWriter, r *http.Request) {

	// Years
	var game_ys []models.Game
	var years []string
	App.Db.Distinct("year").Select("year").Find(&game_ys)
	for _, v := range game_ys {
		if v.Year != 0 {
			years = append(years, strconv.Itoa(int(v.Year)))
		}
	}

	// Data
	data := struct {
		Title string
		Years []string
	}{
		Title: "Nim Griz",
		Years: years,
	}

	TemplateMaker(w, "views/pages/years.html", data)
}
