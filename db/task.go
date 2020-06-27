package db

import (
	"database/sql"

	"github.com/Junkes887/go-server/model"
)

// AddTask adiciona uma tarefa no banco
func AddTask(dbConn *sql.DB, dto model.TaskDTO) error {
	insForm, err := dbConn.Prepare("INSERT INTO task(name, description, status) VALUES($1,$2,$3);")
	if err != nil {
		return err
	}

	_, err = insForm.Exec(dto.Name, dto.Description, dto.Status)

	if err != nil {
		return err
	}

	return nil
}

// FindAllTask busca todos os registros
func FindAllTask(dbConn *sql.DB) ([]model.Task, error) {
	rows, err := dbConn.Query("SELECT * FROM task")

	list := []model.Task{}

	for rows.Next() {
		var id int
		var name, description string
		var status model.Status

		err = rows.Scan(&id, &name, &description, &status)
		if err != nil {
			return nil, err
		}

		list = append(list, model.Task{
			ID:          id,
			Name:        name,
			Description: description,
			Status:      status,
		})
	}

	return list, nil
}
