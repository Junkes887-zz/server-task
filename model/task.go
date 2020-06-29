package model

// Task ...
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status" gorm:"foreignkey:IDStatus"`
	IDStatus    int    `gorm:"column:id_status"`
}
