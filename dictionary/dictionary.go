package dictionary

type Dictionary map[string]string

const (
	ErrNoKey        = DictErr("that key does not exist")
	ErrDuplicateKey = DictErr("cannot use existing key to overwrite value")
	ErrNoKeyUpdate  = DictErr("no key found to update for value")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrNoKeyUpdate
	}
	d[key] = value

	return nil
}
