package gorm

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"


	"CRM-GO/internal/domain"
	"CRM-GO/internal/store"
)

var _ store.Storer = (*GORMStore)(nil)


type GORMStore struct { db *gorm.DB }


func New(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil { return nil, err }
	if err := db.AutoMigrate(&domain.Contact{}); err != nil { return nil, err }
	return &GORMStore{db: db}, nil
}


func (s *GORMStore) Create(c *domain.Contact) error { return s.db.Create(c).Error }


func (s *GORMStore) List() ([]domain.Contact, error) {
	var out []domain.Contact
	return out, s.db.Order("id asc").Find(&out).Error
}

func (s *GORMStore) GetByID(id uint) (*domain.Contact, error) {
	var c domain.Contact
	if err := s.db.First(&c, id).Error; err != nil {
	if errors.Is(err, gorm.ErrRecordNotFound) { return nil, errors.New("not found") }
	return nil, err
	}
	return &c, nil
}

func (s *GORMStore) Update(c *domain.Contact) error { return s.db.Save(c).Error }
func (s *GORMStore) Delete(id uint) error { return s.db.Delete(&domain.Contact{}, id).Error }