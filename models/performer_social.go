/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// PerformerSocial --
type PerformerSocial struct {
	Twitter string `json:"twitter,omitempty"`

	Youtube string `json:"youtube,omitempty"`

	Instagram string `json:"instagram,omitempty"`

	Snapchat string `json:"snapchat,omitempty"`
}
