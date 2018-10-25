/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// CustomerCredit --
type CustomerCredit struct {

	// Has the user stored credit information
	InfoStored bool `json:"info_stored"`

	// Key for 3 leg transactions to provider bank
	Key string `json:"key"`
}
