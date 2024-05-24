package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(wr http.ResponseWriter, response any) error {
	wr.Header().Set("Content-Type", "application/json")

	if res, ok := response.(*error_response); ok {
		wr.WriteHeader(res.StatusCode)
	}

	if res, ok := response.(*success_response); ok {
		wr.WriteHeader(res.Status)
	}
	return json.NewEncoder(wr).Encode(response)
}
