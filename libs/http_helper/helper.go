package httphelper

import (
	"encoding/json"
	"net/http"
)

// JSONErrorResponse --
// NotFound
// swagger:response genericJsonError
type JSONErrorResponse struct {
	Error string `json:"error"`
}

// JSONSuccessResponse --
// Success response
// swagger:response ok
type JSONSuccessResponse struct {
	Message    string `json:"message"`
	Identifier string `json:"identifier"`
}

// ReturnAPIOK --
func ReturnAPIOK(w http.ResponseWriter, json []byte) error {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return nil
}

// ReturnAPIError --
func ReturnAPIError(w http.ResponseWriter, err error) {
	payload := JSONErrorResponse{Error: err.Error()}
	js, err := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(js)
}
