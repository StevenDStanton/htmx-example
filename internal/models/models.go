package models

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) Contact {
	return Contact{name, email}
}
