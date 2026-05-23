package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// small utils to convert to json
func toJson(v any) (string, error) {
	jsons, err := json.Marshal(v)
	if err == nil {
		return string(jsons), nil
	} else {
		return "", fmt.Errorf("error: %s", err.Error()) // print error
	}
}

// start a interactive session
func StartSession(b *bufio.Scanner, store *Data) {
	fmt.Println("session started.")
	for {
		fmt.Print(">")
		isOk := b.Scan() // take input
		if !isOk {
			fmt.Println(b.Err()) // possible scanner error case
		}

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
			err, ok := store.Del(parts[1])
			if !ok && err != nil {
				fmt.Println(err.Error())
				continue
			} else {
				fmt.Println("deleted key.")
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

				//  utility variables
				var (
					pass     bool
					reqType  string
					bodyData map[string]string
					jsonData string
				)
				pass = true

				// request initialisation variables
				var (
					cl    *http.Request
					clErr error
				)

				if ok {
					client := http.Client{}
					// cl, err := http.NewRequest(reqType, val, nil) // new request
					if jsonData != "" && reqType == http.MethodPost {
						cl, clErr = http.NewRequest(reqType, val, bytes.NewBuffer([]byte(jsonData)))
					}

					if clErr != nil {
						fmt.Println(err)
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

							if len(headers) < 2 || len(headers) > 2 { // validate length of headers or it will panic during execution
								fmt.Println("please include both header name and value.")
								pass = false
								break
							}
							if headers[0] != "" && headers[1] != "" {
								cl.Header.Add(headers[0], headers[1])
							} else {
								continue
							}
						}

						// to check if argument is for req methods
						if upc == "-X" {

							// validate if values exist (GET or POST)
							if i+1 >= len(parts) {
								fmt.Println("missing argument values, consider either POST or GET.")
								continue
							}

							if parts[i+1] == "POST" {
								reqType = http.MethodPost

								// validate if request body flag also exists aswell as its data
								if i+2 >= len(parts) {
									fmt.Println("missing request body flag, consider: -d [data]")
									continue
								}

								uppercased1 := strings.ToUpper(parts[i+2]) // normalize parts[i+2]

								if uppercased1 == "-D" { // for normal json data

									// validate if values exist
									if i+3 >= len(parts) {
										fmt.Println("please include actual data in json format.")
										continue
									}
									bodyData = map[string]string{
										"data": parts[i+3],
									}

									data, err := toJson(bodyData)
									if err != nil {
										fmt.Println(err.Error())
										continue
									} else {
										jsonData = data
									}

								}

								if uppercased1 == "-F" { // form mode
									if i+3 >= len(parts) {
										fmt.Println("please include the necessary form data with format: 'title:value'")
										continue
									}

									formParts := strings.SplitN(parts[i+3], ":", 2) // needs to be in a format like so:
									// "title:value", and strings.SplitN(...) would return:
									// map[string]string{
									//		"title": "value",
									// }

									// validate if formParts is of correct length now (or will panic)
									if len(formParts) < 2 || len(formParts) > 2 {
										fmt.Println("please use the correct format.")
										continue
									}

									bodyData = map[string]string{ // appending necessary deaails to body
										"title": formParts[0],
										"value": formParts[1],
									}
								}

							} else {
								if parts[i+1] == "GET" {
									reqType = http.MethodGet
								} else {
									fmt.Println("invalid method type.")
									continue
								}
							}
						}

					}

					if !pass {
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
