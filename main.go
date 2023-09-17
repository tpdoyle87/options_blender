package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"options/controllers"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS options (id SERIAL PRIMARY KEY, strike FLOAT, expiry TEXT. opton_type TEXT, underlying TEXT, Credit FLOAT, debit FLOAT, active BOOLEAN, entered TEXT, closed_early BOOLEAN, final_credit FLOAT, notes TEXT, user_id INTEGER REFERENCES users(id))`)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")
	router.HandleFunc("/options/", controllers.GetOptions(db)).Methods("GET")
	router.HandleFunc("/options/{id}", controllers.GetOption(db)).Methods("GET")
	router.HandleFunc("/options", controllers.CreateOption(db)).Methods("POST")
	router.HandleFunc("/options/{id}", controllers.UpdateOption(db)).Methods("PUT")
	router.HandleFunc("/options/{id}", controllers.DeleteOption(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
