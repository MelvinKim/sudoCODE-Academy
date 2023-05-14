package database_test

import (
	"context"
	"testing"

	"github.com/MelvinKim/courses/domain"
	"github.com/MelvinKim/courses/infrastructure/database"
	"github.com/brianvoe/gofakeit/v6"
)

func TestPostgresDB_CreateCourse(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	course := &domain.Course{
		Title:       gofakeit.FarmAnimal(),
		Price:       gofakeit.UintRange(10, 50),
		Description: "A nice course",
		Instructor:  gofakeit.Name(),
		Category:    gofakeit.CarMaker(),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			course, err := p.CreateCourse(tt.args.ctx, tt.args.course)
			if err != nil != tt.wantErr {
				t.Errorf("PostgresDB.CreateCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && course != nil {
				if course.UUID == "" {
					t.Fatalf("expected course to have a valid UUID")
				}
				if course.CreatedAt == nil {
					t.Fatalf("expected course to have a createdAt timestamp")
				}
				if course.UpdatedAt == nil {
					t.Fatalf("expected course to have a updatedAt timestamp")
				}
				if course.Title == "" {
					t.Fatalf("expected course to have a valid Title")
				}
				if course.Category == "" {
					t.Fatalf("expected course to have a valid category")
				}
				if course.Description == "" {
					t.Fatalf("expected course to have a valid description")
				}
				if course.Instructor == "" {
					t.Fatalf("expected course to have a valid instructor name")
				}
			}
		})
	}
}

func TestPostgresDB_GetCourse(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	course := &domain.Course{
		Title:       gofakeit.LastName(),
		Price:       gofakeit.UintRange(10, 50),
		Description: "A nice course",
		Instructor:  gofakeit.Name(),
		Category:    gofakeit.CarMaker(),
	}
	course, err := p.CreateCourse(ctx, course)
	if err != nil {
		t.Fatalf("error while creating test course: %v", err)
	}
	courseTitle := course.Title

	type args struct {
		ctx         context.Context
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
				courseTitle: &courseTitle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			course, err := p.GetCourse(tt.args.ctx, tt.args.courseTitle)
			if err != nil != tt.wantErr {
				t.Errorf("PostgresDB.GetCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && course != nil {
				if course.UUID == "" {
					t.Fatalf("expected course to have a valid UUID")
				}
				if course.Title == "" {
					t.Fatalf("expected course to have a valid title")
				}
				if course.CreatedAt == nil {
					t.Fatalf("expected course to have a createdAt timestamp")
				}
				if course.UpdatedAt == nil {
					t.Fatalf("expected course to have a updatedAt timestamp")
				}
				if course.Category == "" {
					t.Fatalf("expected course to have a valid category")
				}
				if course.Description == "" {
					t.Fatalf("expected course to have a valid description")
				}
				if course.Instructor == "" {
					t.Fatalf("expected course to have a valid instructor name")
				}
			}
		})
	}
}

func TestPostgresDB_AssignCourseToStudent(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	student := &domain.Student{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	student, err := p.CreateStudent(ctx, student)
	if err != nil {
		t.Fatalf("Failed to test create student: %v", err)
	}
	course := &domain.Course{
		Title:       gofakeit.LastName(),
		Price:       gofakeit.UintRange(10, 50),
		Description: "A nice course",
		Instructor:  gofakeit.Name(),
		Category:    gofakeit.CarMaker(),
	}
	course, err = p.CreateCourse(ctx, course)
	if err != nil {
		t.Errorf("error while creating test course: %v", err)
		return
	}

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.AssignCourseToStudent(tt.args.ctx, tt.args.email, tt.args.courseTitle)
			if err != nil != tt.wantErr {
				t.Errorf("PostgresDB.AssignCourseToStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && course != nil {
				var courses []*domain.Course
				if err := p.DB.Model(course).Association("Students").Find(&courses); err != nil {
					t.Fatalf("Failed to find student's courses: %v", err)
				}
				if len(courses) != 1 {
					t.Fatalf("Expected 1 course, got %d", len(courses))
				}
			}
		})
	}
}
