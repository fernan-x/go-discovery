package userinfra

import (
	"errors"

	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type InMemoryUserRepository struct {
	users []userdomain.User
}

var _ userdomain.UserRepository = (*InMemoryUserRepository)(nil)

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []userdomain.User{},
	}
}

func (r *InMemoryUserRepository) Create(u *userdomain.User) error {
	r.users = append(r.users, *u)
	return nil
}

func (r *InMemoryUserRepository) GetAll() ([]userdomain.User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepository) GetByEmail(email string) (*userdomain.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, errors.New("User with email " + email + " not found")
}