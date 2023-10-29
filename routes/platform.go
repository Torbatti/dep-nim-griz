package routes

import (
	"net/http"

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
