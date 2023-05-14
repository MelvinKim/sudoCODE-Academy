package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/MelvinKim/courses/domain"
	"github.com/MelvinKim/courses/repository"
)

type UsecaseContract interface {
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
	GetStudent(
		ctx context.Context,
		email *string,
	) (*domain.Student, error)
	GetCourse(
		ctx context.Context,
		title *string,
	) (*domain.Course, error)
}

// Usecase represents the Courses's service business logic
type Usecase struct {
	Create repository.CreateRepository
	Get    repository.GetRepository
}

// Checkpreconditions asserts all pre-conditions are met
func (u *Usecase) Checkpreconditions() {
	if u.Create == nil {
		log.Panicf("courses usecase has not initialized a create repository")
	}
	if u.Get == nil {
		log.Panicf("courses usecase has not initialized a get repository")
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

// CreateCourse creates a new sudocode course
func (u *Usecase) CreateCourse(
	ctx context.Context,
	course *domain.Course,
) (*domain.Course, error) {
	if course.Description == "" {
		return nil, fmt.Errorf("course's description can not be empty")
	}
	if course.Instructor == "" {
		return nil, fmt.Errorf("course's instructor can not be empty")
	}
	if course.Price == 0 {
		return nil, fmt.Errorf("course's price can not be zero")
	}
	if course.Title == "" {
		return nil, fmt.Errorf("course's title can not be empty")
	}
	if course.Category == "" {
		return nil, fmt.Errorf("course's category can not be empty")
	}
	return u.Create.CreateCourse(ctx, course)
}

// AssignCourseToStudent assign a student a course
func (u *Usecase) AssignCourseToStudent(
	ctx context.Context,
	email *string,
	courseTitle *string,
) (*domain.Student, error) {
	if *email == "" {
		return nil, fmt.Errorf("student's email can not be empty")
	}
	if *courseTitle == "" {
		return nil, fmt.Errorf("course's title can not be empty")
	}
	return u.Create.AssignCourseToStudent(ctx, email, courseTitle)
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

// GetCourse gets a sudoCODE academy course based on the course title
func (u *Usecase) GetCourse(
	ctx context.Context,
	title *string,
) (*domain.Course, error) {
	if title == nil {
		return nil, fmt.Errorf("course title can not be empty")
	}
	return u.Get.GetCourse(ctx, title)
}
