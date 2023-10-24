package main

import (
	"log"
	"net/http"
	"os"

	// chi related
	"github.com/go-chi/chi"

	// .env
	"github.com/joho/godotenv"
)

func main() {
	// Setting Up .env
	godotenv.Load(".env") //os.Setenv(port, "8000")
	portString := os.Getenv("PORT")

	// Setting Chi Router
	router := chi.NewRouter()

	// MiddleWares
	// NOTE: all middlewares must be defined before routes on a mux

	// V1 Router
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	// Setting Up app server
	app := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	// Static Files : Fonts/JsLibs/Imgs/etc.
	// View Routes
	// http.Handle("/foo", routes.Base(http.ResponseWriter, *http.Request))
	// BackEnd Routes
	// Apis Routes
	// Hx Routes
	// Utils

	// Listen And Serve
	err := app.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
