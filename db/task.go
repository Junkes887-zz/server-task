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

// FindByIDTask buscar por id
func FindByIDTask(dbConn *sql.DB, id string) (model.Task, error) {
	sel, err := dbConn.Query("SELECT * FROM task WHERE id=$1", id)
	var task model.Task

	for sel.Next() {
		var id int
		var name, description string
		var status model.Status

		err = sel.Scan(&id, &name, &description, &status)
		if err != nil {
			return task, err
		}

		task = model.Task{
			ID:          id,
			Name:        name,
			Description: description,
			Status:      status,
		}
	}

	return task, nil

}

// UptadeTask adiciona uma tarefa no banco
func UptadeTask(dbConn *sql.DB, dto model.Task) error {
	insForm, err := dbConn.Prepare("UPDATE task SET name = $2, description = $3, status = $4 WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = insForm.Exec(dto.ID, dto.Name, dto.Description, dto.Status)

	if err != nil {
		return err
	}

	return nil
}

// UptadeTask adiciona uma tarefa no banco
func DeleteTask(dbConn *sql.DB, id string) error {
	insForm, err := dbConn.Prepare("delete  from task  where id = $1")
	if err != nil {
		return err
	}

	_, err = insForm.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
