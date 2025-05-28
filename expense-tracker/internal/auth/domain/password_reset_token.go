package auth_domain

import "time"

type PasswordResetToken struct {
	Token string
	UserId string
	ExpiresAt time.Time
}

type PasswordResetTokenRepository interface {
	Save(token PasswordResetToken) error
	GetByToken(token string) (*PasswordResetToken, error)
}