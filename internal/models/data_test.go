package models

import (
	"testing"
)

func TestNewData(t *testing.T) {
	data := NewData()

	if len(data.Contacts) != 2 {
		t.Errorf("expected 2 contacts, got %d", len(data.Contacts))
	}

	expectedContacts := []Contact{
		{Name: "Sam Spade", Email: "sam@spade.com"},
		{Name: "John Doe", Email: "John.doe@gmail.com"},
	}

	for i, contact := range expectedContacts {
		if data.Contacts[i] != contact {
			t.Errorf("expected contact %v, got %v", contact, data.Contacts[i])
		}
	}
}

func TestEmailExists(t *testing.T) {
	data := NewData()

	tests := []struct {
		email    string
		expected bool
	}{
		{"sam@spade.com", true},
		{"John.doe@gmail.com", true},
		{"nonexistent@example.com", false},
	}

	for _, test := range tests {
		result := data.EmailExists(test.email)
		if result != test.expected {
			t.Errorf("for email %s, expected %v, got %v", test.email, test.expected, result)
		}
	}
}
