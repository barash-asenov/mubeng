package redis_test

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	redisStorage "ktbs.dev/mubeng/internal/storage/redis"
	"ktbs.dev/mubeng/pkg/model"
)

func TestAddProxy(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	proxy := &model.Proxy{
		Address: "192.168.1.1:5050",
	}
	proxyStore := redisStorage.NewProxyStorage("root", rdb)

	// expectations
	mock.ExpectHSet("root", "192.168.1.1:5050", []byte(`{"address": "192.168.1.1:5050"}`))

	if err := proxyStore.AddProxy(context.Background(), proxy); err != nil {
		t.Error("expected successfully adding proxy: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
