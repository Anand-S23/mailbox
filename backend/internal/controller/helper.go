package controller

import (
	"encoding/json"
	"net/http"
)

func ErrorMessage(message string) map[string]string {
    return map[string]string {"error": message}
}

func InternalServerError(w http.ResponseWriter) error {
    errMsg := ErrorMessage("Internal server error, please try again")
    return WriteJSON(w, http.StatusInternalServerError, errMsg)
}

func BadRequestError(w http.ResponseWriter, message string) error {
    errMsg := ErrorMessage(message)
    return WriteJSON(w, http.StatusBadRequest, errMsg)
}

func UnauthorizedError(w http.ResponseWriter) error {
    errMsg := ErrorMessage("Unauthorized")
    return WriteJSON(w, http.StatusUnauthorized, errMsg)
}

func PageNotFoundError(w http.ResponseWriter) error {
    errMsg := ErrorMessage("Page not found")
    return WriteJSON(w, http.StatusNotFound, errMsg)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

