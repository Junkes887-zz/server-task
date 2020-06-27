package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Junkes887/go-server/handler"
)

var db *sql.DB

func dbConn() (db *sql.DB) {
	dbDriver := "postgres"
	dbUser := "postgres"
	dbPass := "go"
	dbName := "postgres"

	db, err := sql.Open(dbDriver, "postgres://"+dbUser+":"+dbPass+"@localhost:5433/"+dbName+"?sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db = dbConn()
	defer db.Close()
	task := handler.Task{DB: db}
	http.HandleFunc("/", handler.HelloServer)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			task.CreateTask(w, r)

		case "GET":
			task.FindAllTask(w, r)
		}
	})
	http.ListenAndServe(":8080", nil)
}
