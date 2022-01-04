package jwt

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// VerifyDataToken validate the authentication header and fills the data variable with the information found
func (j *JWT) VerifyDataToken(header string, data interface{}) error {
	bearerToken := strings.Split(header, " ")
	if len(bearerToken) != 2 {
		return fmt.Errorf("bearer token malformed or not present")
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error verification jwtoken method")
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("token valid")
	}

	return json.Unmarshal([]byte(token.Claims.(jwt.MapClaims)["user"].(string)), data)
}

// GenerateUserToken generate a jwt token string using an HS256 signing method from the given secret
func (j *JWT) GenerateUserToken(userData string) (token Token, err error) {
	expiresAt := time.Now().Add(time.Hour * 30 * 24).Unix()
	claimMap := jwt.MapClaims{
		"user": userData,
		"exp":  expiresAt,
	}

	tokenValue := jwt.NewWithClaims(jwt.SigningMethodHS256, claimMap)

	tokenString, err := tokenValue.SignedString([]byte(j.secret))
	if err != nil {
		return token, fmt.Errorf("cannot sign token string: %s", err.Error())
	}

	return Token{Token: tokenString, ExpiresAt: expiresAt}, nil
}
