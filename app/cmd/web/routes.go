package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(app.headersMW, app.loggingMW)

	r.HandleFunc("/", app.index).Methods(http.MethodGet)

	auth := r.PathPrefix("/auth/").Subrouter()
	auth.Use(app.jsonMW)
	auth.HandleFunc("/login", app.login).Methods(http.MethodPost)
	auth.HandleFunc("/register", app.createUser).Methods(http.MethodPost)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(app.jsonMW, app.authMW)
	api.HandleFunc("/invoice/{id:[0-9]+}", app.getInvoice).Methods(http.MethodGet)
	api.HandleFunc("/invoice", app.getInvoices).Methods(http.MethodGet)
	api.HandleFunc("/invoice", app.createInvoice).Methods(http.MethodPost)
	api.HandleFunc("/invoice/payment", app.setPaid).Methods(http.MethodPost)
	api.HandleFunc("/invoice/pdf/{id:[0-9]+}", app.printInvoice).Methods(http.MethodGet)

	api.HandleFunc("/customer/{id:[0-9]+}", app.getCustomer).Methods(http.MethodGet)
	api.HandleFunc("/customer", app.getCustomers).Methods(http.MethodGet)
	api.HandleFunc("/customer", app.createCustomer).Methods(http.MethodPost)
	api.HandleFunc("/customer/{id:[0-9]+}", app.updateCustomer).Methods(http.MethodPut)
	api.HandleFunc("/customer/{id:[0-9]+}", app.removeCustomer).Methods(http.MethodDelete)

	api.HandleFunc("/user/{uuid:[a-z0-9-]{36}}", app.getUser).Methods(http.MethodGet)
	api.HandleFunc("/user/{uuid:[a-z0-9-]{36}}", app.updateUser).Methods(http.MethodPut)

	if app.config.production == false {
		r.Methods(http.MethodOptions).HandlerFunc(app.cors)
		r.Use(app.corsMW)
	}

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./ui/dist/assets/"))))

	return r
}
