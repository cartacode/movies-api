package api

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"time"
)

type res struct {
	Response interface{} `json:"response,omitempty"`
	Message  string      `json:"message,omitempty"`
}

// Respond -- standard implementation for API responses to
// front-end application
// **** @Richard - if we can use gzipped responses for all API requests, then we could use requests.ReturnAPIOK
func Respond(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	var res = res{}
	cacheSince := time.Now().Format(http.TimeFormat)
	w.Header().Set("Last-Modified", cacheSince)
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Encoding", "gzip")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error()
	} else {
		w.WriteHeader(http.StatusOK)
	}
	gz := gzip.NewWriter(w)
	res.Response = data
	json.NewEncoder(gz).Encode(res)
	gz.Close()
}
