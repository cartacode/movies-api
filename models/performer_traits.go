/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// PerformerTraits --
type PerformerTraits struct {
	Ethnicity string `json:"ethnicity"`
	HairColor string `json:"haircolor"`
	Piercings bool   `json:"piercings"`
	Tattoos   bool   `json:"tattoos"`
	StarSign  string `json:"sign"`
}
