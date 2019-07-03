package models

import "time"

// ModelBase declare to override gorm default
type ModelBase struct {
	ID        int       `gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// User table
type User struct {
	ModelBase        // replaces gorm.Model
	Email     string `gorm:"not null; unique"`
	Name      string `gorm:"not null; type:varchar(100)"` // unique_index
	Password  string `gorm:"not null" json:"-"`           // not include in response's json
}
