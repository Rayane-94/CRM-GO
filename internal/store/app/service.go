package app
import (
	"errors"
	"strings"


	"CRM-GO/internal/domain"
	"CRM-GO/internal/store"
)


type ContactService struct { s store.Storer }


func NewContactService(s store.Storer) *ContactService { return &ContactService{s: s} }


func (svc *ContactService) Add(c domain.Contact) (*domain.Contact, error) {
	if strings.TrimSpace(c.Email) == "" && strings.TrimSpace(c.Phone) == "" {
		return nil, errors.New("email ou téléphone requis")
	}
	if err := svc.s.Create(&c); err != nil { return nil, err }
		return &c, nil
	}


func (svc *ContactService) List() ([]domain.Contact, error) { return svc.s.List() }


func (svc *ContactService) Update(c domain.Contact) error {
	if c.ID == 0 { return errors.New("id requis") }
		return svc.s.Update(&c)
	}


func (svc *ContactService) Delete(id uint) error { return svc.s.Delete(id) }