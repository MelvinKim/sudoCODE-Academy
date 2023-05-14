package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MelvinKim/courses/application/common/dto"
	"github.com/MelvinKim/courses/domain"
	"github.com/MelvinKim/courses/presentation/interactor"
)

// PresentationHandlers represents all the REST API logic
type PresentationHandlers interface {
	CreateStudent() http.HandlerFunc
	CreateCourse() http.HandlerFunc
	AssignCourseToStudent() http.HandlerFunc
	GetStudent() http.HandlerFunc
	GetCourse() http.HandlerFunc
}

// PresentationHandlersImpl represents the usecase implementation object
type PresentationHandlersImpl struct {
	interactor *interactor.Interactor
}

// NewPresentationHandlers initializes a new REST handlers usecase
func NewPresentationHandlers(
	i *interactor.Interactor,
) PresentationHandlers {
	return &PresentationHandlersImpl{i}
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p PresentationHandlersImpl) CreateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.StudentCreationPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request body to struct: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		student := domain.Student{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
		}
		createdStudent, err := p.interactor.Courses.CreateStudent(ctx, &student)
		if err != nil {
			msg := fmt.Sprintf("error creating student: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, createdStudent, http.StatusCreated)
	}
}

func (p PresentationHandlersImpl) CreateCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.CourseCreationPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request body to struct: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		course := domain.Course{
			Title:       payload.Title,
			Price:       payload.Price,
			Description: payload.Description,
			Instructor:  payload.Instructor,
			Category:    payload.Category,
		}
		createdStudent, err := p.interactor.Courses.CreateCourse(ctx, &course)
		if err != nil {
			msg := fmt.Sprintf("error creating course: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, createdStudent, http.StatusCreated)
	}
}

func (p PresentationHandlersImpl) AssignCourseToStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.StudentCourseAssigningPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request body to struct: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}
		student, err := p.interactor.Courses.AssignCourseToStudent(ctx, &payload.Email, &payload.CourseTitle)
		if err != nil {
			msg := fmt.Sprintf("error assigning course to student: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, student, http.StatusCreated)
	}
}

func (p PresentationHandlersImpl) GetStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.GetStudentPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request boy to struct: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		email := payload.Email
		student, err := p.interactor.Courses.GetStudent(ctx, &email)
		if err != nil {
			msg := fmt.Sprintf("error getting student: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, student, http.StatusOK)
	}
}

func (p PresentationHandlersImpl) GetCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.GetCoursePayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request boy to struct: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		course, err := p.interactor.Courses.GetCourse(ctx, &payload.CourseTitle)
		if err != nil {
			msg := fmt.Sprintf("error getting course: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, course, http.StatusOK)
	}
}
