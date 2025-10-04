package store


import "CRM-GO/internal/domain"


type Storer interface {
Create(contact *domain.Contact) error
List() ([]domain.Contact, error)
GetByID(id uint) (*domain.Contact, error)
Update(contact *domain.Contact) error
Delete(id uint) error
}