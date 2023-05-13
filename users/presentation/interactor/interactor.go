package interactor

import "github.com/MelvinKim/users/usecase"

type Interactor struct {
	Users usecase.UsecaseContract
}

func NewUsersInteractor(
	users usecase.UsecaseContract,
) (*Interactor, error) {
	return &Interactor{
		Users: users,
	}, nil
}
