package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Container struct {
	ID             int       `json:"id"`
	IPAddress      string    `json:"ip_address"`
	PingTime       float64   `json:"ping_time"`
	LastSuccessful time.Time `json:"last_successful"`
}

var db *sql.DB

func initDB() {
	time.Sleep(5 * time.Second)
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!")
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS containers (
        id SERIAL PRIMARY KEY,
        ip_address TEXT NOT NULL,
        ping_time FLOAT8,
        last_successful TIMESTAMP
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Обработка POST-запроса (Добавление данных)
func addContainer(w http.ResponseWriter, r *http.Request) {
	var container Container
	err := json.NewDecoder(r.Body).Decode(&container)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO containers (ip_address, ping_time, last_successful) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, container.IPAddress, container.PingTime, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Обработка GET-запроса (Получение данных)
func getContainers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, ip_address, ping_time, last_successful FROM containers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var containers []Container
	for rows.Next() {
		var c Container
		err := rows.Scan(&c.ID, &c.IPAddress, &c.PingTime, &c.LastSuccessful)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		containers = append(containers, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}

func main() {
	initDB()
	createTable()

	http.HandleFunc("/add", addContainer)
	http.HandleFunc("/containers", getContainers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
