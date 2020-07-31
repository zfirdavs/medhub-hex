package rest

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// RESTful provides interface for handlers
type RESTful interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func PrintAllRoutes(r chi.Router) {
	walkFunc := func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Fatalf("Logging err: %s\n", err.Error())
	}
}
