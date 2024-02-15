package cache

import (
	"encoding/json"

	e_files "github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	e_user "github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/bradfitz/gomemcache/memcache"
)

type Mencacher[T Entity] interface {
	Set(key string, i T) error
	Get(key string) (*memcache.Item, error)
	GetMulti(key []string) (map[string]*memcache.Item, error)
	Delete(key string) error
}

type Memcached[T Entity] struct {
	Client *memcache.Client
}

func NewMencached[T Entity](server ...string) *Memcached[T] {
	client := memcache.New(server...)
	return &Memcached[T]{
		Client: client,
	}
}

type Entity interface {
	*e_user.User | *e_files.Contract | *e_files.Extract | *e_files.Invoice
}

func (m *Memcached[T]) Set(key string, i T) error {
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

func (m *Memcached[T]) Get(key string) (*memcache.Item, error) {
	item, err := m.Client.Get(key)

	if err != nil {
		return nil, errors.ErrCtx(err, "Error get cache")
	}
	return item, nil
}

func (m *Memcached[T]) GetMulti(key []string) (map[string]*memcache.Item, error) {

	items, err := m.Client.GetMulti(key)
	if err != nil {
		return nil, errors.ErrCtx(err, "Error get multi cache")
	}
	return items, nil
}

func (m *Memcached[T]) Delete(key string) error {
	err := m.Client.Delete(key)
	if err != nil {
		return errors.ErrCtx(err, "Error to set cache")
	}
	return nil
}
