package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"ktbs.dev/mubeng/pkg/model"
)

type ProxyStorage struct {
	rootKey string
	rdb     *redis.Client
}

func NewProxyStorage(rootKey string, rdb *redis.Client) *ProxyStorage {
	return &ProxyStorage{
		rootKey: rootKey,
		rdb:     rdb,
	}
}

func (s *ProxyStorage) AddProxy(ctx context.Context, proxy *model.Proxy) error {
	rawProxy, err := json.Marshal(proxy)
	if err != nil {
		return err
	}

	return s.rdb.HSet(ctx, s.rootKey, proxy.Address, rawProxy).Err()
}

func (s *ProxyStorage) DeleteProxy(ctx context.Context, proxyAddress string) error {
	return s.rdb.HDel(ctx, s.rootKey, proxyAddress).Err()
}

func (s *ProxyStorage) GetProxy(ctx context.Context, proxyAddress string) (*model.Proxy, error) {
	rawProxy, err := s.rdb.HGet(ctx, s.rootKey, proxyAddress).Result()
	if err != nil {
		return nil, err
	}

	var proxy = &model.Proxy{}
	err = json.Unmarshal([]byte(rawProxy), &proxy)
	if err != nil {
		return nil, err
	}

	return proxy, nil
}

func (s *ProxyStorage) GetAllProxies(ctx context.Context) ([]*model.Proxy, error) {
	rawProxies, err := s.rdb.HGetAll(ctx, s.rootKey).Result()
	if err != nil {
		return nil, err
	}

	var proxies = make([]*model.Proxy, 0, len(rawProxies))
	for _, rawProxy := range rawProxies {
		var proxy = &model.Proxy{}
		err := json.Unmarshal([]byte(rawProxy), &proxy)
		if err != nil {
			return nil, err
		}

		proxies = append(proxies, proxy)
	}

	return proxies, nil
}
