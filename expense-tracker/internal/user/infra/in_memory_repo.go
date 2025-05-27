package user_infra

import (
	"errors"

	user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

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

func (r *InMemoryUserRepository) GetByEmail(email string) (*user_domain.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, errors.New("User with email " + email + " not found")
}