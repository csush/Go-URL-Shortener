package storage

import (
	"errors"
	"sync"
)

const ErrNotFoundMsg = "url doesn't exist"

type IStore interface {
	Save(string, string)
	Get(string) (string, error)
}

type Store struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{urls: make(map[string]string)}
}

func (s *Store) Save(shortURL, longURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = longURL
}

func (s *Store) Get(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, exists := s.urls[shortURL]

	if exists {
		return longURL, nil
	} else {
		return "", errors.New(ErrNotFoundMsg)
	}
}
