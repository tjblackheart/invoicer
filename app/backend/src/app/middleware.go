package app

import (
	"encoding/json"
	"net/http"
)

func (app Application) jsonMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func (app Application) authMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts, err := app.getTokenAsString(r)
		if err != nil {
			json, _ := json.Marshal(err.Error)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(json)

			return
		}

		err = app.validateToken(ts)
		if err != nil {
			json, _ := json.Marshal(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(json)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app Application) corsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Add("Vary", "Access-Control-Request-Headers")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, Authorization")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		next.ServeHTTP(w, r)
	})
}
