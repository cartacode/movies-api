/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// Customer --
type Customer struct {

	// Unique email for this customer, read only, cognito controlled
	Email string `json:"email"`

	Credit *CustomerCredit `json:"credit,omitempty"`

	Purchased *CustomerPurchased `json:"purchased,omitempty"`

	// True/False. Is the user active
	Active bool `json:"active,omitempty"`

	Wishlist *CustomerWishlist `json:"wishlist,omitempty"`
}
