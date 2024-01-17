package cache

import (
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/bradfitz/gomemcache/memcache"
)

type Mencacher interface {
	Get(key string) (item *memcache.Item, err error)
	Set(item *memcache.Item) error
	Delete(key string) error
	GetMulti(key []string) (map[string]*memcache.Item, error)
}

type Mencached struct {
	Client *memcache.Client
}

func NewMencached(client *memcache.Client) *Mencached {
	return &Mencached{
		Client: client,
	}
}

func (m *Mencached) Get(key string) (*memcache.Item, error) {
	item, err := m.Client.Get(key)
	if err != nil {
		return nil, errors.NewError(err, "Error get cache")
	}
	return item, nil
}

func (m *Mencached) Set(item *memcache.Item) error {
	err := m.Client.Set(item)
	if err != nil {
		return errors.NewError(err, "Error to set cache")
	}
	return nil
}

func (m *Mencached) Delete(key string) error {
	err := m.Client.Delete(key)
	if err != nil {
		return errors.NewError(err, "Error to set cache")
	}
	return nil
}

func (m *Mencached) GetMulti(key []string) (map[string]*memcache.Item, error) {
	items, err := m.Client.GetMulti(key)
	if err != nil {
		return nil, errors.NewError(err, "Error get multi cache")
	}
	return items, nil
}
