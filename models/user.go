package models

type User struct {
	Name      string `json:"name" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	TaxId     string `json:"taxId" validate:"required"`
	BirthDate string `json:"birthDate" validate:"required"`
}

func NewUser(name, lastName, taxId, birthDate string) User {
	return User{Name: name, LastName: lastName, TaxId: taxId, BirthDate: birthDate}
}
