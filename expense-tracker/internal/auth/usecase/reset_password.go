package authusecase

import (
	"fmt"
	"time"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
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
}

func NewResetPasswordUsecase(
	authService authdomain.AuthService,
	userRepo userdomain.UserRepository,
	tokenRepo authdomain.PasswordResetTokenRepository,
) *ResetPasswordUsecase {
	return &ResetPasswordUsecase{authService, userRepo, tokenRepo}
}

func (u *ResetPasswordUsecase) Execute(resetToken string, newPassword string) error {
	// 1️⃣ Input:
	// resetToken (string)
	// newPassword (string)

	t, err := u.tokenRepo.GetByToken(resetToken)
	if err != nil {
		return fmt.Errorf("invalid reset token")
	}

	if t.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("reset token expired")
	}

	// 4️⃣ Hash the new password (using bcrypt or AuthService.HashPassword)
	// 5️⃣ Update the user’s password in the database
	// 6️⃣ Delete or invalidate the reset token (for security!)
	// userId := t.UserId


	return fmt.Errorf("not implemented")
}