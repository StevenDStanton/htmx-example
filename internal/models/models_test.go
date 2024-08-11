package models

import (
	"testing"
)

func TestNewContact(t *testing.T) {
	name := "John Doe"
	email := "john.doe@example.com"

	contact := NewContact(name, email)

	if contact.Name != name {
		t.Errorf("expected name %v, got %v", name, contact.Name)
	}

	if contact.Email != email {
		t.Errorf("expected email %v, got %v", email, contact.Email)
	}
}
