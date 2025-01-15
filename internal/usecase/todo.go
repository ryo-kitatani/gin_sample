package usecase

import (
	"gin-sample/internal/domain/entity"
	"gin-sample/internal/domain/repository"
)

type TodoUseCase interface {
	Create(title string) error
	GetAll() ([]*entity.Todo, error)
	GetByID(id uint) (*entity.Todo, error)
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
	GetByStatus(status string) ([]*entity.Todo, error)
}

type todoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}

// Create 新規作成
func (tu *todoUseCase) Create(title string) error {
	todo := &entity.Todo{
		Title: title,
	}

	if err := todo.Validate(); err != nil {
		return err
	}

	return tu.todoRepo.Create(todo)
}

func (u *todoUseCase) GetAll() ([]*entity.Todo, error) {
	return u.todoRepo.FindAll()
}

func (u *todoUseCase) GetByID(id uint) (*entity.Todo, error) {
	return u.todoRepo.FindByID(id)
}

func (u *todoUseCase) UpdateStatus(id uint, status string) error {
	todo, err := u.todoRepo.FindByID(id)
	if err != nil {
		return err
	}

	switch status {
	case "todo":
		todo.MarkAsTodo()
	case "doing":
		todo.MarkAsDoing()
	case "done":
		todo.MarkAsDone()
	}

	if err := todo.Validate(); err != nil {
		return err
	}

	return u.todoRepo.Update(todo)
}

func (u *todoUseCase) Delete(id uint) error {
	return u.todoRepo.Delete(id)
}

func (u *todoUseCase) GetByStatus(status string) ([]*entity.Todo, error) {
	return u.todoRepo.FindByStatus(status)
}
