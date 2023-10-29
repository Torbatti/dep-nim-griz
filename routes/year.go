package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

type YGame struct {
	Name          string `json:"Game"`
	Year          uint   `json:"Year"`
	Publisher     string `json:"Publisher"`
	Developer     string `json:"Dev"`
	Platform      string `json:"Platform"`
	GameLink      string `json:"GameLink"`
	DevLink       string `json:"DevLink"`
	PublisherLink string `json:"PublisherLink"`
	PlatformLink  string `json:"PlatformLink"`
}

func Year(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")
	// Year
	var games []models.Game
	App.Db.Where("year = ?", year).Find(&games)

	// Data
	data := struct {
		Title string
		Year  string
		Games []models.Game
	}{
		Title: "Nim Griz",
		Year:  year,
		Games: games,
	}

	TemplateMaker(w, "views/pages/year.html", data)
}
