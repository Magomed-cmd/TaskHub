package model

import "github.com/google/uuid"

type User struct {
	ID       uint64    `gorm:"primaryKey"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `gorm:"type:text;not null" json:"Password,omitempty"`
	Name     string    `gorm:"type:varchar(100);not null"`
	UUID     uuid.UUID `gorm:"type:uuid;not null"`
}
