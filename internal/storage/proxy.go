package storage

import (
	"golang.org/x/net/context"
	"ktbs.dev/mubeng/pkg/model"
)

type ProxyStorer interface {
	AddProxy(ctx context.Context, proxy *model.Proxy) error
	DeleteProxy(ctx context.Context, proxyAddr string) error
	GetProxy(ctx context.Context, proxyAddr string) (*model.Proxy, error)
	GetAllProxies(ctx context.Context) ([]*model.Proxy, error)
}
