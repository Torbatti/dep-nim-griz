package hx

import (
	"log"
	"net/http"

	"github.com/torbatti/nim-griz/models"
	"github.com/torbatti/nim-griz/routes"
)

type Form struct {
	Search string `form:"search"`
}

func Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form["search"][0])

	var model_games []models.Game
	if len(r.Form["search"][0]) > 3 {
		App.Db.Distinct("name").Select("name", "year", "platform").Where("name LIKE ?", "%"+r.Form["search"][0]+"%").Find(&model_games)
	} else {
		log.Println("INPUT LENGTH IS LESS THAN 4")
	}

	var games []routes.Game
	for _, v := range model_games {
		game := routes.Game{Year: v.Year, Name: v.Name, Platform: v.Platform}
		games = append(games, game)
	}

	// Data
	data := struct {
		Title string
		Games []routes.Game
	}{
		Title: "Nim Griz",
		Games: games,
	}

	routes.TemplateMaker(w, "views/partials/search.html", data)
}
