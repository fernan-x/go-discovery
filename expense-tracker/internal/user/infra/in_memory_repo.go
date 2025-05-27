package user_infra

import user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"

type InMemoryUserRepository struct {
	users []user_domain.User
}

var _ user_domain.UserRepository = (*InMemoryUserRepository)(nil)

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []user_domain.User{},
	}
}

func (r *InMemoryUserRepository) Create(u *user_domain.User) error {
	r.users = append(r.users, *u)
	return nil
}

func (r *InMemoryUserRepository) GetAll() ([]user_domain.User, error) {
	return r.users, nil
}