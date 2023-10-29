package main

import (
	"log"
	"net/http"
	"os"

	// Insiders

	"github.com/torbatti/nim-griz/apis"
	"github.com/torbatti/nim-griz/middlewares"
	"github.com/torbatti/nim-griz/models"
	"github.com/torbatti/nim-griz/routes"
	"github.com/torbatti/nim-griz/utils"

	// chi related
	"github.com/go-chi/chi"

	// .env
	"github.com/joho/godotenv"

	// Gorm
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// https://stackoverflow.com/questions/35038864/how-to-access-global-variables
var app *utils.App

func makeApp() *utils.App {
	app = &utils.App{}
	routes.App = app

	// Database: Opening
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database .\n ", err.Error())
		os.Exit(2)
	} // db.Logger = logger.Default.LogMode(logger.Info)

	// Database: Migrations
	db.AutoMigrate(&models.Game{}, &models.List{}, &models.User{})

	// FIND IF DB IS EMPTY
	// if (DB_EMPTY){}
	// apis.Iterate("data")
	var game models.Game
	if result := db.Find(&game, "name = ?", "1000m Zombie Escape!"); result.Error != nil { //TODO: VALIDATE IF ERROR IS COLUMN NOT EXIST ERROR
		apis.Start(db, "datas")
	}

	// Connecting
	app.Router = chi.NewRouter()
	app.Db = db

	return app
}

func main() {
	app := makeApp()

	// Public Routes
	// Private ROutes

	// Setting Up .env
	godotenv.Load(".env") //os.Setenv(port, "8000")
	portString := os.Getenv("PORT")

	// MiddleWares // NOTE: all middlewares must be defined before routes on a mux(example: group routes)
	middlewares.Base(app.Router)

	// Root
	root := chi.NewRouter()
	app.Router.Mount("/", root)
	root.Get("/hi", utils.HelloWorld)

	// Public
	public := http.FileServer(http.Dir("./public"))
	root.Mount("/", public)

	// Apis
	api := chi.NewRouter() // NOTE: Use mux Then Mount
	root.Mount("/api", api)

	// View Routes
	root.Get("/", routes.Index)
	root.Get("/test", routes.Test)

	// Hx Routes
	hx := chi.NewRouter()
	root.Mount("/hx", hx)

	// Initial server
	server := &http.Server{
		Handler: app.Router,
		Addr:    ":" + portString,
	}

	// Listen And Serve
	log.Println("Listening On " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
