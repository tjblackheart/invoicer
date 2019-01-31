package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tjblackheart/invoicer/pkg/handlers"
)

// Header adds response headers
func Header(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

// Logger adds a logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsDev {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	})
}

// Auth checks for a valid JWT.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts, err := handlers.GetTokenAsString(r)
		if err != nil {
			json, _ := json.Marshal(err.Error)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(json)
			return
		}

		err = handlers.ValidateToken(ts)
		if err != nil {
			json, _ := json.Marshal(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(json)
			return
		}

		next.ServeHTTP(w, r)
	})
}
