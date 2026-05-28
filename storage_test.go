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

	_, ok2 := store.Set("test", "hi");
	if !ok2 {
		fmt.Println("couldnt set data.")
	}
	
	vals, ok3 := store.GetAll()
	if !ok3 {
		fmt.Println("not vars set.")
	}
	for i, v := range vals {
		fmt.Println(v, i)
	}
}
