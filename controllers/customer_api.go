/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package controllers

import (
	"net/http"

	"github.com/VuliTv/api/models"
)

// JSONPaginationResponseCustomer --
// HTTP status code 200 and repository model in data
// swagger:response customerResp
type JSONPaginationResponseCustomer struct {
	// in: body
	Results       []models.Customer `json:"results"`
	TotalResults  int               `json:"total"`
	RecordsOnPage int               `json:"recordsonpage"`
	Page          int               `json:"page"`
	TotalPages    int               `json:"totalpages"`
}

// CustomerCustomerIDDelete --
func CustomerCustomerIDDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerCustomerIDGet --
func CustomerCustomerIDGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerCustomerIDPatch --
func CustomerCustomerIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerGet --
func CustomerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerPost --
func CustomerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
