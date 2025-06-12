package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type DateOnly struct {
	time.Time
}

const layoutISO = "2006-01-02"

// For JSON serialization
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // remove quotes
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(layoutISO))
}

// For GORM to store in database
func (d DateOnly) Value() (driver.Value, error) {
	return d.Time, nil
}

// For GORM to read from database
func (d *DateOnly) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid time format")
	}
	d.Time = t
	return nil
}

type Task struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	UserID      uint      `gorm:"not null;index" json:"-"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdated:CASCADE,OnDelete:CASCADE" json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   DateOnly  `json:"start_date"`
	EndDate     DateOnly  `json:"end_date"`
	Status      string    `gorm:"type:text;default:'Pending'" json:"status,omitempty"`
	Priority    string    `gorm:"type:text;default:'Medium'" json:"priority,omitempty"`
	Active      bool      `gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
