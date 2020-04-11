package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tjblackheart/Invoicer/backend/pkg/models"
	"github.com/tjblackheart/Invoicer/backend/pkg/pdf"

	log "github.com/sirupsen/logrus"
)

// Map simplifies a json response
type Map map[string]interface{}

// CORS preflight OPTIONS handler
func (app Application) cors(w http.ResponseWriter, r *http.Request) {
	return
}

// REST
func (app Application) login(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	c := &models.Credentials{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	u, err := models.Authenticate(c)
	if err != nil {
		switch err.(type) {
		case models.ErrInvalidCredentials:
			app.error(w, err, http.StatusForbidden)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	ts, err := app.generateToken(u)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, Map{"token": ts, "user": u})
}

func (app Application) createUser(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	if err := u.Create(); err != nil {
		switch err.(type) {
		case models.ErrUnique, models.ErrValidation:
			app.error(w, err, http.StatusBadRequest)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	app.json(w, u)
}

func (app Application) getUser(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	vars := mux.Vars(r)

	u, err := models.FindUser(vars["uuid"])
	if err != nil {
		switch err.(type) {
		case models.ErrNotFound:
			app.error(w, err, http.StatusNotFound)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, u)
}

func (app Application) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u, err := models.FindUser(vars["uuid"])
	if err != nil {
		switch err.(type) {
		case models.ErrNotFound:
			app.error(w, err, http.StatusNotFound)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	patch := &models.User{}

	if err = json.NewDecoder(r.Body).Decode(&patch); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	if u, err = u.Update(patch); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, u)
}

func (app Application) updatePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u, err := models.FindUser(vars["uuid"])
	if err != nil {
		switch err.(type) {
		case models.ErrNotFound:
			app.error(w, err, http.StatusNotFound)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	var pwReq models.PasswordChangeRequest
	if err = json.NewDecoder(r.Body).Decode(&pwReq); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	if err := u.UpdatePassword(&pwReq); err != nil {
		switch err.(type) {
		case models.ErrValidation, models.ErrInvalidCredentials:
			app.error(w, err, http.StatusBadRequest)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, Map{})
}

//

func (app Application) getInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	invoice, err := models.FindInvoice(uuid, vars["id"])
	if err != nil {
		app.error(w, err, http.StatusNotFound)
		return
	}

	app.json(w, invoice)
}

func (app Application) getInvoices(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	invoices, err := models.InvoiceGetAll(uuid)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, invoices)
}

func (app Application) createInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	invoice := &models.Invoice{}
	err = json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	err = invoice.Create(uuid)
	if err != nil {
		switch err.(type) {
		case models.ErrValidation:
			app.error(w, err, http.StatusBadRequest)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app Application) setPaid(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	var payload models.PaidPayload
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	if err = models.InvoiceSetPaid(uuid, &payload); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	invoices, err := models.InvoiceGetAll(uuid)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, invoices)
}

func (app Application) cancelInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	invoice, err := models.FindInvoice(uuid, vars["id"])
	if err != nil {
		app.error(w, err, http.StatusNotFound)
		return
	}

	invoice.IsCancelled = true
	app.db.Save(&invoice)

	w.WriteHeader(http.StatusNoContent)
}

//

func (app Application) getCustomer(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	vars := mux.Vars(r)

	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	customer, err := models.CustomerGet(uuid, vars["id"])
	if err != nil {
		app.error(w, err, http.StatusNotFound)
		return
	}

	app.json(w, customer)
}

func (app Application) getCustomers(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	customers, err := models.CustomerGetAll(uuid)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, customers)
}

func (app Application) createCustomer(w http.ResponseWriter, r *http.Request) {
	customer := &models.Customer{}
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	customer, err = models.CustomerCreate(uuid, customer)
	if err != nil {
		switch err.(type) {
		case models.ErrUnique, models.ErrValidation:
			app.error(w, err, http.StatusBadRequest)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	app.json(w, customer)
}

func (app Application) updateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid, err := app.getUUID(r)

	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	customer, err := models.CustomerGet(uuid, id)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	var patch models.Customer

	if err = json.NewDecoder(r.Body).Decode(&patch); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	if err = customer.Update(&patch); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, customer)
}

func (app Application) removeCustomer(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.CustomerDelete(uuid, id); err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}
}

// PDF

func (app Application) printInvoice(w http.ResponseWriter, r *http.Request) {
	uuid, err := app.getUUID(r)

	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	i, err := models.FindInvoice(uuid, vars["id"])

	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	u, err := models.FindUser(uuid)
	if err != nil {
		switch err.(type) {
		case models.ErrNotFound:
			app.error(w, err, http.StatusNotFound)
			return
		}

		app.error(w, err, http.StatusInternalServerError)
		return
	}

	g := pdf.Generator{Invoice: i, User: u}
	filename, err := g.Generate()

	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	b64, err := g.Base64(filename)
	if err != nil {
		app.error(w, err, http.StatusInternalServerError)
		return
	}

	app.json(w, Map{"filename": filename, "content": b64})
}

//

func (app Application) json(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
}

func (app Application) error(w http.ResponseWriter, err error, status int) {
	if status >= http.StatusInternalServerError {
		log.Error(err)
	}

	http.Error(w, err.Error(), status)
}
