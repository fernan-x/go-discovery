package user_usecase

import (
	password_hasher "github.com/fernan-x/expense-tracker/internal/shared/password-hasher"
	user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type CreateUserUseCaseInterface interface {
	Execute(email string, password string, firstName string, lastName string) error
}

var _ CreateUserUseCaseInterface = &CreateUserUseCase{}

type CreateUserUseCase struct {
	repo user_domain.UserRepository
	passwordHasher password_hasher.PasswordHasher
}

func NewCreateUserUseCase(repo user_domain.UserRepository, passwordHasher password_hasher.PasswordHasher) *CreateUserUseCase {
	return &CreateUserUseCase{repo, passwordHasher}
}

func (u *CreateUserUseCase) Execute(email string, password string, firstName string, lastName string) error {
	hash, err := u.passwordHasher.Hash(password)
	if err != nil {
		return err
	}

	err = u.repo.Create(&user_domain.User{
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