package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/m3-app/backend/types"
)

// Given a JSON file, map the contents into any struct dest
func FileMapper(filename string, dest interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%s not found", filename)
	}
	if err = json.Unmarshal(file, dest); err != nil {
		return err
	}
	return nil
}

func ParseTokens() (types.Tokens, error) {
	var tokens types.Tokens
	err := FileMapper("tokens.json", &tokens)
	return tokens, err
}

// Given a bearer token ("Bearer <TOKEN>") returns the token or an error if parsing was unsuccessful
func GetBearerToken(authorization string) (string, error) {
	parsed_authorization := strings.Split(authorization, " ")
	if parsed_authorization[0] != "Bearer" || len(parsed_authorization) < 2 {
		return "", fmt.Errorf("could not parse bearer token")
	}
	token := parsed_authorization[1]
	return token, nil
}
