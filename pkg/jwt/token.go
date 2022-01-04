package jwt

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// VerifyDataToken validate the authentication header and fills the data variable with the information found
func VerifyDataToken(header, secret string, data interface{}) error {
	bearerToken := strings.Split(header, " ")
	if len(bearerToken) != 2 {
		return fmt.Errorf("bearer token malformed or not present")
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error verification jwtoken method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("token valid")
	}

	return json.Unmarshal([]byte(token.Claims.(jwt.MapClaims)["user"].(string)), data)
}
