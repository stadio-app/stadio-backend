package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Given a JSON file, map the contents into any struct dest
func FileMapper[T any](filename string, dest T) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%s not found", filename)
	}
	if err = json.Unmarshal(file, dest); err != nil {
		return err
	}
	return nil
}
