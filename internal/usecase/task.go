package usecase

import (
	"fmt"
	"github.com/vgekko/go-tasks-graphql/internal/entity"
	"github.com/vgekko/go-tasks-graphql/internal/repository"
)

type TaskUseCase struct {
	taskRepo repository.Task
}

func NewTaskUseCase(taskRepo repository.Task) *TaskUseCase {
	return &TaskUseCase{taskRepo}
}

func (uc *TaskUseCase) Create(input entity.TaskInput) (int, error) {
	id, err := uc.taskRepo.Create(input)
	if err != nil {
		return 0, fmt.Errorf("TasUseCase.Create: %w", err)
	}

	return id, nil
}

func (uc *TaskUseCase) GetByID(id int) (entity.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return entity.Task{}, fmt.Errorf("TaskUseCase.GetByID: %w", err)
	}

	return task, nil
}

func (uc *TaskUseCase) GetAll() ([]entity.Task, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.GetAll: %w", err)
	}

	return tasks, nil
}

func (uc *TaskUseCase) GetOpened() ([]entity.Task, error) {
	openedTasks, err := uc.taskRepo.GetOpened()
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.GetOpened: %w", err)
	}

	return openedTasks, nil
}

func (uc *TaskUseCase) GetCompleted() ([]entity.Task, error) {
	completedTasks, err := uc.taskRepo.GetCompleted()
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.GetCompleted: %w", err)
	}

	return completedTasks, nil
}

func (uc *TaskUseCase) Update(id int, input entity.TaskInput) error {
	if err := uc.taskRepo.Update(id, input); err != nil {
		return fmt.Errorf("TaskUseCase.Update: %w", err)
	}

	return nil
}

func (uc *TaskUseCase) Delete(id int) error {
	if err := uc.taskRepo.Delete(id); err != nil {
		return fmt.Errorf("TaskUseCase.Delete: %w", err)
	}

	return nil
}

func (uc *TaskUseCase) Complete(id int) error {
	if err := uc.taskRepo.Complete(id); err != nil {
		return fmt.Errorf("TaskUseCase.Complete: %w", err)
	}

	return nil
}

func (uc *TaskUseCase) Reopen(id int) error {
	if err := uc.taskRepo.Reopen(id); err != nil {
		return fmt.Errorf("TaskUseCase.Reopen: %w", err)
	}

	return nil
}
