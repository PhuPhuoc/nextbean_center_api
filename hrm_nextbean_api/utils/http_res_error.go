package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type error_response struct {
	StatusCode int         `json:"status-code"`
	Message    interface{} `json:"message"`
	Log        string      `json:"log"`
}

func validate_nil_err(e error) string {
	var log string
	if e == nil {
		log = ""
	} else {
		log = e.Error()
	}
	return log
}

func ErrorResponse_NewFull(status int, mess string, e error) *error_response {
	return &error_response{
		StatusCode: status,
		Message:    mess,
		Log:        validate_nil_err(e),
	}
}

func ErrorResponse_NoPermission(mess string) *error_response {
	return &error_response{
		StatusCode: http.StatusForbidden, // 403
		Message:    "You have no permission",
		Log:        mess,
	}
}

func ErrorResponse_Unauthorized() *error_response {
	return &error_response{
		StatusCode: http.StatusUnauthorized, // 401
		Message:    "You have no permission",
		Log:        "",
	}
}

func ErrorResponse_TokenExpired() *error_response {
	return &error_response{
		StatusCode: http.StatusUnauthorized, // 401
		Message:    "You have no permission",
		Log:        "Token expired, please login again",
	}
}

func ErrorResponse_BadRequest(mess string, e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    mess,
		Log:        validate_nil_err(e),
	}
}

func ErrorResponse_BadRequest_ListError(mess interface{}, e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    mess,
		Log:        validate_nil_err(e),
	}
}

func ErrorResponse_InvalidRequest(e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid request",
		Log:        validate_nil_err(e),
	}
}

func ErrorResponse_Server(e error) *error_response {
	return &error_response{
		StatusCode: http.StatusInternalServerError,
		Message:    "something went wrong with server",
		Log:        e.Error(),
	}
}

func ErrorResponse_DB(e error) *error_response {
	return &error_response{
		StatusCode: http.StatusInternalServerError,
		Message:    "something went wrong with DB",
		Log:        e.Error(),
	}
}

func ErrorResponse_CannotListEntity(entity string, e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		Log:        e.Error(),
	}
}

func ErrorResponse_CannotGetEntity(entity string, e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		Log:        e.Error(),
	}
}

func ErrorResponse_CannotUpdateEntity(entity string, e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		Log:        e.Error(),
	}
}

func ErrorResponse_CannotDeleteEntity(entity string, e error) *error_response {
	return &error_response{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		Log:        e.Error(),
	}
}
