package json


import (
	"encoding/json"
	"fmt"
	"os"
	"sort"


	"CRM-GO/internal/domain"
	"CRM-GO/internal/store"
)

var _ store.Storer = (*JSONStore)(nil)

type JSONStore struct {
path string
seq uint
}


type snapshot struct {
Seq uint `json:"seq"`
Contacts []domain.Contact `json:"contacts"`
}


func New(path string) (*JSONStore, error) {
	s := &JSONStore{path: path}
	if err := s.bootstrap(); err != nil { return nil, err }
	return s, nil
}


func (s *JSONStore) bootstrap() error {
	b, err := os.ReadFile(s.path)
	if err != nil { return err }
	var snap snapshot
	if err := json.Unmarshal(b, &snap); err != nil { return err }
	s.seq = snap.Seq
	return nil
}

func (s *JSONStore) load() (snapshot, error) {
	b, err := os.ReadFile(s.path)
	if err != nil { return snapshot{}, err }
	var snap snapshot
	if err := json.Unmarshal(b, &snap); err != nil { return snapshot{}, err }
	return snap, nil
}


func (s *JSONStore) persist(snap snapshot) error {
	b, err := json.MarshalIndent(snap, "", " ")
	if err != nil { return err }
	return os.WriteFile(s.path, b, 0o644)
}


func (s *JSONStore) Create(c *domain.Contact) error {
	snap, err := s.load(); if err != nil { return err }
	s.seq = snap.Seq + 1
	c.ID = s.seq
	snap.Seq = s.seq
	snap.Contacts = append(snap.Contacts, *c)
	return s.persist(snap)
}


func (s *JSONStore) List() ([]domain.Contact, error) {
	snap, err := s.load(); if err != nil { return nil, err }
	sort.Slice(snap.Contacts, func(i, j int) bool { return snap.Contacts[i].ID < snap.Contacts[j].ID })
	return snap.Contacts, nil
}


func (s *JSONStore) GetByID(id uint) (*domain.Contact, error) {
	snap, err := s.load(); if err != nil { return nil, err }
	for _, c := range snap.Contacts { if c.ID == id { cc := c; return &cc, nil } }
	return nil, fmt.Errorf("not found")
}


func (s *JSONStore) Update(c *domain.Contact) error {
	snap, err := s.load(); if err != nil { return err }
	for i := range snap.Contacts {
		if snap.Contacts[i].ID == c.ID { snap.Contacts[i] = *c; return s.persist(snap) }
	}
	return fmt.Errorf("not found")
}


func (s *JSONStore) Delete(id uint) error {
	snap, err := s.load(); if err != nil { return err }
	for i := range snap.Contacts {
		if snap.Contacts[i].ID == id {
			snap.Contacts = append(snap.Contacts[:i], snap.Contacts[i+1:]...)
			return s.persist(snap)
		}
	}
	return fmt.Errorf("not found")
}