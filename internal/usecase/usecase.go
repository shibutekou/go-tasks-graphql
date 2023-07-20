package usecase

import (
	"github.com/vgekko/go-tasks-graphql/internal/entity"
	"github.com/vgekko/go-tasks-graphql/internal/repository"
)

type Usecase struct {
	Task
}

func NewUseCase(repo *repository.Repository) *Usecase {
	return &Usecase{Task: NewTaskUseCase(repo.Task)}
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
