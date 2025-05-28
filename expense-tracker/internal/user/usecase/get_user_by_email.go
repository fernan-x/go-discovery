package userusecase

import userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"

type GetUserByEmailUseCaseInterface interface {
	Execute(email string) (*userdomain.User, error)
}

var _ GetUserByEmailUseCaseInterface = &GetUserByEmail{}

type GetUserByEmail struct {
	repo userdomain.UserRepository
}

func NewGetUserByEmailUseCase(repo userdomain.UserRepository) *GetUserByEmail {
	return &GetUserByEmail{repo}
}

func (u *GetUserByEmail) Execute(email string) (*userdomain.User, error) {
	return u.repo.GetByEmail(email)
}