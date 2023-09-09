package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if d[key] == "" {
		return "", errors.New("that key does not exist")
	}
	return d[key], nil
}
