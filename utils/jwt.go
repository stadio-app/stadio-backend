package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stadio-app/go-backend/ent"
)

func GenerateJWT(key string, user *ent.User) (string, error) {
	today := time.Now()
	month := (time.Hour * 24) * 30 // 24 hours * 30 days
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"avatar": user.Avatar,
		"iat": today.Unix(),
		"exp": today.Add(month).Unix(),
	})
	return jwt.SignedString([]byte(key))
}

func GetJwtClaims(jwt_token string, key string) (jwt.MapClaims, error) {
	token, token_err := jwt.Parse(jwt_token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(key), nil
	})
	if token_err != nil {
		return nil, token_err
	}
	
	// Get claims stored in parsed JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not fetch jwt claims")
	}
	return claims, nil
}
