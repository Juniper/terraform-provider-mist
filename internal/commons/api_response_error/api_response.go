package mist_api_error

import (
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

func ProcessApiErrorInventory(statusCode int, body io.ReadCloser) (errorInterface interface{}) {
	if statusCode == 400 {
		bodyBytes, e := io.ReadAll(body)
		if e == nil {
			err := json.Unmarshal(bodyBytes, errorInterface)
			if err != nil {
				return nil
			}
		}
	}
	return errorInterface
}

func ProcessApiError(statusCode int, body io.ReadCloser, err error) (errorResponse string) {
	if statusCode >= 300 {
		bodyBytes, e := io.ReadAll(body)
		if e == nil {
			errorBody := ErrorDetail{}
			err := json.Unmarshal(bodyBytes, &errorBody)
			if err != nil {
				return fmt.Sprintf("unexpected error: %s", err.Error())
			}
			errorResponse = errorBody.Detail
		}
	}
	if err != nil && errorResponse == "" {
		errorResponse = fmt.Sprintf("unexpected error: %s", err.Error())
	}
	return errorResponse
}

func ProcessInventoryApiError(
	action string,
	statusCode int,
	body io.ReadCloser,
	err error,
) (errorResponse []string, vcMemberAssignWarning bool) {
	vcMemberAssignWarning = false
	if statusCode == 400 || statusCode == 404 {
		var valueType = "information"
		switch action {
		case "assign":
			valueType = "MAC Address"
		case "unassign":
			valueType = "MAC Address"
		case "unclaim":
			valueType = "Serial"
		case "claim":
			valueType = "Claim Code"
		}
		bodyBytes, e := io.ReadAll(body)
		if e == nil {
			errorBody := ErrorInventory{}
			err = json.Unmarshal(bodyBytes, &errorBody)
			if err != nil {
				return nil, vcMemberAssignWarning
			}
			for i, mac := range errorBody.Error {
				reason := errorBody.Reason[i]
				if reason == "VC member switches could not be assigned" {
					vcMemberAssignWarning = true
				} else {
					errorResponse = append(errorResponse, fmt.Sprintf("Invalid %s \"%s\". Reason from Mist: %s", valueType, mac, reason))
				}
			}
		}
	}
	if err != nil && len(errorResponse) == 0 && !vcMemberAssignWarning {
		errorResponse = append(errorResponse, fmt.Sprintf("unexpected error: %s", err.Error()))
	}
	return errorResponse, vcMemberAssignWarning
}
