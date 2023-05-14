package repository

import (
	"context"

	"github.com/MelvinKim/courses/domain"
)

// CreateRepository defines create contract
type CreateRepository interface {
	CreateStudent(
		ctx context.Context,
		student *domain.Student,
	) (*domain.Student, error)
	CreateCourse(
		ctx context.Context,
		course *domain.Course,
	) (*domain.Course, error)
	AssignCourseToStudent(
		ctx context.Context,
		email *string,
		courseTitle *string,
	) (*domain.Student, error)
}

// GetRepository defines get contract
type GetRepository interface {
	GetStudent(
		ctx context.Context,
		email *string,
	) (*domain.Student, error)
	GetCourse(
		ctx context.Context,
		title *string,
	) (*domain.Course, error)
}
