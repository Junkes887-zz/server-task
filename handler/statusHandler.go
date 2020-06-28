package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/go-server/db"
	"github.com/Junkes887/go-server/model"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// Status ...
type Status struct {
	DB *gorm.DB
}

// FindAllSatus ...
func (t Status) FindAllSatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	list := db.FindAllStatus(t.DB)

	js, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// CreateStatus ...
func (t Status) CreateStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var p model.Status

	json.NewDecoder(r.Body).Decode(&p)
	db.CreateStatus(t.DB, p)
}

// UpdateStatus ...
func (t Status) UpdateStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var p model.Status

	json.NewDecoder(r.Body).Decode(&p)
	db.UpdateStatus(t.DB, p)
}
