package todo

import (
	"time"

	"github.com/google/uuid"
)

type (
	Todo struct {
		ID          uuid.UUID `form:"id" json:"id" binding:"omitempty"`
		Title       string    `form:"title" json:"title" binding:"omitempty"`
		Description string    `form:"description" json:"description" binding:"omitempty"`
		DueDate     time.Time `form:"due_date" json:"due_date" binding:"omitempty" time_format:"2006-01-02T15:04:05Z"`
		Priority    string    `form:"priority" json:"priority" binding:"omitempty,oneof=LOW MEDIUM HIGH"`
		Completed   bool      `form:"completed" json:"completed" binding:"omitempty"`
	}

	CreateTodo struct {
		ID          uuid.UUID `form:"id" json:"id" binding:"omitempty"`
		Title       string    `form:"title" json:"title" binding:"omitempty"`
		Description string    `form:"description" json:"description" binding:"omitempty"`
		DueDate     time.Time `form:"due_date" json:"due_date" binding:"omitempty" time_format:"2006-01-02T15:04:05Z"`
		Priority    string    `form:"priority" json:"priority" binding:"omitempty,oneof=LOW MEDIUM HIGH"`
	}

	UpdateTodo struct {
		Title       string    `form:"title" json:"title" binding:"omitempty"`
		Description string    `form:"description" json:"description" binding:"omitempty"`
		DueDate     time.Time `form:"due_date" json:"due_date" binding:"omitempty" time_format:"2006-01-02T15:04:05Z"`
		Priority    string    `form:"priority" json:"priority" binding:"omitempty,oneof=LOW MEDIUM HIGH"`
	}

	SearchTodoReq struct {
		Todo
		PaginateReq
	}

	SearchTodoRes struct {
		Items []CreateTodo `json:"items"`
		PaginateReq
		Total int64 `json:"total"`
	}

	CompletedReq struct {
		Completed bool `json:"completed"`
	}

	PaginateReq struct {
		Start int `form:"start" json:"start" binding:"min=0"`
		Limit int `form:"limit" json:"limit" binding:"required,min=2,max=100"`
	}
)
