package dto

import (
	"time"

	"github.com/kshitij/taskManager/models"
)

type Task struct {
	ID          uint      `json:"id,omitempty"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Active      bool      `json:"-"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UpdateTask struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	StartDate   models.DateOnly `json:"start_date"`
	EndDate     models.DateOnly `json:"end_date"`
	Status      string          `json:"status"`
	Priority    string          `json:"priority"`
	UpdatedAt   time.Time       `json:"-"`
}
