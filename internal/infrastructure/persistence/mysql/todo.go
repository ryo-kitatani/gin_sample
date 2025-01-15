package mysql

import (
	"gin-sample/internal/domain/entity"
	"gin-sample/internal/domain/repository"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository は repository.TodoRepository の MySQL 実装を生成します
func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *entity.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) FindAll() ([]*entity.Todo, error) {
	var todos []*entity.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) FindByID(id uint) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Update(todo *entity.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Todo{}, id).Error
}

func (r *todoRepository) FindByStatus(status string) ([]*entity.Todo, error) {
	var todos []*entity.Todo
	err := r.db.Where("status = ?", status).Find(&todos).Error
	return todos, err
}
