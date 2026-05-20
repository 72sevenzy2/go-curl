package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

// start a interactive session
func StartSession(b *bufio.Scanner, store *Data) {
	fmt.Println("session started.")
	for {
		fmt.Print(">")
		b.Scan() // take input

		input := b.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue // skip current iteration if no input
		}

		upperInput := strings.ToUpper(parts[0]) // "VAR", "GET", "DEL", "EXIT"

		switch upperInput {
		// declare variables (can be to store headers, urls, etc)
		case "VAR":
			err, ok := store.Set(parts[1], parts[2])
			if err != nil && !ok {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("successful.")
			continue
		case "GET":
			val, ok, err := store.Get(parts[1])
			if err != nil && !ok {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(val)

		case "DEL":
			fmt.Println("deleted key.")
			err, ok := store.Del(parts[1])
			if !ok && err != nil {
				fmt.Println(err.Error())
				continue
			}

		// actual api testing logic (GET only for now)
		case "TEST":
			if parts[1] == "" { // parts[1] will be the url the user wanted to test
				fmt.Println("please include a valid url aswell.")
				continue
			} else {
				client := http.Client{}
				cl, err := http.NewRequest(http.MethodGet, parts[1], nil) // new request
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
 
				resp, err2 := client.Do(cl) // send the request to the url provided
				if err2 == nil {
					// shadow auth header from resp
					clonedH := resp.Header.Clone()
					clonedH.Del("authorization")

					respB := resp.Body
					defer respB.Close() // close the connection at the end of this block

					// outputting
					fmt.Println("response headers:")
					fmt.Println(clonedH)
					for range 10 {
						fmt.Println("-") // seperator for headers and resp body so its easier to read
					}
					fmt.Println("response body:")
					fmt.Println(respB)
				}
			}
		// for exiting
		case "EXIT":
			fmt.Println("exiting will remove saved variables.")
			return
		}
	}
}
