package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdated:CASCADE,OnDelete:CASCADE" json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
