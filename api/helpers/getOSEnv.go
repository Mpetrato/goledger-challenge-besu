package helpers

import (
	"errors"
	"os"
)

func GetOSEnv(key string) (*string, error) {
	value := os.Getenv(key)
	if value == "" {
		return nil, errors.New("error on get env")
	}

	return &value, nil
}
