package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vgekko/go-tasks-graphql/internal/entity"
)

const tableTasks = "tasks"

type Repository struct {
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Task: NewTaskRepository(db)}
}

type Task interface {
	Create(input entity.TaskInput) (int, error)
	GetByID(id int) (entity.Task, error)
	GetAll() ([]entity.Task, error)
	GetOpened() ([]entity.Task, error)
	GetCompleted() ([]entity.Task, error)
	Update(id int, input entity.TaskInput) error
	Delete(id int) error
	Complete(id int) error
	Reopen(id int) error
}
