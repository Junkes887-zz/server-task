package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/go-server/db"
	"github.com/Junkes887/go-server/model"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// HelloServer ...
func HelloServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Server run port 8080..."))
}

// Task ...
type Task struct {
	DB *gorm.DB
}

// FindAllTask ...
func (t Task) FindAllTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	list := db.FindAllTask(t.DB)

	js, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// CreateTask ...
func (t Task) CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) model.Task {
	var p model.Task

	json.NewDecoder(r.Body).Decode(&p)
	db.AddTask(t.DB, p)
	return p
}

// UptadeTask ...
func (t Task) UptadeTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var p model.Task

	json.NewDecoder(r.Body).Decode(&p)
	db.UptadeTask(t.DB, p)
}

// DeleteTask ...
func (t Task) DeleteTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	db.DeleteTask(t.DB, id)
}
