package usecase_test

import (
	"context"
	"testing"

	"github.com/MelvinKim/courses/domain"
	"github.com/MelvinKim/courses/infrastructure/database"
	course "github.com/MelvinKim/courses/usecase"
	"github.com/brianvoe/gofakeit/v6"
)

func newTestUsecase() *course.Usecase {
	create := database.NewPostgresDB()
	get := database.NewPostgresDB()
	u := course.NewUsecase(create, get)
	return u
}

func TestUsecase_CreateStudent(t *testing.T) {
	u := newTestUsecase()
	ctx := context.Background()
	student := &domain.Student{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	incompleteStudent := &domain.Student{
		FirstName: "",
		LastName:  "",
		Email:     "",
	}

	type args struct {
		ctx     context.Context
		student *domain.Student
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:     ctx,
				student: student,
			},
			wantErr: false,
		},
		{
			name: "Sad case - missing fields",
			args: args{
				ctx:     ctx,
				student: incompleteStudent,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student, err := u.CreateStudent(ctx, student)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && student != nil {
				if student.UUID == "" {
					t.Fatalf("expected guest profile to have a valid UUID.")
				}
				if student.CreatedAt == nil {
					t.Fatalf("expected guest profile to have a created at timestamp.")
				}
				if student.UpdatedAt == nil {
					t.Fatalf("expected guest profile to have an updated at timestamp.")
				}
			}
		})
	}
}

func TestUsecase_CreateCourse(t *testing.T) {
	u := newTestUsecase()
	ctx := context.Background()
	course := &domain.Course{
		Title:       gofakeit.LastName(),
		Price:       23,
		Description: gofakeit.Address().City,
		Instructor:  gofakeit.Name(),
		Category:    gofakeit.Car().Brand,
	}
	incompleteCourse := &domain.Course{
		Title:       "",
		Description: "",
		Instructor:  "",
		Price:       0,
		Category:    "",
	}

	type args struct {
		ctx    context.Context
		course *domain.Course
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:    ctx,
				course: course,
			},
			wantErr: false,
		},
		{
			name: "Sad case - missing fields",
			args: args{
				ctx:    ctx,
				course: incompleteCourse,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			course, err := u.CreateCourse(tt.args.ctx, tt.args.course)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && course != nil {
				if course.UUID == "" {
					t.Fatalf("expected course to have a valid UUID.")
				}
				if course.CreatedAt == nil {
					t.Fatalf("expected course to have a created at timestamp.")
				}
				if course.UpdatedAt == nil {
					t.Fatalf("expected course to have an updated at timestamp.")
				}
			}
		})
	}
}

func TestUsecase_AssignCourseToStudent(t *testing.T) {
	u := newTestUsecase()
	ctx := context.Background()
	student := &domain.Student{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	student, err := u.CreateStudent(ctx, student)
	if err != nil {
		t.Errorf("error while creating test student, err: %v", err)
	}
	course := &domain.Course{
		Title:       gofakeit.LastName(),
		Price:       23,
		Description: gofakeit.Address().City,
		Instructor:  gofakeit.Name(),
		Category:    gofakeit.Car().Brand,
	}
	course, err = u.CreateCourse(ctx, course)
	if err != nil {
		t.Errorf("error while creating test course, err: %v", err)
	}
	emptyEmail := ""
	emptyCourseTitle := ""

	type args struct {
		ctx         context.Context
		email       *string
		courseTitle *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:         ctx,
				email:       &student.Email,
				courseTitle: &course.Title,
			},
			wantErr: false,
		},
		{
			name: "Sad case - empty email and title",
			args: args{
				ctx:         ctx,
				email:       &emptyEmail,
				courseTitle: &emptyCourseTitle,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student, err := u.AssignCourseToStudent(tt.args.ctx, tt.args.email, tt.args.courseTitle)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AssignCourseToStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && course != nil {
				if student.UUID == "" {
					t.Fatalf("expected student to have a valid UUID.")
				}
				if student.CreatedAt == nil {
					t.Fatalf("expected student to have a created at timestamp.")
				}
				if student.UpdatedAt == nil {
					t.Fatalf("expected student to have an updated at timestamp.")
				}
			}
		})
	}
}
