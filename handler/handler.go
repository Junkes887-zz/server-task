package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Junkes887/go-server/db"
	"github.com/Junkes887/go-server/model"
	"github.com/julienschmidt/httprouter"
)

// HelloServer ...
func HelloServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Server run port 8080..."))
}

// Task ...
type Task struct {
	DB *sql.DB
}

// CreateTask ...
func (t Task) CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var p model.TaskDTO

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.AddTask(t.DB, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// FindAllTask ...
func (t Task) FindAllTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list, err := db.FindAllTask(t.DB)

	js, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// FindByIDTask ...
func (t Task) FindByIDTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	list, err := db.FindByIDTask(t.DB, id)

	js, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UptadeTask ...
func (t Task) UptadeTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var p model.Task

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UptadeTask(t.DB, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// DeleteTask ...
func (t Task) DeleteTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	err := db.DeleteTask(t.DB, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
