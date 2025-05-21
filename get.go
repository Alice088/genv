package genv

import (
	"errors"
	"fmt"
	"os"
)

func Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key must not be empty")
	}

	v := os.Getenv(key)
	if len(v) == 0 {
		return "", fmt.Errorf("%s is not set", key)
	}

	return v, nil
}
