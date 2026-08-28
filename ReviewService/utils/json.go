package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter,statusCode int,data any) error {
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonSuccessResponse(w http.ResponseWriter,status int, message string ,data any ) error {
	w.Header().Set("content-type","application/json")
	w.WriteHeader(status)
	response := map[string]any{}
	response["status"] = "success"
	response["message"] = message
	response["data"] = data

	return WriteJsonResponse(w,status,response)

}


func WriteJsonErrorResponse(w http.ResponseWriter, status int, errorMessage string,err error) error {
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(status)
	response := map[string]any{}
	response["message"] = errorMessage
	response["statusCode"] = "error"
	response["error"] = err.Error()

	return WriteJsonResponse(w,status,response)
}


func ReadJsonBody(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)

}