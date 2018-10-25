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
	Weight int32 `json:"weight"`

	Waist int32 `json:"waist"`

	Bust string `json:"bust"`

	Height int32 `json:"height"`
}
