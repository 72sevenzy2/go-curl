package main

import (
	"net/http"
	"time"
)

func Log(v *http.Client, req *http.Request) (time.Duration, *http.Response, error) {
	start := time.Now()
	resp, err := v.Do(req)
	end := time.Since(start)

	if err != nil {
		return 0, nil, err
	}

	return end, resp, nil
}
