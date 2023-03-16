package utils

import (
	"strings"

	"github.com/stadio-app/go-backend/types"
)

func CreateFullName(name types.FullName) string {
	full_name := []string{}
	full_name = append(full_name, name.FirstName)
	if name.MiddleName != nil {
		full_name = append(full_name, *name.MiddleName)
	}
	full_name = append(full_name, name.LastName)
	return strings.Join(full_name, " ")
}
