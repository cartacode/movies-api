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
	Credit *CustomerCredit `json:"credit"`

	// Purchased Items
	Purchased *CustomerPurchased `json:"purchased"`

	// True/False. Is the user active
	Active bool `json:"active"`

	// User wishlist
	Wishlist *CustomerWishlist `json:"wishlist"`
}
