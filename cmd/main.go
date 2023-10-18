package main

import (
	"log"
	"net/http"

	"github.com/google-play-scraper-api/internal/middlewares"
	"github.com/google-play-scraper-api/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	// init app
	r := mux.NewRouter()

	// set content-type: application/json
	r.Use(middlewares.JsonContentTypeMiddleware)

	// add routers
	routes.AppRoutes(r)

	// start
	log.Fatal(http.ListenAndServe(":8000", r))
}
