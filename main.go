package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"options/controllers/users"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", users.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", users.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", users.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", users.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", users.DeleteUser(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
