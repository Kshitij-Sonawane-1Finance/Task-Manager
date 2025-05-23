package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAT time.Time	`json:"updated_at,omitempty"`
}

type Admin struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
