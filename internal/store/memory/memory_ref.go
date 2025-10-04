package memory
import (
	"errors"
	"sync"


	"CRM-GO/internal/domain"
	"CRM-GO/internal/store"
)


var _ store.Storer = (*MemoryStore)(nil)


type MemoryStore struct {
	mu sync.RWMutex
	seq uint
	contacts map[uint]domain.Contact
}


func New() *MemoryStore {
	return &MemoryStore{contacts: make(map[uint]domain.Contact)}
}


func (s *MemoryStore) Create(c *domain.Contact) error {
	s.mu.Lock(); defer s.mu.Unlock()
	s.seq++
	c.ID = s.seq
	s.contacts[c.ID] = *c
	return nil
}


func (s *MemoryStore) List() ([]domain.Contact, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	out := make([]domain.Contact, 0, len(s.contacts))
	for _, c := range s.contacts { out = append(out, c) }
	return out, nil
}


func (s *MemoryStore) GetByID(id uint) (*domain.Contact, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	c, ok := s.contacts[id]
	if !ok { return nil, errors.New("not found") }
	return &c, nil
}


func (s *MemoryStore) Update(c *domain.Contact) error {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.contacts[c.ID]; !ok { return errors.New("not found") }
	s.contacts[c.ID] = *c
	return nil
}


func (s *MemoryStore) Delete(id uint) error {
	s.mu.Lock(); defer s.mu.Unlock()
	if _, ok := s.contacts[id]; !ok { return errors.New("not found") }
	delete(s.contacts, id)
	return nil
}