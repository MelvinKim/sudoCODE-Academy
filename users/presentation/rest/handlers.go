package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MelvinKim/users/application/common/dto"
	"github.com/MelvinKim/users/domain"
	"github.com/MelvinKim/users/presentation/interactor"
)

// PresentationHandlers represents all the REST API logic
type PresentationHandlers interface {
	CreateStudent() http.HandlerFunc
	GetStudent() http.HandlerFunc
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
			msg := fmt.Sprintf("error unmarshalling request boy to struct: %v", err)
			w.Header().Set("Content-Type", "application/json")
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		student := domain.Student{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
		}
		createdStudent, err := p.interactor.Users.CreateStudent(ctx, &student)
		if err != nil {
			msg := fmt.Sprintf("error creating student: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, createdStudent, http.StatusCreated)
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
		student, err := p.interactor.Users.GetStudent(ctx, &email)
		if err != nil {
			msg := fmt.Sprintf("error getting student: %v", err)
			jsonResponse(w, map[string]string{"error": msg}, http.StatusBadRequest)
			return
		}

		jsonResponse(w, student, http.StatusOK)
	}
}
