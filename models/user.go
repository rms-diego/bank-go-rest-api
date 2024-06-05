package models

type User struct {
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	TaxId     string `json:"taxId"`
	BirthDate string `json:"birthDate"`
}

func NewUser(name, lastName, taxId, birthDate string) User {
	return User{Name: name, LastName: lastName, TaxId: taxId, BirthDate: birthDate}
}
