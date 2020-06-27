package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Junkes887/go-server/db"
	"github.com/Junkes887/go-server/model"
)

// HelloServer ...
func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Server run port 8080..."))
}

// Task ...
type Task struct {
	DB *sql.DB
}

// CreateTask ...
func (t Task) CreateTask(w http.ResponseWriter, r *http.Request) {
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
func (t Task) FindAllTask(w http.ResponseWriter, r *http.Request) {
	list, err := db.FindAllTask(t.DB)

	js, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
