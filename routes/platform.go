package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/torbatti/nim-griz/models"
)

func Platforms(w http.ResponseWriter, r *http.Request) {
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
		Platforms []string
	}{
		Title:     "Nim Griz",
		Platforms: platforms,
	}

	TemplateMaker(w, "views/pages/platforms.html", data)
}
func Platform(w http.ResponseWriter, r *http.Request) {
	platform := chi.URLParam(r, "platform")
	log.Println(platform)

	var games []models.Game
	App.Db.Where("platform = ?", platform).Find(&games)

	// Data
	data := struct {
		Title    string
		Platform string
		Games    []models.Game
	}{
		Title:    "Nim Griz",
		Platform: platform,
		Games:    games,
	}

	TemplateMaker(w, "views/pages/platform.html", data)
}
