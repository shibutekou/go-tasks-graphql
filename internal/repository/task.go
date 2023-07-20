package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vgekko/go-tasks-graphql/internal/entity"
)

const (
	statusCompleted = "completed"
	statusOpened    = "opened"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(input entity.TaskInput) (int, error) {
	var taskID int

	queryInsert := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", tableTasks)

	err := r.db.QueryRow(queryInsert, input.Title, input.Description).Scan(&taskID)
	if err != nil {
		return 0, fmt.Errorf("TaskRepository.Create: %w", err)
	}

	return taskID, nil
}

func (r *TaskRepository) GetByID(id int) (entity.Task, error) {
	var task entity.Task

	querySelectByID := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", tableTasks)
	if err := r.db.Get(&task, querySelectByID, id); err != nil {
		return entity.Task{}, fmt.Errorf("TaskRepository.GetByID: %w", err)
	}

	return task, nil
}

func (r *TaskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task

	querySelectAll := fmt.Sprintf("SELECT * FROM %s", tableTasks)
	if err := r.db.Select(&tasks, querySelectAll); err != nil {
		return nil, fmt.Errorf("TaskRepository.GetAll: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) GetOpened() ([]entity.Task, error) {
	var tasks []entity.Task

	querySelectOpened := fmt.Sprintf("SELECT * FROM %s WHERE status='%s'", tableTasks, statusOpened)
	if err := r.db.Select(&tasks, querySelectOpened); err != nil {
		return nil, fmt.Errorf("TaskRepository.GetOpened: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) GetCompleted() ([]entity.Task, error) {
	var tasks []entity.Task

	querySelectCompleted := fmt.Sprintf("SELECT * FROM %s WHERE status='%s'", tableTasks, statusCompleted)
	if err := r.db.Select(&tasks, querySelectCompleted); err != nil {
		return nil, fmt.Errorf("TaskRepository.GetCompleted: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) Update(id int, input entity.TaskInput) error {
	queryUpdate := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 FROM %s WHERE id=$3 ", tableTasks)

	_, err := r.db.Exec(queryUpdate, input.Title, input.Description, id)
	if err != nil {
		return fmt.Errorf("TaskRepository.Update: %w", err)
	}

	return nil
}

func (r *TaskRepository) Delete(id int) error {
	queryDelete := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableTasks)

	_, err := r.db.Exec(queryDelete, id)
	if err != nil {
		return fmt.Errorf("TaskRepository.Delete: %w", err)
	}

	return nil
}

func (r *TaskRepository) Complete(id int) error {
	queryUpdate := fmt.Sprintf("UPDATE %s SET status='%s' WHERE id=$1", tableTasks, statusCompleted)

	_, err := r.db.Exec(queryUpdate, id)
	if err != nil {
		return fmt.Errorf("TaskRepository.Complete: %w", err)
	}

	return nil
}

func (r *TaskRepository) Reopen(id int) error {
	queryUpdate := fmt.Sprintf("UPDATE %s SET status='%s' WHERE id=$1", tableTasks, statusOpened)

	_, err := r.db.Exec(queryUpdate, id)
	if err != nil {
		return fmt.Errorf("TaskRepository.Reopen: %w", err)
	}

	return nil
}
