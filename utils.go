package main

import (
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
	"strings"
	"strconv"
)

func (h *HeaderFlags) String() string { // gets called internally by go's flag pkg, (type flag.Value expects a String() and Set() func)
	return fmt.Sprint(*h)
}

func (h *HeaderFlags) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func Validate(args []string) error {
	if len(args) < 1 {
		UsageMsg := errors.New("usage > main.go <URL> [-H key:value]")
		return fmt.Errorf("%s", UsageMsg.Error())
	}
	return nil
}

// func for adding headers
func AddHeaders(req *http.Request, args HeaderFlags) error {
	for _, h := range args {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid input type %s", h)
		}

		// appending headers
		req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}
	return nil
}

// normalize key types to string (for storage.go)
func Normalize(keyname any) (string, error) {
	switch v := keyname.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	default:
		errms := errors.New("invalid type: consider only string or int.")
		return "", errms
	}
}

// small utils to convert to json (for session.go)
func ToJson(v any) (string, error) {
	jsons, err := json.Marshal(v)
	if err == nil {
		return string(jsons), nil
	} else {
		return "", fmt.Errorf("error: %s", err.Error()) // print error
	}
}