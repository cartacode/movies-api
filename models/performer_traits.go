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
	Ethnicity string `json:"ethnicity,omitempty"`
	HairColor string `json:"haircolor,omitempty"`
	Piercings bool   `json:"piercings,omitempty"`
	Tattoos   bool   `json:"tattoos,omitempty"`
	StarSign  bool   `json:"sign,omitempty"`
}
