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
func (d *Data) Del(keyname any) bool {
	newk, err := Normalize(keyname)
	if err != nil {
		return false
	}

	// check if key exists first.
	if _, v := d.data_storage[newk]; v {
		delete(d.data_storage, newk)
		return true
	} else {
		return false
	}
}

// func to get all values from data_storage
func (d *Data) GetAll() (map[string]string, bool) {
	res := make(map[string]string, len(d.data_storage)) // initialise the size as number of elements in data_storage to reduce size allocated for this map

	if len(d.data_storage) != 0 {
		for v, i := range d.data_storage {
			res[i] = v
			return res, true
		}
	}
	return nil, false
}
