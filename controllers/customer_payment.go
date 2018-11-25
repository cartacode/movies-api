/*
 * Vuli API
 *
 * Vuli Customer Payment Methods API
 *
 * API version: 3
 */

package controllers

import "net/http"

// CustomerPaymentAdd --
func CustomerPaymentAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerPaymentDelete --
func CustomerPaymentDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerPaymentUpdate --
func CustomerPaymentUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
