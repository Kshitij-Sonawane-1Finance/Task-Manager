package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAT time.Time
}

type Admin struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Email string `gorm:"unique"`
	Password string
}
