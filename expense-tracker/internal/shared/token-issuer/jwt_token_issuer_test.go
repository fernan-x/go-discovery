package token_issuer_test

import (
	"testing"

	token_issuer "github.com/fernan-x/expense-tracker/internal/shared/token-issuer"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var secret = []byte("secret")
var jti = token_issuer.NewJwtTokenIssuer(secret)

func TestJwtTokenIssuer_GenerateAccessToken_Success(t *testing.T) {
	token, err := jti.GenerateAccessToken("12345")
	if err != nil {
		t.Fatal(err)
	}

	jwtDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	claims, ok := jwtDecoded.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("Could not parse claims")
	}

	assert.Equal(t, "12345", claims["sub"])
	assert.Equal(t, "expense-tracker", claims["iss"])
	assert.Equal(t, "access", claims["type"])
	assert.NotEmpty(t, claims["iat"])
	assert.NotEmpty(t, claims["exp"])
}

func TestJwtTokenIssuer_GenerateRefreshToken_Success(t *testing.T) {
	token, err := jti.GenerateRefreshToken("12345")
	if err != nil {
		t.Fatal(err)
	}

	jwtDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	claims, ok := jwtDecoded.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("Could not parse claims")
	}

	assert.Equal(t, "12345", claims["sub"])
	assert.Equal(t, "expense-tracker", claims["iss"])
	assert.Equal(t, "refresh", claims["type"])
	assert.NotEmpty(t, claims["iat"])
	assert.NotEmpty(t, claims["exp"])
}

func TestJwtTokenIssuer_Parse_Success(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDg1MzE0MzEsImlhdCI6MTc0ODQ0NTAzMSwiaXNzIjoiRXhwZW5zZSBUcmFja2VyIiwic3ViIjoiMTIzNDUifQ.lqRuN8G2PQKiGhEdQOXd92EN0imzrdmDrncfF8CDxvE"
	claims, err := jti.Parse(token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, "12345", claims["sub"])
}

func TestJwtTokenIssuer_Parse_Failure_Signature(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDg1MzE0MzEsImlhdCI6MTc0ODQ0NTAzMSwiaXNzIjoiRXhwZW5zZSBUcmFja2VyIiwic3ViIjoiMTIzNDUifQ.lqRuN8G2PQKiGhEdQOXd92EN0imzrdmDrncfF8CDxvE"
	jti := token_issuer.NewJwtTokenIssuer([]byte("wrong-secret"))
	_, err := jti.Parse(token)
	assert.Error(t, err)
}