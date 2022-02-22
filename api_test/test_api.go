package test_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DrawflowDiagram struct {
	Name            string
	DrawflowContent string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var params DrawflowDiagram
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			fmt.Printf("Error on / post")
		}
		w.Write([]byte(params.Name))
		w.Write([]byte(params.DrawflowContent))
		fmt.Printf(params.Name)
	})

	http.ListenAndServe(":3333", r)
}
