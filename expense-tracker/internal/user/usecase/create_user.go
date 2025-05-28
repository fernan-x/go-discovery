package userusecase

import (
	passwordhasher "github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type CreateUserUseCaseInterface interface {
	Execute(email string, password string, firstName string, lastName string) error
}

var _ CreateUserUseCaseInterface = &CreateUserUseCase{}

type CreateUserUseCase struct {
	repo userdomain.UserRepository
	passwordHasher passwordhasher.PasswordHasher
}

func NewCreateUserUseCase(repo userdomain.UserRepository, passwordHasher passwordhasher.PasswordHasher) *CreateUserUseCase {
	return &CreateUserUseCase{repo, passwordHasher}
}

func (u *CreateUserUseCase) Execute(email string, password string, firstName string, lastName string) error {
	hash, err := u.passwordHasher.Hash(password)
	if err != nil {
		return err
	}

	err = u.repo.Create(&userdomain.User{
		Email: email,
		Password: string(hash),
		FirstName: firstName,
		LastName: lastName,
	})
	if err != nil {
		return err
	}

	return nil
}