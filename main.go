package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Junkes887/go-server/handler"
	"github.com/julienschmidt/httprouter"
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
	router := httprouter.New()
	task := handler.Task{DB: db}
	router.GET("/", handler.HelloServer)
	router.GET("/task", task.FindAllTask)
	router.GET("/task/:id", task.FindByIDTask)
	router.POST("/task", task.CreateTask)
	router.PUT("/task", task.UptadeTask)
	http.ListenAndServe(":8080", router)
}
