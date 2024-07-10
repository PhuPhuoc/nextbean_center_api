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

func WriteFileExcel(w http.ResponseWriter, fileContent []byte) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    w.Header().Set("Content-Disposition", "attachment; filename=Book.xlsx")
    w.WriteHeader(http.StatusOK)

    if _, err := w.Write(fileContent); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
}