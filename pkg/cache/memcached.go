package cache

import (
	"encoding/json"

	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/bradfitz/gomemcache/memcache"
)

type Mencacher interface {
	Set(key string, i interface{}) error
	Get(key string) (*memcache.Item, error)
	GetMulti(key []string) (map[string]*memcache.Item, error)
	Delete(key string) error
}

type Memcached struct {
	Client *memcache.Client
}

func NewMencached(server ...string) *Memcached {
	client := memcache.New(server...)
	return &Memcached{
		Client: client,
	}
}

func (m *Memcached) Set(key string, i interface{}) error {
	value, err := json.Marshal(i)
	if err != nil {
		return errors.ErrCtx(err, "Error marshaling data to JSON")
	}
	item := &memcache.Item{
		Key:   key,
		Value: value,
	}
	err = m.Client.Set(item)
	if err != nil {
		return errors.ErrCtx(err, "Error setting cache")
	}
	return nil
}

func (m *Memcached) Get(key string) (*memcache.Item, error) {
	item, err := m.Client.Get(key)

	if err != nil {
		return nil, errors.ErrCtx(err, "Error get cache")
	}
	return item, nil
}

func (m *Memcached) GetMulti(key []string) (map[string]*memcache.Item, error) {

	items, err := m.Client.GetMulti(key)
	if err != nil {
		return nil, errors.ErrCtx(err, "Error get multi cache")
	}
	return items, nil
}

func (m *Memcached) Delete(key string) error {
	err := m.Client.Delete(key)
	if err != nil {
		return errors.ErrCtx(err, "Error to set cache")
	}
	return nil
}
