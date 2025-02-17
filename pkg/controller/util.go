package controller

import (
	"encoding/json"
	"net/http"
)

func HTTPResponse(w http.ResponseWriter, message any, statusCode int) map[string]any {
	resp := map[string]any{
		"StatusCode": statusCode,
		"Message": message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	return resp
}