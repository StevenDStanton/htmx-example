package forms

import (
	"reflect"
	"testing"
)

func TestNewFormData(t *testing.T) {
	formData := NewFormData()

	if formData.Values == nil {
		t.Errorf("Expected non-nil map for Values, got nil")
	}
	if formData.Errors == nil {
		t.Errorf("Expected non-nil map for Errors, got nil")
	}
}

func TestNewPageData(t *testing.T) {
	data := "some data"
	formData := NewFormData()
	pageData := NewPageData(data, formData)

	if pageData.Data != data {
		t.Errorf("Expected Data to be %v, got %v", data, pageData.Data)
	}
	if !reflect.DeepEqual(pageData.Form, formData) {
		t.Errorf("Expected Form to be %v, got %v", formData, pageData.Form)
	}
}
