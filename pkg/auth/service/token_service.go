package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthJWT struct {
	secretKey string
}

func NewAuthJWTService(secretKey string) *AuthJWT {
	return &AuthJWT{
		secretKey: secretKey,
	}
}

func (s *AuthJWT) GenerateTokenJWT() (string, error) {
	claims := jwt.MapClaims{"timestamp": time.Now().UTC().Format((time.RFC3339)), }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthJWT) VerifyToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return false, err
	}

	if parsedToken.Valid {
		return true, nil
	}

	return false, err
}