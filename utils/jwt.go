package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/m3-app/backend/ent"
)

func GenerateJWT(key string, user *ent.User) (string, error) {
	today := time.Now()
	month := (time.Hour * 24) * 30 // 24 hours * 30 days
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID.String(),
		"name": user.Name,
		"email": user.Email,
		"avatar": user.Avatar,
		"creation": today.Unix(),
		"expiration": today.Add(month).Unix(),
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetJwtClaims(jwt_token string, key string) (jwt.MapClaims, error) {
	token, token_err := jwt.Parse(jwt_token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(key), nil
	})
	if token_err != nil {
		return nil, fmt.Errorf("could not parse jwt token")
	}
	
	// Get claims stored in parsed JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not fetch jwt claims")
	}
	return claims, nil
}
