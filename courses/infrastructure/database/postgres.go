package database

import (
	"context"
	"fmt"
	"os"

	"github.com/MelvinKim/courses/domain"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB sets up a database layer within the service
type PostgresDB struct {
	DB *gorm.DB
}

// Checkpreconditions assert all conditions required to run the service are met
func (p *PostgresDB) Checkpreconditions() {
	if p.DB == nil {
		log.Fatalf("postgres database ORM has not been initialized.")
	}
}

// NewPostgresDB initializes a new postgres DB instance
func NewPostgresDB() *PostgresDB {
	db := PostgresDB{
		DB: Init(),
	}
	db.Checkpreconditions()
	return &db
}

// Migrate runs the databas's migrations
func Migrate(db *gorm.DB) {
	tables := []interface{}{
		&domain.Student{},
		&domain.Course{},
	}
	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			log.Panicf("can't run db migrations on table %v in the user's service: err: %v", table, err)
		}
	}
}

// Init initializes a new gorm instance by connecting to a postgres DB instance
func Init() *gorm.DB {
	var dsn string
	if os.Getenv("ENVIRONMENT") == "prod" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("TEST_DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't open postgres db connection for the courses service: %v", err)
	}
	log.Info("Database connected successfully.")
	Migrate(db)
	log.Info("Database migrations ran successfully.")
	return db
}

// CreateCourse creates a new course in sudocode acaddemy
func (p *PostgresDB) CreateCourse(
	ctx context.Context,
	course *domain.Course,
) (*domain.Course, error) {
	if err := p.DB.Create(course).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new course: %v", err)
	}
	return course, nil
}

// GetCourse returns a single course
func (p *PostgresDB) GetCourse(
	ctx context.Context,
	title *string,
) (*domain.Course, error) {
	filters := &domain.Course{
		Title: *title,
	}
	var course domain.Course
	if err := p.DB.Where(filters).Find(&course).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't get course by title: %v err: %v", title, err)
	}
	if course.UUID == "" {
		return nil, nil
	}

	return &course, nil
}

// AssignCourseToStudent assigns a course to a student after they have purchased them
func (p *PostgresDB) AssignCourseToStudent(
	ctx context.Context,
	email *string,
	courseTitle *string,
) (*domain.Student, error) {
	student, err := p.GetStudent(ctx, email)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, fmt.Errorf("kindly create an account with us in order to purchase a course") // maybe to do some redirect to the signup page
	}

	course, err := p.GetCourse(ctx, courseTitle)
	if err != nil {
		return nil, err
	}
	for _, c := range student.Courses {
		if c.UUID == course.UUID {
			return nil, fmt.Errorf("course %v is already assigned to student %v", c.UUID, student.UUID)
		}
	}
	student.Courses = append(student.Courses, *course)
	if err = p.DB.Save(student).Error; err != nil {
		return nil, fmt.Errorf("error while assigning a student course with UUID %v, title %v", course.UUID, course.Title)
	}

	return student, nil
}

// GetStudent returns a single student
func (p *PostgresDB) GetStudent(
	ctx context.Context,
	email *string,
) (*domain.Student, error) {
	filters := &domain.Student{
		Email: *email,
	}
	var student domain.Student
	if err := p.DB.Where(filters).Find(&student).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't get student by email: %v err: %v", email, err)
	}
	if student.UUID == "" {
		return nil, nil
	}

	return &student, nil
}

// CreateStudentProfile creates a new student profile in sudoCODE academy
// func (p *PostgresDB) CreateStudentProfile(
// 	ctx context.Context,
// 	studentProfile *domain.StudentProfile,
// ) (*domain.StudentProfile, error) {
// 	if err := p.DB.Create(studentProfile).Error; err != nil {
// 		return nil, fmt.Errorf("infrastructure: can't create a new student profile: %v", err)
// 	}
// 	return studentProfile, nil
// }

// // GetStudentProfile returns a Student's Profile
// func (p *PostgresDB) GetStudentProfile(
// 	ctx context.Context,
// 	studentUUID *string,
// ) (*domain.StudentProfile, error) {
// 	filters := &domain.StudentProfile{
// 		StudentUUID: *studentUUID,
// 	}
// 	var studentProfile domain.StudentProfile
// 	if err := p.DB.Where(filters).Find(&studentProfile).Error; err != nil {
// 		return nil, fmt.Errorf("infrastructure: can't get student profile with UUID %v err: %v", studentUUID, err)
// 	}
// 	if studentProfile.UUID == "" {
// 		return nil, nil
// 	}

// 	return &studentProfile, nil
// }