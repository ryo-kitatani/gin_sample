package entity

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Status    string    `json:"status" gorm:"type:enum('todo','doing','done');default:'todo'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

// NewTodo は新しいTodoエンティティを作成します
func NewTodo(title string) *Todo {
	now := time.Now()
	return &Todo{
		Title:     title,
		Status:    "todo", // デフォルト値
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Todo) MarkAsDone() {
	t.Status = "done"
	t.UpdatedAt = time.Now()
}

func (t *Todo) MarkAsDoing() {
	t.Status = "doing"
	t.UpdatedAt = time.Now()
}

func (t *Todo) MarkAsTodo() {
	t.Status = "todo"
	t.UpdatedAt = time.Now()
}

func (t *Todo) Validate() error {
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Status == "" {
		return errors.New("status is required")
	}
	return nil
}
