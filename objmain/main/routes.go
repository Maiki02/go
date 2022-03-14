package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tomiok/course-phones-review/gadgets/smartphones/web"
	reviews "github.com/pruebaapi/general/reviews/web"
	"net/http"
)

func Routes(
	sph *web.CreateSmartphoneHandler,
	reviewHandler *reviews.ReviewHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Get("/hello", helloHandler)
	mux.Post("/reviews", alumnoHandler.AddAlumnoHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "miqueas")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}

