/*
 * Vuli API
 *
 * Vuli Category Delivery API
 *
 * API version: 3

 */

package controllers

import (
	"net/http"
	"testing"
)

func TestCategoryCategoryIDDelete(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CategoryCategoryIDDelete(tt.args.w, tt.args.r)
		})
	}
}

func TestCategoryPost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CategoryPost(tt.args.w, tt.args.r)
		})
	}
}

func TestCategoryGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CategoryGet(tt.args.w, tt.args.r)
		})
	}
}

func TestCategoryCategoryIDPatch(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CategoryCategoryIDPatch(tt.args.w, tt.args.r)
		})
	}
}

func TestCategorySlugGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CategorySlugGet(tt.args.w, tt.args.r)
		})
	}
}
