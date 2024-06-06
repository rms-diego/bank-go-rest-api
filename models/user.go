package models

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	TaxId     string `json:"taxId" validate:"required"`
	BirthDate string `json:"birthDate" validate:"required"`
}
