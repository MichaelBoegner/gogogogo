package dictionary

import "errors"

type Dictionary map[string]string

var ErrNoKey = errors.New("that key does not exist")
var ErrDuplicateKey = errors.New("cannot use existing key to overwrite value")

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]
	if !found {
		return "", ErrNoKey
	}
	return value, nil
}

func (d Dictionary) Submit(key, value string) error {
	if d[key] != "" {
		return ErrDuplicateKey
	}
	d[key] = value
	return nil
}
