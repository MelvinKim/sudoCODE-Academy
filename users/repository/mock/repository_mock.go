package mock

import (
	"context"

	"github.com/MelvinKim/users/domain"
)

// MockCreateRepository mocks the database create repository
type MockCreateRepository struct {
	MockCreateStudent func(
		ctx context.Context,
		student *domain.Student,
	) (*domain.Student, error)
}

// NewMockCreateRepository initializes a new MockCreateRepository
func NewMockCreateRepository() *MockCreateRepository {
	return &MockCreateRepository{
		MockCreateStudent: func(ctx context.Context, student *domain.Student) (*domain.Student, error) {
			return &domain.Student{}, nil
		},
	}
}

// CreateStudent mocks CreateStudent
func (c *MockCreateRepository) CreateStudent(
	ctx context.Context,
	student *domain.Student,
) (*domain.Student, error) {
	return c.MockCreateStudent(ctx, student)
}

// MockGetRepository mocks the database get repository
type MockGetRepository struct {
	MockGetStudent func(
		ctx context.Context,
		email *string,
	) (*domain.Student, error)
}

// NewMockGetRepository initializes a new MockGetRepository
func NewMockGetRepository() *MockGetRepository {
	return &MockGetRepository{
		MockGetStudent: func(ctx context.Context, email *string) (*domain.Student, error) {
			return &domain.Student{}, nil
		},
	}
}

// GetStudent mocks GetStudent
func (c *MockGetRepository) GetStudent(
	ctx context.Context,
	email *string,
) (*domain.Student, error) {
	return c.MockGetStudent(ctx, email)
}
