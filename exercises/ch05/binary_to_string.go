package ch05

import (
	"errors"
)

func ToString(i float64) (string, error) {
	if i == 1 || i == 0 {
		return "", nil
	}
	j := i * 2
	current := ""
	if j >= 1 {
		j = j - 1
		current = "1"
	} else {
		current = "0"
	}

	r, err := ToString(j)
	if err != nil {
		return "", err
	}
	if len(r) > 32 {
		return "", errors.New("overflow")
	}

	return current + r, nil
}
