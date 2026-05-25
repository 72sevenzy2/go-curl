package main

import (
	"fmt"
	"testing"
)

func TestStorage(t *testing.T) {
	store := NewStore()
	_, ok := store.Set(1, "test")
	if ok {
		fmt.Println("data set successfully")
	} else {
		fmt.Println("could not save data")
	}

	k := store.Del(1)
	if !k {
		fmt.Println("key does no exist.")
	}

	val, exists, _ := store.Get(1)
	if exists {
		fmt.Println("data exists:", val)
	} else {
		fmt.Println("couldnt retrieve data") // expected output for del testing
	}
}
