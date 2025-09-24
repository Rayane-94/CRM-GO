package storage

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetByID(id int) (*Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

//err perso

var ErrContactNotFound = func(id int) error { return fmt.Errorf("Contact not found %d", id)}