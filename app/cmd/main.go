package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tjblackheart/invoicer/pkg/db"
	"github.com/tjblackheart/invoicer/pkg/handlers"
)

var (
	Build                = "unknown"
	Version              = "unknown"
	IsDev                = false
	host, secret, dbFile string
)

func main() {
	setupEnv()

	db.InitDB(dbFile, IsDev)
	defer db.Conn.Close()

	handlers.JWTSecret = secret
	migrateTables()

	r := mux.NewRouter()
	if IsDev {
		r.Methods("OPTIONS").HandlerFunc(handlers.CORS)
	}

	handleStatus(r)
	handleAuth(r)
	handleAPI(r)
	handleStatic(r)

	fmt.Printf("Server running at http://%s ... CTRL+C to quit\n", host)
	log.Fatal(http.ListenAndServe(host, r))
}

func setupEnv() {
	if os.Getenv("APP_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("Error reading environment variables.")
		}

		IsDev = true
	}

	host = os.Getenv("APP_HOSTNAME")
	if host == "" {
		log.Fatalln("Invalid hostname")
	}

	secret = os.Getenv("APP_SECRET")
	if secret == "" {
		log.Fatalln("Invalid secret")
	}

	dbFile = os.Getenv("APP_DB_FILE")
	if dbFile == "" {
		log.Fatalln("Invalid database file")
	}
}

func migrateTables() {
	handlers.InvoiceMigrate()
	handlers.CustomerMigrate()
	handlers.UserMigrate()
}

func handleStatus(r *mux.Router) {
	health := r.PathPrefix("/api/").Subrouter()

	health.HandleFunc("/ping", handlers.Status).Methods("GET")

	health.Use(Header)
	health.Use(Logger)
}

func handleAPI(r *mux.Router) {
	api := r.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/invoice", handlers.InvoiceGetAll).Methods("GET")
	api.HandleFunc("/invoice", handlers.InvoiceCreate).Methods("POST")
	api.HandleFunc("/invoice/payment", handlers.InvoiceSetPaid).Methods("POST")
	api.HandleFunc("/invoice/{id:[0-9]+}", handlers.InvoiceGet).Methods("GET")

	api.HandleFunc("/customer", handlers.CustomerGetAll).Methods("GET")
	api.HandleFunc("/customer", handlers.CustomerCreate).Methods("POST")
	api.HandleFunc("/customer/{id:[0-9]+}", handlers.CustomerGet).Methods("GET")
	api.HandleFunc("/customer/{id:[0-9]+}", handlers.CustomerUpdate).Methods("PUT")
	api.HandleFunc("/customer/{id:[0-9]+}", handlers.CustomerDelete).Methods("DELETE")

	api.HandleFunc("/user/{uuid:[a-z0-9-]{36}}", handlers.UserGet).Methods("GET")
	api.HandleFunc("/user/{uuid:[a-z0-9-]{36}}", handlers.UserUpdate).Methods("PUT")

	api.Use(Header)
	api.Use(Logger)
	api.Use(Auth)
}

func handleAuth(r *mux.Router) {
	auth := r.PathPrefix("/auth/").Subrouter()

	auth.HandleFunc("/login", handlers.Login).Methods("POST")
	auth.HandleFunc("/register", handlers.UserCreate).Methods("POST")

	auth.Use(Header)
	auth.Use(Logger)
}

func handleStatic(r *mux.Router) {
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./ui/dist/assets/"))))

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/dist/index.html")
	})

	r.Use(Logger)
}
