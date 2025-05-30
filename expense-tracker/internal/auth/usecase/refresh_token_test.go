package authusecase_test

import (
	"testing"
	"time"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	authinfra "github.com/fernan-x/expense-tracker/internal/auth/infra"
	authusecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	passwordhasher "github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	tokenissuer "github.com/fernan-x/expense-tracker/internal/shared/tokenissuer"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupTest() authdomain.AuthService {
	var passwordHasher = &passwordhasher.BcryptPasswordHasher{}
	var tokenIssuer = tokenissuer.NewJwtTokenIssuer([]byte("secret"))
	return authinfra.NewAuthService(passwordHasher, tokenIssuer)
}

func generateExpiredRefreshToken(secret string, userID string) string {
	claims := jwt.MapClaims{
		"sub":  userID,
		"iat":  time.Now().Add(-48 * time.Hour).Unix(), // issued 2 days ago
		"exp":  time.Now().Add(-24 * time.Hour).Unix(), // expired 1 day ago
		"iss":  "expense-tracker",
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestRefreshToken_Success(t *testing.T) {
	authService := setupTest()
	uc := authusecase.NewRefreshTokenUseCase(authService)
	refreshToken, _ := authService.GenerateRefreshToken("12345")

	res, err := uc.Execute(refreshToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.AccessToken)
	assert.NotEmpty(t, res.RefreshToken)

	claims, err := authService.ParseToken(res.RefreshToken)
	assert.NoError(t, err)
	assert.Equal(t, "12345", claims["sub"])
	assert.Equal(t, "refresh", claims["type"])

	claims, err = authService.ParseToken(res.AccessToken)
	assert.NoError(t, err)
	assert.Equal(t, "access", claims["type"])
}

func TestRefreshToken_Failure_InvalidToken(t *testing.T) {
	authService := setupTest()
	uc := authusecase.NewRefreshTokenUseCase(authService)

	_, err := uc.Execute("invalid-token")
	assert.Error(t, err)
}

func TestRefreshToken_Failure_ExpiredToken(t *testing.T) {
	authService := setupTest()
	expiredToken := generateExpiredRefreshToken("secret", "12345")
	uc := authusecase.NewRefreshTokenUseCase(authService)

	_, err := uc.Execute(expiredToken)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token is expired")
}