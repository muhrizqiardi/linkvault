package utils

import "errors"

func ValidateEnumString(input string, enums ...string) error {
	for _, enum := range enums {
		if input != enum {
			continue
		}
		return nil
	}

	return errors.New("Input string invalid")
}
