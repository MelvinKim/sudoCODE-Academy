package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AbstractBase is an abstract struct that can be embedded in other structs
type AbstractBase struct {
	UUID      string `gorm:"primaryKey"`
	Active    bool   `gorm:"default:true"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate ensures a UUID and createdAt data is inserted
func (ab *AbstractBase) BeforeCreate(tx *gorm.DB) (err error) {
	ab.UUID = uuid.New().String()
	return
}

// Student ...
type Student struct {
	AbstractBase `gorm:"embedded"`
	FirstName    string `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName     string `json:"last_name" gorm:"type:varchar(255);not null"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
}
