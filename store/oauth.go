package store

import (
	"github.com/RangelReale/osin"
)

// OAuthStore represents an in-memory store for OAuth clients and tokens.
type OAuthStore struct {
	clients   map[string]osin.Client
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

// NewOAuthStore creates a OAuthStore instance.
func NewOAuthStore() *OAuthStore {
	r := &OAuthStore{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

	// TODO: Pass in client ID and secret
	r.clients["admin"] = &osin.DefaultClient{
		Id:          "admin",
		Secret:      "Digitox123",
		RedirectUri: "http://localhost/",
	}

	return r
}

// Clone clones store.
func (s *OAuthStore) Clone() osin.Storage {
	return s
}

// Close closes store. Since this store is in-memory, there's no need to close any resources.
func (s *OAuthStore) Close() {}

// GetClient retrieves a client by ID from the store.
func (s *OAuthStore) GetClient(id string) (osin.Client, error) {
	if c, ok := s.clients[id]; ok {
		return c, nil
	}

	return nil, osin.ErrNotFound
}

// SetClient sets a client by ID.
func (s *OAuthStore) SetClient(id string, client osin.Client) error {
	s.clients[id] = client
	return nil
}

// SaveAuthorize saves authorization data to the store.
func (s *OAuthStore) SaveAuthorize(data *osin.AuthorizeData) error {
	s.authorize[data.Code] = data
	return nil
}

// LoadAuthorize loads authorization data from the store.
func (s *OAuthStore) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}

	return nil, osin.ErrNotFound
}

// RemoveAuthorize removes authorization code from store.
func (s *OAuthStore) RemoveAuthorize(code string) error {
	delete(s.authorize, code)
	return nil
}

// SaveAccess saves access token data to the store.
func (s *OAuthStore) SaveAccess(data *osin.AccessData) error {
	s.access[data.AccessToken] = data
	if data.RefreshToken != "" {
		s.refresh[data.RefreshToken] = data.AccessToken
	}

	return nil
}

// LoadAccess loads access token data from the store.
func (s *OAuthStore) LoadAccess(code string) (*osin.AccessData, error) {
	if d, ok := s.access[code]; ok {
		return d, nil
	}

	return nil, osin.ErrNotFound
}

// RemoveAccess removes access token data from the store.
func (s *OAuthStore) RemoveAccess(code string) error {
	delete(s.access, code)
	return nil
}

// LoadRefresh loads refresh token data from the store.
func (s *OAuthStore) LoadRefresh(code string) (*osin.AccessData, error) {
	if d, ok := s.refresh[code]; ok {
		return s.LoadAccess(d)
	}

	return nil, osin.ErrNotFound
}

// RemoveRefresh removes refresh token data from the store.
func (s *OAuthStore) RemoveRefresh(code string) error {
	delete(s.refresh, code)
	return nil
}
