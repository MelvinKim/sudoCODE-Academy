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
	FirstName    string   `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName     string   `json:"last_name" gorm:"type:varchar(255);not null"`
	Email        string   `json:"email" gorm:"uniqueIndex;not null"`
	Courses      []Course `gorm:"many2many:user_courses;"`
}

// Course ...
type Course struct {
	AbstractBase `gorm:"embedded"`
	Title        string `json:"title" gorm:"uniqueIndex;not null"`
	Price        uint   `json:"price"`
	Description  string `json:"description"`
	Instructor   string `json:"instructor"`
	Category     string `json:"category"`
}

// if You decide to add the studentProfile table, DO NOT FORGET TO INCLUDE IT IN THE MIGRATIONS TO BE RUN
// StudentProfile ...
// type StudentProfile struct {
// 	AbstractBase `gorm:"embedded"`
// 	StudentUUID  string  `json:"student_uuid"`
// 	Student      Student `json:"student,omitempty" gorm:"foreignKey:StudentUUID"`
// 	CourseUUID   string  `json:"course_uuid"`
// 	Course       Course  `json:"course,omitempty" gorm:"foreignKey:CourseUUID"`
// }

// type Course struct {
// 	AbstractBase `gorm:"embedded"`
// 	Title        string  `json:"title"`
// 	Price        uint    `json:"price"`
// 	Description  string  `json:"description"`
// 	Instructor   string  `json:"instructor"`
// 	Category     string  `json:"category"`
// 	StudentUUID  string  `json:"student_uuid"`
// 	Student      Student `json:"student,omitempty" gorm:"foreignKey:StudentUUID"`
// }
