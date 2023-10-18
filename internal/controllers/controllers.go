package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"github.com/n0madic/google-play-scraper/pkg/app"
	"github.com/n0madic/google-play-scraper/pkg/developer"
	"github.com/n0madic/google-play-scraper/pkg/reviews"
	"github.com/n0madic/google-play-scraper/pkg/search"
	"github.com/n0madic/google-play-scraper/pkg/similar"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Error404(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&Response{
		Status:  404,
		Message: "Request not found!",
		Data:    nil,
	})
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&Response{
		Status:  200,
		Message: "App is running...",
		Data:    nil,
	})
}

func Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	if len(strings.Trim(q, " ")) > 0 {

		query := search.NewQuery(q, search.PriceAll, search.Options{
			Country:  "vi",
			Language: "vi",
			Number:   50,
		})

		err := query.Run()

		if err != nil {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}
		errors := query.LoadMoreDetails(20)
		if len(errors) > 0 {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}

		json.NewEncoder(w).Encode(&Response{
			Status:  200,
			Message: "Success",
			Data:    query.Results,
		})

	} else {
		json.NewEncoder(w).Encode(&Response{
			Status:  400,
			Message: "Keyword not found!",
			Data:    nil,
		})

	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	pkgName := path.Base(r.URL.Path)
	if len(strings.Trim(pkgName, " ")) > 0 {
		a := app.New(pkgName, app.Options{
			Country:  "vi",
			Language: "vi",
		})
		err := a.LoadDetails()
		if err != nil {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}
		err = a.LoadPermissions()
		if err != nil {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}
		json.NewEncoder(w).Encode(&Response{
			Status:  200,
			Message: "Success",
			Data:    a,
		})
	} else {
		json.NewEncoder(w).Encode(&Response{
			Status:  400,
			Message: "Package name not found",
			Data:    nil,
		})
	}

}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	pkgName := path.Base(r.URL.Path)
	if len(strings.Trim(pkgName, " ")) > 0 {
		r := reviews.New("com.activision.callofduty.shooter", reviews.Options{
			Number:   100,
			Language: "vi",
			Country:  "vi",
		})
		err := r.Run()
		if err != nil {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}

		json.NewEncoder(w).Encode(&Response{
			Status:  200,
			Message: "Success",
			Data:    r.Results,
		})
	} else {
		json.NewEncoder(w).Encode(&Response{
			Status:  400,
			Message: "Package name not found",
			Data:    nil,
		})
	}
}

func GetSimilars(w http.ResponseWriter, r *http.Request) {
	pkgName := path.Base(r.URL.Path)
	if len(strings.Trim(pkgName, " ")) > 0 {
		si := similar.New(pkgName, similar.Options{
			Number:   100,
			Language: "vi",
			Country:  "vi",
		})
		err := si.Run()
		if err != nil {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}
		json.NewEncoder(w).Encode(&Response{
			Status:  200,
			Message: "Success",
			Data:    si.Results,
		})

	} else {
		json.NewEncoder(w).Encode(&Response{
			Status:  400,
			Message: "Package name not found",
			Data:    nil,
		})
	}
}

// func GetCategory(w http.ResponseWriter, r *http.Request) {
// 	clusters, err := category.New(store.Game, store., category.Options{
// 		Country:  "vi",
// 		Language: "vi",
// 		Number:   100,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(&Response{
// 		Status:  200,
// 		Message: "Success",
// 		Data:    clusters,
// 	})
// }

func GetDeveloper(w http.ResponseWriter, r *http.Request) {
	devName := path.Base(r.URL.Path)

	if len(strings.Trim(devName, " ")) > 0 {
		dev := developer.New(devName, developer.Options{
			Number: 100,
		})
		err := dev.Run()
		if err != nil {
			json.NewEncoder(w).Encode(&Response{
				Status:  401,
				Message: "Error, something went wrong",
				Data:    nil,
			})
			return
		}
		json.NewEncoder(w).Encode(&Response{
			Status:  200,
			Message: "Success",
			Data:    dev.Results,
		})

	} else {
		json.NewEncoder(w).Encode(&Response{
			Status:  400,
			Message: "Package name not found",
			Data:    nil,
		})
	}
}
