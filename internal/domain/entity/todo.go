package entity

import (
	"errors"
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Status    string    `json:"status" gorm:"type:enum('todo','doing','done');default:'todo'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
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
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Status == "" {
		return errors.New("status is required")
	}
	return nil
}
