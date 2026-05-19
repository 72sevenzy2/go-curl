package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func intPtr(i int) *int { return &i } // returns *int

func Log(v *http.Client, req *http.Request, bodyAllowed *bool, bodySize *int) (time.Duration, *http.Response, string, *int, error) {
	start := time.Now()
	resp, err := v.Do(req)
	if err != nil {
		return 0, nil, "", intPtr(0), err
	}
	end := time.Since(start)

	// fmt.Println("visited to:", req.URL.Path)
	// fmt.Println("method:", req.Method)
	fmt.Printf("visited to %s, with method %s", req.URL.Path, req.Method)
	fmt.Println(" full url:", req.URL)
	// request query
	if req.URL.RawQuery != "" {
		fmt.Println("with query:", req.URL.RawQuery)
	}

	// exlude sensitive headers
	fmt.Println("request header details:")
	newHeaders := req.Header.Clone()
	newHeaders.Del("Authorization")
	fmt.Println(newHeaders) // then display

	// request body printing
	var bodyprev string // body preview (string) var
	var max *int // body size var
	var bodybytes []byte // body size (in bytes)

	if *bodyAllowed {
		max = bodySize
		bodybytes, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		resp.Body = io.NopCloser(bytes.NewBuffer(bodybytes))
		if *max > 0 && len(bodybytes) > *max {
			bodybytes = bodybytes[:*max]
		}
		bodyprev = string(bodybytes)

	}

	return end, resp, bodyprev, max, nil
}
