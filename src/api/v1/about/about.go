package about

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

/*
FIXME
*/
type About struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

/*
FIXME
*/
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAbout)
	return router
}

/*
FIXME
*/
func GetAbout(w http.ResponseWriter, r *http.Request) {

	aboutBody := About{
		Name:    "StatsAPI",
		Version: "Server version 0.1, API version v1",
	}
	render.JSON(w, r, aboutBody)
}
