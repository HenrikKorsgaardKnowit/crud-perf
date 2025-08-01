package server

import (
	"net/http"

	"github.com/henrikkorsgaard/crud-perf/database"
)

func New(db *database.UserDatabase) http.Handler {

	var handler http.Handler = addRoutes(db)

	return handler
}

// refactored into independent route function to aid testing
func addRoutes(db *database.UserDatabase) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/healthy", healthy())
	//Returns JSON
	mux.Handle("GET /users/{email}", userGetHandler(db))
	mux.Handle("POST /users", userPostHandler(db))
	mux.Handle("GET /users", userHandler(db))
	return mux
}

func healthy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("I'm alive"))
	})
}
