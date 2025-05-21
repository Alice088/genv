package genv

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Init(path string) {
	if len(path) == 0 {
		panic("path is empty")
	}

	err := godotenv.Load(path)
	if err != nil {
		panic(fmt.Sprintf("cannot load .env file: %s", err.Error()))
	}
}
