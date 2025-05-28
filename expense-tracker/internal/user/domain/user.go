package userdomain

type User struct {
	ID string
	Email string
	Password string
	FirstName string
	LastName string
}

type UserRepository interface {
	Create(u *User) error
	GetByEmail(email string) (*User, error)
}