package routes

import (
	"github.com/google-play-scraper-api/internal/controllers"
	"github.com/gorilla/mux"
)

func AppRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/api/search", controllers.Search).Methods("GET")
	r.HandleFunc("/api/detail/{package_name}", controllers.GetDetail).Methods("GET")
	r.HandleFunc("/api/reviews/{package_name}", controllers.GetReviews).Methods("GET")
	r.HandleFunc("/api/similar/{package_name}", controllers.GetSimilars).Methods("GET")
	r.HandleFunc("/api/developer/{dev_name}", controllers.GetDeveloper).Methods("GET")
	// r.HandleFunc("/api/category/{category_name}", controllers.GetCategory).Methods("GET")

	// 404
	r.HandleFunc("*", controllers.Error404).Methods("GET")
	r.HandleFunc("*", controllers.Error404).Methods("POST")
	r.HandleFunc("*", controllers.Error404).Methods("PUST")
	r.HandleFunc("*", controllers.Error404).Methods("PATCH")
	r.HandleFunc("*", controllers.Error404).Methods("DELETE")
}
