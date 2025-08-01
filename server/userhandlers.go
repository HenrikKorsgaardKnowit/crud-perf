package server

import (
	"net/http"

	"github.com/henrikkorsgaard/crud-perf/db"
)

func userGetHandler(db *db.UserDatabase) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")

			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))

		},
	)
}

func userPostHandler(db *db.UserDatabase) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			id := r.PathValue("id")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))

		},
	)
}

func userHandler(db *db.UserDatabase) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))

		},
	)
}
