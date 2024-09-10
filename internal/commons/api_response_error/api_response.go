package mist_api_error

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type ErrorDetail struct {
	Detail string `json:"detail,omitempty"`
}

func ProcessApiErrorInventory(ctx context.Context, status_code int, body io.ReadCloser) interface{} {
	var error_interface interface{}
	if status_code == 400 {
		bodyBytes, e := io.ReadAll(body)
		if e == nil {
			json.Unmarshal(bodyBytes, error_interface)
		}
	}
	return error_interface
}

func ProcessApiError(ctx context.Context, status_code int, body io.ReadCloser, err error) string {
	var error_response string = ""
	if status_code == 400 {
		bodyBytes, e := io.ReadAll(body)
		if e == nil {
			error_body := ErrorDetail{}
			json.Unmarshal(bodyBytes, &error_body)
			error_response = error_body.Detail
		}
	}
	if err != nil && error_response == "" {
		error_response = fmt.Sprintf("unexpected error: %s", err.Error())
	}
	return error_response
}
