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

type ErrorInventory struct {
	Op      string   `json:"op,omitempty"`
	Success []string `json:"success"`
	Error   []string `json:"error"`
	Reason  []string `json:"reason"`
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

func ProcessInventoryApiError(ctx context.Context, action string, status_code int, body io.ReadCloser, err error) []string {
	var error_response []string
	if status_code == 400 || status_code == 404 {
		var value_type string = "information"
		switch action {
		case "assign":
			value_type = "MAC Address"
		case "unassign":
			value_type = "MAC Address"
		case "unclaim":
			value_type = "Serial"
		case "claim":
			value_type = "Claim Code"
		}
		bodyBytes, e := io.ReadAll(body)
		if e == nil {
			error_body := ErrorInventory{}
			json.Unmarshal(bodyBytes, &error_body)

			for i, mac := range error_body.Error {
				reason := error_body.Reason[i]
				error_response = append(error_response, fmt.Sprintf("Invalid %s \"%s\". Reason from Mist: %s", value_type, mac, reason))
			}
		}
	}
	if err != nil && len(error_response) == 0 {
		error_response = append(error_response, fmt.Sprintf("unexpected error: %s", err.Error()))
	}
	return error_response
}
