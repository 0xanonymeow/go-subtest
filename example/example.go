package example

import "errors"

func Print(s string) (string, error) {
	if s == "error" {
		return "", errors.New("error expected")
	}

	return s, nil
}
