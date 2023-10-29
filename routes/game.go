package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/torbatti/nim-griz/models"
)

func GameR(w http.ResponseWriter, r *http.Request) {
	platform := chi.URLParam(r, "game")

	var model_games []models.Game
	App.Db.Distinct("name").Select("*").Where("name = ?", platform).First(&model_games)

	var games []Game
	for _, v := range model_games {
		game := Game{Year: v.Year, Name: v.Name, Platform: v.Platform, Developer: v.Developer, Publisher: v.Publisher}
		games = append(games, game)
	}
	// Data
	data := struct {
		Title string
		Games []Game
	}{
		Title: "Nim Griz",
		Games: games,
	}

	TemplateMaker(w, "views/pages/game.html", data)
}
