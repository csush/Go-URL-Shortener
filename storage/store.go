package storage

import "errors"

const ErrNotFoundMsg = "URL doesn't exist"

type Store struct {
	urls map[string]string
}

func NewStore() *Store {
	return &Store{urls: make(map[string]string)}
}

func (s *Store) Save(shortURL, longURL string) {
	s.urls[shortURL] = longURL
}

func (s *Store) Get(shortURL string) (string, error) {
	longURL, exists := s.urls[shortURL]

	if exists {
		return longURL, nil
	} else {
		return "", errors.New(ErrNotFoundMsg)
	}
}
