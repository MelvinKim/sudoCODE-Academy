package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/MelvinKim/users/domain"
	"github.com/MelvinKim/users/repository"
)

type UsecaseContract interface {
	CreateStudent(
		ctx context.Context,
		student *domain.Student,
	) (*domain.Student, error)
	GetStudent(
		ctx context.Context,
		email *string,
	) (*domain.Student, error)
}

// Usecase represents the User's service business logic
type Usecase struct {
	Create repository.CreateRepository
	Get    repository.GetRepository
}

// Checkpreconditions asserts all pre-conditions are met
func (u *Usecase) Checkpreconditions() {
	if u.Create == nil {
		log.Panicf("users usecase has not initialized a create repository")
	}
	if u.Get == nil {
		log.Panicf("users usecase has not initialized a get repository")
	}
}

// NewUsecase creates a new usecase instance
func NewUsecase(
	create repository.CreateRepository,
	get repository.GetRepository,
) *Usecase {
	uc := &Usecase{
		Create: create,
		Get:    get,
	}
	uc.Checkpreconditions()
	return uc
}

// CreateStudent creates a new sudocode student
func (u *Usecase) CreateStudent(
	ctx context.Context,
	student *domain.Student,
) (*domain.Student, error) {
	if student.Email == "" {
		return nil, fmt.Errorf("email can not be empty")
	}
	if student.FirstName == "" {
		return nil, fmt.Errorf("first name can not be empty")
	}
	if student.LastName == "" {
		return nil, fmt.Errorf("last name can not be empty")
	}
	return u.Create.CreateStudent(ctx, student)
}

// GetStudent gets student ny their email address
func (u *Usecase) GetStudent(
	ctx context.Context,
	email *string,
) (*domain.Student, error) {
	if email == nil {
		return nil, fmt.Errorf("email can not be empty")
	}
	return u.Get.GetStudent(ctx, email)
}
