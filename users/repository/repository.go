package repository

import (
	"context"

	"github.com/MelvinKim/users/domain"
)

// CreateRepository defines create contract
type CreateRepository interface {
	CreateStudent(
		ctx context.Context,
		student *domain.Student,
	) (*domain.Student, error)
}

// GetRepository defines the get contract
type GetRepository interface {
	GetStudent(
		ctx context.Context,
		email *string,
	) (*domain.Student, error)
}
