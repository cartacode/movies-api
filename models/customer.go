/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// Customer Document
//
// A customer on the site and all of their preferences and profile
//
// swagger:model
type Customer struct {

	// Unique email for this customer, read only, cognito controlled
	// read only: true
	Email string `json:"email"`

	// Credit Information
	Credit *CustomerCredit `json:"credit,omitempty"`

	// Purchased Items
	Purchased *CustomerPurchased `json:"purchased,omitempty"`

	// True/False. Is the user active
	Active bool `json:"active,omitempty"`

	// User wishlist
	Wishlist *CustomerWishlist `json:"wishlist,omitempty"`
}
