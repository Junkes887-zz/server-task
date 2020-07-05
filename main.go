package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/Junkes887/go-server/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var db *gorm.DB

func dbConn() (db *gorm.DB) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		fmt.Errorf("$URL not set")
		return
	}
	db, err := gorm.Open("postgres", dbUrl)
	db.SingularTable(true)

	if err != nil {
		fmt.Println(err.Error)
		panic(err.Error())
	}
	return db
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Errorf("$PORT not set")
		return
	}
	db = dbConn()
	defer db.Close()
	router := httprouter.New()
	task := handler.Task{DB: db}
	status := handler.Status{DB: db}
	router.GET("/", handler.HelloServer)
	router.GET("/status", status.FindAllSatus)
	router.POST("/status", status.CreateStatus)
	router.PUT("/status", status.UpdateStatus)
	router.GET("/task", task.FindAllTask)
	router.POST("/task", task.CreateTask)
	router.PUT("/task", task.UptadeTask)
	router.DELETE("/task/:id", task.DeleteTask)
	c := cors.AllowAll()
	handlerCors := c.Handler(router)
	http.ListenAndServe(":"+port, handlerCors)
	fmt.Println("Server runner in port: " + port)
}
