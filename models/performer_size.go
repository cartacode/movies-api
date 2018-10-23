/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// PerformerSize --
type PerformerSize struct {
	Weight int32 `json:"weight,omitempty"`

	Waist int32 `json:"waist,omitempty"`

	Bust string `json:"bust,omitempty"`

	Height int32 `json:"height,omitempty"`
}
