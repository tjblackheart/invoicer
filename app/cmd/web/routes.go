package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(app.headersMW, app.loggingMW)
	r.HandleFunc("/", app.index).Methods("GET")

	auth := r.PathPrefix("/auth/").Subrouter()
	auth.Use(app.jsonMW)
	auth.HandleFunc("/login", app.login).Methods("POST")
	auth.HandleFunc("/register", app.createUser).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(app.jsonMW, app.authMW)
	api.HandleFunc("/invoice/{id:[0-9]+}", app.getInvoice).Methods("GET")
	api.HandleFunc("/invoice", app.getInvoices).Methods("GET")
	api.HandleFunc("/invoice", app.createInvoice).Methods("POST")
	api.HandleFunc("/invoice/payment", app.setPaid).Methods("POST")
	api.HandleFunc("/invoice/pdf/{id:[0-9]+}", app.printInvoice).Methods("GET")

	api.HandleFunc("/customer/{id:[0-9]+}", app.getCustomer).Methods("GET")
	api.HandleFunc("/customer", app.getCustomers).Methods("GET")
	api.HandleFunc("/customer", app.createCustomer).Methods("POST")
	api.HandleFunc("/customer/{id:[0-9]+}", app.updateCustomer).Methods("PUT")
	api.HandleFunc("/customer/{id:[0-9]+}", app.removeCustomer).Methods("DELETE")

	api.HandleFunc("/user/{uuid:[a-z0-9-]{36}}", app.getUser).Methods("GET")
	api.HandleFunc("/user/{uuid:[a-z0-9-]{36}}", app.updateUser).Methods("PUT")

	if app.config.production == false {
		api.Use(app.corsMW)
	}

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./ui/dist/assets/"))))

	return r
}
