package genv

import (
	"errors"
	"fmt"
	"strconv"
)

type mutantType interface {
	int | bool | []byte
}

// MuGet - Convert ENV value(string) to T
func MuGet[T mutantType](key string) (T, error) {
	var mutant T
	var rawValue string
	var err error

	if rawValue, err = Get(key); err != nil {
		return mutant, err
	}

	switch any(mutant).(type) {
	case int:
		i, err := strconv.Atoi(rawValue)
		if err != nil {
			return mutant, err
		}
		return any(i).(T), nil
	case bool:
		parseBool, err := strconv.ParseBool(rawValue)
		if err != nil {
			return mutant, err
		}

		return any(parseBool).(T), nil
	case []byte:
		return any(rawValue).(T), nil
	default:
		return mutant, errors.New(fmt.Sprintf("not support type: %T", any(mutant)))
	}
}
