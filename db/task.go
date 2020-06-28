package db

import (
	"github.com/Junkes887/go-server/model"
	"github.com/jinzhu/gorm"
)

// FindAllTask busca todos os registros
func FindAllTask(dbConn *gorm.DB) []model.Task {
	var tasks []model.Task
	dbConn.Preload("Status").Find(&tasks)

	return tasks
}

// AddTask adiciona uma tarefa no banco
func AddTask(dbConn *gorm.DB, dto model.Task) {
	dbConn.Create(&dto)
}

// UptadeTask adiciona uma tarefa no banco
func UptadeTask(dbConn *gorm.DB, dto model.Task) {
	dbConn.Save(&dto)
}

// DeleteTask adiciona uma tarefa no banco
func DeleteTask(dbConn *gorm.DB, id string) {
	var task model.Task
	dbConn.Where("id = ?", id).Find(&task)
	dbConn.Delete(&task)
}
