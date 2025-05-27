package user_usecase

import user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"

type GetUserByEmailUseCaseInterface interface {
	Execute(email string) (*user_domain.User, error)
}

var _ GetUserByEmailUseCaseInterface = &GetUserByEmail{}

type GetUserByEmail struct {
	repo user_domain.UserRepository
}

func NewGetUserByEmailUseCase(repo user_domain.UserRepository) *GetUserByEmail {
	return &GetUserByEmail{repo}
}

func (u *GetUserByEmail) Execute(email string) (*user_domain.User, error) {
	return u.repo.GetByEmail(email)
}