package main

import (
	"bufio"
	"fmt"
	"io"
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
			if len(parts) < 2 { // validate arguments before continuing (other it will panic)
				fmt.Println("variable as second argument does not exit, consider setting a var.")
				continue
			} else {
				val, ok, err := store.Get(parts[1]) // check if parts[1] exists as a var first
				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				// flag value to confirm if header values are validated
				pass := true

				if ok {
					client := http.Client{}
					cl, err := http.NewRequest(http.MethodGet, val, nil) // new request
					if err != nil {
						fmt.Println(err.Error())
						continue
					}

					// parse header arguments
					for i := range len(parts) {
						upc := strings.ToUpper(parts[i]) // normalize "-h" to all uppercase

						if upc == "-H" {
							// check if header values exist
							if i+1 >= len(parts) {
								fmt.Println("missing header values.")
								continue
							}

							headers := strings.SplitN(parts[i+1], ":", 2) // splits headers as:
							// for example: header1:value, splitN() would split it so:
							// map[string]string{
							// 	"header1",
							//  "value",
							// }

							if len(headers) < 2 { // validate length of headers or it will panic during execution
								fmt.Println("please include both header name and value.")
								pass = false
								break
							}
							if headers[0] != "" && headers[1] != "" {
								cl.Header.Add(headers[0], headers[1])
							} else {
								continue
							}
						} else {
							continue
						}
					}

					if !pass { continue }
					resp, err2 := client.Do(cl) // send the request to the url provided
					if err2 == nil {
						// shadow auth header from resp
						clonedH := resp.Header.Clone()
						clonedH.Del("authorization")

						respB := resp.Body
						defer respB.Close() // close the connection at the end of this block

						// outputting
						fmt.Println("response headers:")
						// fmt.Println(clonedH)
						for header, val := range clonedH {
							fmt.Println("header:", header)
							for _, v := range val {
								fmt.Println("value:", v)
								for range 5 {
									fmt.Print("-")
								}
							}
						}
						for range 10 {
							fmt.Print("-") // seperator for headers and resp body so its easier to read
						}
						body, err := io.ReadAll(respB) // read respB bytes
						if err != nil {
							continue // skip current iteration if no body
						} else {
							fmt.Println("\nresponse body:")
							for range 10 {
								fmt.Print("-")
							}
							fmt.Println("\n", string(body))
						}
					}

				} else { // if key does not exit
					fmt.Println("variable does not exit, consider setting one.")
					continue
				}
			}

		// for exiting
		case "EXIT":
			fmt.Println("exiting will remove saved variables.")
			return
		}
	}
}
