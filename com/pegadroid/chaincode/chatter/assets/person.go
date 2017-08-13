/**
*********************************************************
Copyright (c) 2017 Pegadroid IQ Solutions Private Limited
All rights reserved
*********************************************************
*/

package assets

// PersonAssetInterface is the interface for our Person asset
type PersonAssetInterface interface {
	SetID(id int)
	SetName(name string)
	SetEmail(email string)
	GetID() int
	GetName() string
	GetEmail() string
}

// Person stuct is the struct implmenting the PersonAssetInterface interface
type Person struct {
	ID    int
	Name  string
	Email string
}

func (person *Person) SetID(id int) {
	person.ID = id
}

func (person *Person) GetID() int {
	return person.ID
}

func (person *Person) SetEmail(email string) {
	person.Email = email
}

func (person *Person) GetEmail() string {
	return person.Email
}

func (person *Person) SetName(name string) {
	person.Name = name
}

func (person *Person) GetName() string {
	return person.Name
}
