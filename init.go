package genv

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
)

func Init(path string) error {
	if len(path) == 0 {
		return errors.New("path is empty")
	}

	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("cannot load .env file: %w", err)
	}

	return nil
}
