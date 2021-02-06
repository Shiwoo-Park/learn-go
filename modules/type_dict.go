package modules

import (
	"errors"
	"log"
)

// Dictionary type
type Dictionary map[string]string

// multiple variable declaration
var (
	errorNotFound  = errors.New("Not found")
	errorKeyExists = errors.New("Key exists")
)

// Search for key and return value
func (d Dictionary) Search(key string) (string, error) {
	//
	val, exists := d[key]
	if exists {
		return val, nil
	}
	return "", errorNotFound
}

// Add new data in dict
func (d Dictionary) Add(key string, val string) error {
	_, err := d.Search(key)

	switch err {
	case errorNotFound:
		d[key] = val
	case nil:
		return errorKeyExists
	}

	return nil
}

// Update data in dict
func (d Dictionary) Update(key string, val string) error {
	_, err := d.Search(key)

	switch err {
	case errorNotFound:
		return err
	case nil:
		d[key] = val
	}

	return nil
}

// Delete KV pair in dict
func (d Dictionary) Delete(key string) {
	_, err := d.Search(key)
	if err == errorNotFound {
		log.Println("Warn: key doesn't exists")
	}
	delete(d, key)
}
