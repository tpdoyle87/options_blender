package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Option struct {
	ID          int     `json:"id"`
	Strike      float64 `json:"strike"`
	Expiry      string  `json:"expiry"`
	OptionType  string  `json:"option_type"`
	Underlying  string  `json:"underlying"`
	Credit      float64 `json:"credit"`
	Debit       float64 `json:"debit"`
	Active      bool    `json:"active"`
	Entered     string  `json:"entered"`
	ClosedEarly bool    `json:"closed_early"`
	FinalCredit float64 `json:"final_credit"`
	Notes       string  `json:"notes"`
	UserID      int     `json:"user_id"`
}

// get all options
func GetOptions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM options")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		options := []Option{}
		for rows.Next() {
			var o Option
			if err := rows.Scan(&o.ID, &o.Strike, &o.Expiry, &o.OptionType, &o.Underlying, &o.Credit, &o.Debit, &o.Active, &o.Entered, &o.ClosedEarly, &o.FinalCredit, &o.Notes, &o.UserID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			options = append(options, o)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(options)
	}
}

// get an option by id
func GetOption(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var o Option
		err := db.QueryRow("SELECT * FROM options WHERE id=$1", id).Scan(&o.ID, &o.Strike, &o.Expiry, &o.OptionType, &o.Underlying, &o.Credit, &o.Debit, &o.Active, &o.Entered, &o.ClosedEarly, &o.FinalCredit, &o.Notes, &o.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(o)
	}
}

// create an option
func CreateOption(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o Option
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		_, err := db.Exec("INSERT INTO options (strike, expiry, option_type, underlying, credit, debit, active, entered, closed_early, final_credit, notes, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", o.Strike, o.Expiry, o.OptionType, o.Underlying, o.Credit, o.Debit, o.Active, o.Entered, o.ClosedEarly, o.FinalCredit, o.Notes, o.UserID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// update an option
func UpdateOption(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var o Option
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		_, err := db.Exec("UPDATE options SET strike=$1, expiry=$2, option_type=$3, underlying=$4, credit=$5, debit=$6, active=$7, entered=$8, closed_early=$9, final_credit=$10, notes=$11, user_id=$12 WHERE id=$13", o.Strike, o.Expiry, o.OptionType, o.Underlying, o.Credit, o.Debit, o.Active, o.Entered, o.ClosedEarly, o.FinalCredit, o.Notes, o.UserID, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// delete an option
func DeleteOption(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("DELETE FROM options WHERE id=$1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
