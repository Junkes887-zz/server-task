package db

import (
	"fmt"

	"github.com/Junkes887/go-server/model"
	"github.com/jinzhu/gorm"
)

// FindAllStatus busca todos os registros
func FindAllStatus(dbConn *gorm.DB) []model.Status {
	var list []model.Status
	dbConn.Preload("Tasks").Find(&list)

	fmt.Println(list)
	return list
}

// CreateStatus ..
func CreateStatus(dbConn *gorm.DB, dto model.Status) {
	dbConn.Create(&dto)
}

// UpdateStatus ..
func UpdateStatus(dbConn *gorm.DB, dto model.Status) {
	dbConn.Save(&dto)
}
