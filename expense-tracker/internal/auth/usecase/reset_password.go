package authusecase

import (
	"fmt"
	"time"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	"github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type ResetPasswordUsecaseInterface interface {
	Execute(resetToken string, newPassword string) error
}

var _ ResetPasswordUsecaseInterface = &ResetPasswordUsecase{}

type ResetPasswordUsecase struct {
	authService authdomain.AuthService
	userRepo    userdomain.UserRepository
	tokenRepo   authdomain.PasswordResetTokenRepository
	pwdHasher   passwordhasher.PasswordHasher
}

func NewResetPasswordUsecase(
	authService authdomain.AuthService,
	userRepo userdomain.UserRepository,
	tokenRepo authdomain.PasswordResetTokenRepository,
	pwdHasher passwordhasher.PasswordHasher,
) *ResetPasswordUsecase {
	return &ResetPasswordUsecase{authService, userRepo, tokenRepo, pwdHasher}
}

func (u *ResetPasswordUsecase) Execute(resetToken string, newPassword string) error {
	t, err := u.tokenRepo.GetByToken(resetToken)
	if err != nil {
		return fmt.Errorf("invalid reset token")
	}

	if t.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("reset token expired")
	}

	hash, err := u.pwdHasher.Hash(newPassword)
	if err != nil {
		return err
	}

	err = u.userRepo.Update(t.UserId, userdomain.UserUpdateFields{
		Password: &hash,
	})
	if err != nil {
		return err
	}

	u.tokenRepo.Delete(resetToken)

	return nil
}