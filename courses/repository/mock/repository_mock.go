package mock

import (
	"context"

	"github.com/MelvinKim/courses/domain"
)

// MockCreateRepository mocks the database create repository
type MockCreateRepository struct {
	MockCreateStudent func(
		ctx context.Context,
		student *domain.Student,
	) (*domain.Student, error)
	MockCreateCourse func(
		ctx context.Context,
		course *domain.Course,
	) (*domain.Course, error)
	MockAssignCourseToStudent func(
		ctx context.Context,
		email *string,
		courseTitle *string,
	) (*domain.Student, error)
}

// NewMockCreateRepository initializes a new MockCreateRepository
func NewMockCreateRepository() *MockCreateRepository {
	return &MockCreateRepository{
		MockCreateStudent: func(ctx context.Context, student *domain.Student) (*domain.Student, error) {
			return &domain.Student{}, nil
		},
		MockCreateCourse: func(ctx context.Context, course *domain.Course) (*domain.Course, error) {
			return &domain.Course{}, nil
		},
		MockAssignCourseToStudent: func(ctx context.Context, email, courseTitle *string) (*domain.Student, error) {
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

// CreateCourse mocks CreateCourse
func (c *MockCreateRepository) CreateCourse(
	ctx context.Context,
	course *domain.Course,
) (*domain.Course, error) {
	return c.MockCreateCourse(ctx, course)
}

// AssignCourseToStudent mocks AssignCourseToStudent
func (c *MockCreateRepository) AssignCourseToStudent(
	ctx context.Context,
	email *string,
	courseTitle *string,
) (*domain.Student, error) {
	return c.MockAssignCourseToStudent(ctx, email, courseTitle)
}

// MockGetRepository mocks the database get repository
type MockGetRepository struct {
	MockGetStudent func(
		ctx context.Context,
		email *string,
	) (*domain.Student, error)
	MockGetCourse func(
		ctx context.Context,
		title *string,
	) (*domain.Course, error)
}

// NewMockGetRepository initializes a new MockGetRepository
func NewMockGetRepository() *MockGetRepository {
	return &MockGetRepository{
		MockGetStudent: func(ctx context.Context, email *string) (*domain.Student, error) {
			return &domain.Student{}, nil
		},
		MockGetCourse: func(ctx context.Context, title *string) (*domain.Course, error) {
			return &domain.Course{}, nil
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

// GetCourse mocks GetCourse
func (c *MockGetRepository) GetCourse(
	ctx context.Context,
	title *string,
) (*domain.Course, error) {
	return c.MockGetCourse(ctx, title)
}
