package main

import (
	"log"
	"net/http"
	"os"

	// Insiders
	"github.com/torbatti/nim-griz/middlewares"
	"github.com/torbatti/nim-griz/routes"

	// chi related
	"github.com/go-chi/chi"

	// .env
	"github.com/joho/godotenv"

	// Gorm

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	Db     *gorm.DB
}

func makeApp() *App {
	app := &App{}

	// Database: Opening
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database .\n ", err.Error())
		os.Exit(2)
	} // db.Logger = logger.Default.LogMode(logger.Info)

	// Database: Migrations
	db.AutoMigrate()

	// Connecting
	app.Router = chi.NewRouter()
	app.Db = db

	return app
}

func main() {
	app := makeApp()

	// Setting Up .env
	godotenv.Load(".env") //os.Setenv(port, "8000")
	portString := os.Getenv("PORT")

	// MiddleWares // NOTE: all middlewares must be defined before routes on a mux
	middlewares.Base(app.Router)

	// Root
	root := chi.NewRouter()
	app.Router.Mount("/", root)
	root.Get("/hi", routes.HelloWorld)

	api := chi.NewRouter() // NOTE: Use mux Then Mount
	root.Mount("/api", api)

	// Static Files : Fonts/JsLibs/Imgs/etc.

	// View Routes
	// http.Handle("/foo", routes.Base(http.ResponseWriter, *http.Request))
	// BackEnd Routes
	// Apis Routes
	// Hx Routes
	// Utils

	// Initial server
	server := &http.Server{
		Handler: app.Router,
		Addr:    ":" + portString,
	}

	// Listen And Serve
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
