package interactor

import "github.com/MelvinKim/courses/usecase"

type Interactor struct {
	Courses usecase.UsecaseContract
}

func NewUsersInteractor(
	courses usecase.UsecaseContract,
) (*Interactor, error) {
	return &Interactor{
		Courses: courses,
	}, nil
}
