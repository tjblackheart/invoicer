package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tjblackheart/invoicer/pkg/models"
	"github.com/tjblackheart/invoicer/pkg/pdf"
)

// HTML

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/dist/index.html")
}

// REST

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	c := &models.Credentials{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := models.Authenticate(c)
	if err != nil {
		if err == models.ErrInvalidCredentials {
			http.Error(w, "Invalid credentials!", http.StatusForbidden)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ts, err := app.generateToken(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"token": ts, "user": u})
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := u.Create()

	if err != nil {
		if err == models.ErrUnique {
			http.Error(w, "This email is already in use.", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// remove from response
	u.Password = ""
	u.RepeatPassword = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	vars := mux.Vars(r)

	u, err := models.FindUser(vars["uuid"])
	if err != nil {
		if err == models.ErrUserNotFound {
			http.Error(w, "No such user.", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Password = ""
	json.NewEncoder(w).Encode(u)
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u, err := models.FindUser(vars["uuid"])
	if err != nil {
		if err == models.ErrUserNotFound {
			http.Error(w, "No such user.", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	patch := &models.User{}

	if err = json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u, err = u.Update(patch); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Password = ""

	json.NewEncoder(w).Encode(u)
}

//

func (app *application) getInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	invoice, err := models.FindInvoice(uuid, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(invoice)
}

func (app *application) getInvoices(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	invoices, err := models.InvoiceGetAll(uuid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(invoices)
}

func (app *application) createInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	invoice := &models.Invoice{}
	err = json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = invoice.Create(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) setPaid(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := models.PaidPayload{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = models.InvoiceSetPaid(uuid, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	invoices, err := models.InvoiceGetAll(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(invoices)
}

//

func (app *application) getCustomer(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	vars := mux.Vars(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer, err := models.CustomerGet(uuid, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func (app *application) getCustomers(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customers, err := models.CustomerGetAll(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customers)
}

func (app *application) createCustomer(w http.ResponseWriter, r *http.Request) {
	customer := &models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer, err = models.CustomerCreate(uuid, customer)
	if err != nil {
		if err == models.ErrUnique {
			http.Error(w, "The customer already exists.", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func (app *application) updateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid, err := app.getUUID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := models.CustomerGet(uuid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var patch models.Customer
	if err = json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = customer.Update(&patch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func (app *application) removeCustomer(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.CustomerDelete(uuid, id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// PDF

func (app *application) printInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	invoice, err := models.FindInvoice(uuid, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := models.FindUser(uuid)
	if err != nil {
		if err == models.ErrUserNotFound {
			http.Error(w, "No such user.", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO
	file, err := pdf.Generate(invoice, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"file": file})
}
