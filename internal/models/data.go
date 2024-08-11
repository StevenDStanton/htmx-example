package models

type Data struct {
	Contacts []Contact
}

func NewData() *Data {
	return &Data{
		Contacts: []Contact{
			{
				Name:  "Sam Spade",
				Email: "sam@spade.com",
			},
			{
				Name:  "John Doe",
				Email: "John.doe@gmail.com",
			},
		},
	}
}

func (d *Data) EmailExists(email string) bool {
	for _, user := range d.Contacts {
		if user.Email == email {
			return true
		}
	}
	return false
}
