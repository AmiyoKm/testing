package dictionary

import "errors"

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var ErrNotFound = errors.New("could not fund the word you were looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")
var ErrWordDoesNotExist = errors.New("cannot perform operation on word because it does not exit")

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return res, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	if err != nil {
		switch err {
		case ErrNotFound:
			return ErrWordDoesNotExist
		default:
			return err
		}
	}
	d[word] = newDefinition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err != nil {
		switch err {
		case ErrNotFound:
			return ErrWordDoesNotExist
		default:
			return err
		}
	}
	delete(d, word)
	return nil
}
