package errors

import "net/http"

type ErrorDetails struct {
  Code string `json:"code"`
  Message string `json:"message"`
}

type HttpError struct {
  Status int `json:"status"`
  Error ErrorDetails `json:"error"`
}

func NewHttpError (status int, message string, code string) (int, HttpError) {
  return status, HttpError{
    Status: status,
    Error: ErrorDetails{ Code: code, Message: message },
  }
}

func NewNotFoundError(details ErrorDetails) (int, HttpError) {
  message := details.Message
  code := details.Code

  if details.Message == "" { message = "the requested resource was not found" }
  if details.Code == "" { code = "not_found" }

  return http.StatusNotFound, HttpError{
    Status: http.StatusNotFound,
    Error: ErrorDetails{
      Message: message,
      Code: code,
    },
  }
}
