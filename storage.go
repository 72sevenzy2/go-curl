// to hold user inputs during session mode - storing variables in maps to then be retrieved later using dynamic lookups.

package main

import (
	"errors"
)

type Data struct {
	data_storage map[string]string 
}

// new db
func NewStore() *Data {
	return &Data{
		make(map[string]string),
	}
}

// utility get/set functions for data map:

// for both strings and ints
func (d *Data) Get(keyname any) (string, bool, error) {
	newKey, err := Normalize(keyname)
	if err != nil {
		return "", false, err
	}

	val, ok := d.data_storage[newKey]
	return val, ok, nil
}

func (d *Data) Set(keyname any, value string) (error, bool) {
	if value == "" {
		errM := errors.New("please include a value")
		return errM, false
	}

	newK, err := Normalize(keyname)
	if err != nil {
		return err, false
	}
	d.data_storage[newK] = value
	return nil, true
}

// del func
func (d *Data) Del(keyname any) (error, bool) {
	newk, err := Normalize(keyname)
	if err == nil {
		delete(d.data_storage, newk)
		return nil, true
	} else {
		return errors.New("key does not exist."), false
	}
}
