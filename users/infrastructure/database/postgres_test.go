package database_test

import (
	"context"
	"testing"

	"github.com/MelvinKim/users/domain"
	"github.com/MelvinKim/users/infrastructure/database"
	"github.com/brianvoe/gofakeit/v6"
)

func TestPostgresDB_CreateStudent(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	student := &domain.Student{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student, err := p.CreateStudent(tt.args.ctx, tt.args.student)
			if err != nil != tt.wantErr {
				t.Errorf("PostgresDB.CreateStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && student != nil {
				if student.UUID == "" {
					t.Fatalf("expected student to have a valid UUID")
				}
				if student.CreatedAt == nil {
					t.Fatalf("expected student to have a createdAt timestamp")
				}
				if student.UpdatedAt == nil {
					t.Fatalf("expected student to have a updatedAt timestamp")
				}
				if student.Email == "" {
					t.Fatalf("expected student to have a valid email")
				}
				if student.FirstName == "" {
					t.Fatalf("expected student to have a valid first name")
				}
				if student.LastName == "" {
					t.Fatalf("expected student to have a valid last name")
				}
			}
		})
	}
}

func TestPostgresDB_GetStudent(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	student := &domain.Student{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	student, err := p.CreateStudent(ctx, student)
	if err != nil {
		t.Errorf("error while creating test student: %v", err)
		return
	}
	randomEmail := gofakeit.Email()

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
		{
			name: "Sad case - random email",
			args: args{
				ctx:   ctx,
				email: &randomEmail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student, err := p.GetStudent(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.name == "Sad case - random email" && student != nil {
					t.Fatalf("expcted not student to be found, byt got %v", student)
				}
				if tt.name == "Happy case" {
					if student.Email == "nil" {
						t.Fatalf("expected student to have a valid Email")
					}
					if student.FirstName == "" {
						t.Fatalf("expected student to have a valid FirstName")
					}
					if student.CreatedAt == nil {
						t.Fatalf("expected student to have a valid created at timestamp")
					}
					if student.UpdatedAt == nil {
						t.Fatalf("expected student to have a valid created at timestamp")
					}
				}
			}
		})
	}
}
