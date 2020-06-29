package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Junkes887/go-server/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var db *gorm.DB

func dbConn() (db *gorm.DB) {

	dsn := url.URL{
		User:     url.UserPassword("postgres", "go"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", "localhost", 5433),
		Path:     "postgres",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open("postgres", dsn.String())
	db.SingularTable(true)

	if err != nil {
		fmt.Println(err.Error)
		panic(err.Error())
	}
	return db
}

func main() {
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
	http.ListenAndServe(":3333", handlerCors)
}
