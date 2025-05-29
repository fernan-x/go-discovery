package userinfra

import (
	"fmt"

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

	return nil, fmt.Errorf("user with email %s not found", email)
}

func updateFields(u *userdomain.User, fields userdomain.UserUpdateFields) {
	if fields.FirstName != nil {
		u.FirstName = *fields.FirstName
	}
	if fields.LastName != nil {
		u.LastName = *fields.LastName
	}
	if fields.Password != nil {
		u.Password = *fields.Password
	}
}

func (r *InMemoryUserRepository) Update(userId string, fields userdomain.UserUpdateFields) error {
	for i, u := range r.users {
		if u.ID == userId {
			updateFields(&u, fields)
			r.users[i] = u
			return nil
		}
	}

	return fmt.Errorf("user with id %s not found", userId)
}