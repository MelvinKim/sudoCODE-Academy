package usecase_test

import (
	"context"
	"testing"

	"github.com/MelvinKim/users/domain"
	"github.com/MelvinKim/users/infrastructure/database"
	student "github.com/MelvinKim/users/usecase"
	"github.com/brianvoe/gofakeit/v6"
)

// var (
// 	mockCreate = mock.NewMockCreateRepository()
// 	mockGet    = mock.NewMockGetRepository()
// )

// newTestUseCase initializes a new test Usecase
func newTestUsecase() *student.Usecase {
	create := database.NewPostgresDB()
	get := database.NewPostgresDB()
	u := student.NewUsecase(create, get)
	return u
}

// newMockTestUseCase
// func newMockTestUseCase() *student.Usecase {
// 	mockUsecase := student.NewUsecase(mockCreate, mockGet)
// 	return mockUsecase
// }

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

func TestUsecase_GetStudent(t *testing.T) {
	u := newTestUsecase()
	ctx := context.Background()
	student := &domain.Student{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	student, err := u.CreateStudent(ctx, student)
	if err != nil {
		t.Errorf("error while creating test user: %v", err)
		return
	}

	type args struct {
		ctx   context.Context
		email *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:   ctx,
				email: &student.Email,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student, err := u.GetStudent(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetStudent() error = %v, wantErr %v", err, tt.wantErr)
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
