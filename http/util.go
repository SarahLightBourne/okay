package http

import (
	"errors"
	"regexp"
	"strings"
)

func validateKey(inputKey string) (string, error) {
	key := strings.Trim(inputKey, " ")
	regex := regexp.MustCompile("^[a-zA-Z0-9_]+$")

	if regex.MatchString(key) {
		return key, nil
	}

	return key, errors.New("invalid key")
}
