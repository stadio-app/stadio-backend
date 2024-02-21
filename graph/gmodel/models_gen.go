// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Auth struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type CreateAccountInput struct {
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Name        string  `json:"name"`
	Password    string  `json:"password"`
}

type Mutation struct {
}

type Query struct {
}

type User struct {
	ID           int64             `json:"id" sql:"primary_key"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Email        string            `json:"email"`
	PhoneNumber  *string           `json:"phoneNumber,omitempty"`
	Name         string            `json:"name"`
	Avatar       *string           `json:"avatar,omitempty"`
	BirthDate    *time.Time        `json:"birthDate,omitempty"`
	Bio          *string           `json:"bio,omitempty"`
	Active       *bool             `json:"active,omitempty"`
	AuthPlatform *AuthPlatformType `json:"authPlatform,omitempty"`
}

type AuthPlatformType string

const (
	AuthPlatformTypeInternal AuthPlatformType = "INTERNAL"
	AuthPlatformTypeApple    AuthPlatformType = "APPLE"
	AuthPlatformTypeGoogle   AuthPlatformType = "GOOGLE"
)

var AllAuthPlatformType = []AuthPlatformType{
	AuthPlatformTypeInternal,
	AuthPlatformTypeApple,
	AuthPlatformTypeGoogle,
}

func (e AuthPlatformType) IsValid() bool {
	switch e {
	case AuthPlatformTypeInternal, AuthPlatformTypeApple, AuthPlatformTypeGoogle:
		return true
	}
	return false
}

func (e AuthPlatformType) String() string {
	return string(e)
}

func (e *AuthPlatformType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthPlatformType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthPlatformType", str)
	}
	return nil
}

func (e AuthPlatformType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
