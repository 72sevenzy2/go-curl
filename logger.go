package main

import (
	"fmt"
	"net/http"
	"time"
)

func Log(v *http.Client, req *http.Request) (time.Duration, *http.Response, error) {
	start := time.Now()
	resp, err := v.Do(req)
	if err != nil {
		return 0, nil, err
	}
	end := time.Since(start)

	fmt.Println("visited to:", req.URL.Path)
	fmt.Println("method:", req.Method)
	// exlude sensitive headers
	newHeaders := req.Header.Clone()
	newHeaders.Del("Authorization")
	fmt.Println("with headers:", newHeaders) // then display


	return end, resp, nil
}
