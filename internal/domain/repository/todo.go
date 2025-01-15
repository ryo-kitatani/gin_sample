package repository

import "gin-sample/internal/domain/entity"

type TodoRepository interface {
	Create(todo *entity.Todo) error
	FindByID(id uint) (*entity.Todo, error)
	FindAll() ([]*entity.Todo, error)
	Update(todo *entity.Todo) error
	Delete(todo uint) error

	// ステータス検索
	FindByStatus(status string) ([]*entity.Todo, error)
}
