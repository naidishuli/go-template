package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type JWT struct {
	secret string
}

func New(secret string) *JWT {
	return &JWT{secret}
}

func (j JWT) ParseToken(tokenValue string, claim jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenValue, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrVerification
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrTokenInvalid
	}

	return nil
}

func (j JWT) GenerateToken(claim jwt.Claims) (string, error) {
	tokenValue := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return tokenValue.SignedString([]byte(j.secret))
}
