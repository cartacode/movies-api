package requests

import (
	"encoding/json"
	"net/http"
)

// ReturnAPIOK --
// NotFound
// swagger:response apiOk
func ReturnAPIOK(w http.ResponseWriter, json []byte) error {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return nil
}

// JSONPaginationResponse --
// HTTP status code 200 and repository model in data
// swagger:response jsonPaginationResp
type JSONPaginationResponse struct {
	// in: body
	Results       interface{} `json:"results"`
	TotalResults  int         `json:"total"`
	RecordsOnPage int         `json:"recordsonpage"`
	Page          int         `json:"page"`
	TotalPages    int         `json:"totalpages"`
}

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
	Message    string      `json:"message"`
	Identifier string      `json:"identifier"`
	Extra      interface{} `json:"extra"`
}

// ReturnAPIError --
func ReturnAPIError(w http.ResponseWriter, err error) error {
	payload := JSONErrorResponse{Error: err.Error()}
	js, err := json.Marshal(payload)

	log.Error(payload)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(js)

	return err
}

// ReturnOnError --
func ReturnOnError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Error(err)
		ReturnAPIError(w, err)
		return true

	}
	return false
}
