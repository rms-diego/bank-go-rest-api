package models

type user struct {
	Name      string
	LastName  string
	TaxId     string
	BirthDate string
}

func NewUser(name, lastName, taxId, birthDate string) *user {
	return &user{Name: name, LastName: lastName, TaxId: taxId, BirthDate: birthDate}
}
