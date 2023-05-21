package main

import (
	"database/sql"
	"encoding/json"
	"net/http"    
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Meeting struct {
	ID          int    `json:"id"`
	Day         string `json:"day"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Picture     int    `json:"picture"`
	Time        string `json:"time"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/meetings", getMeetings).Methods("GET")
	router.HandleFunc("/meetings", createMeeting).Methods("POST")

	http.ListenAndServe(":8000", router)
}

func getMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	meetings := make([]Meeting, 0)

	rows, err := db.Query("SELECT day, name, location, description, picture, time FROM meetings")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var m Meeting
		err := rows.Scan(&m.Day, &m.Name, &m.Location, &m.Description, &m.Picture, &m.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		meetings = append(meetings, m)
	}

	json.NewEncoder(w).Encode(meetings)
}

func createMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m Meeting
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO meetings(day, name, location, description, picture, time) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(m.Day, m.Name, m.Location, m.Description, m.Picture, m.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(m)
}
