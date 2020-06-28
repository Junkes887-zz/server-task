package model

// Status ...
type Status struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `gorm:"foreignkey:IDStatus"`
}
