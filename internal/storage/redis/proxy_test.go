package redis_test

import (
	"context"
	"encoding/json"
	"sort"
	"testing"

	"github.com/go-redis/redismock/v9"
	redisStorage "ktbs.dev/mubeng/internal/storage/redis"
	"ktbs.dev/mubeng/pkg/model"
)

func TestAddProxy(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	proxyStore := redisStorage.NewProxyStorage("root", rdb)

	proxy := &model.Proxy{
		Address: "192.168.1.1:5050",
	}

	rawProxy, err := json.Marshal(proxy)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	// expectations
	mock.ExpectHSet("root", "192.168.1.1:5050", rawProxy).SetVal(1)

	if err := proxyStore.AddProxy(context.Background(), proxy); err != nil {
		t.Error("expected successfully adding proxy: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDeleteProxy(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	proxyStore := redisStorage.NewProxyStorage("root", rdb)

	mock.ExpectHDel("root", "192.168.1.1:5050").SetVal(1)

	if err := proxyStore.DeleteProxy(context.Background(), "192.168.1.1:5050"); err != nil {
		t.Error("unexpected error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetProxy(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	proxyStore := redisStorage.NewProxyStorage("root", rdb)

	mock.ExpectHGet("root", "192.168.1.1:5050").SetVal(`{"address":"192.168.1.1:5050","country":"DE","last_status":1}`)

	proxy, err := proxyStore.GetProxy(context.Background(), "192.168.1.1:5050")
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if proxy.Address != "192.168.1.1:5050" {
		t.Error("unexpected value on address: ", proxy.Address)
	}

	if proxy.Country != "DE" {
		t.Error("unexpected value on country: ", proxy.Country)
	}

	if proxy.LastStatus != model.ActiveStatus {
		t.Error("unexpected last status: ", proxy.LastStatus)
	}
}

func TestGetAllProxies(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	proxyStore := redisStorage.NewProxyStorage("root", rdb)

	mock.ExpectHGetAll("root").SetVal(map[string]string{
		"192.168.1.1:1000": `{"address":"192.168.1.1:1000","country":"AT","last_status":1}`,
		"192.168.1.1:2000": `{"address":"192.168.1.1:2000","country":"NL","last_status":2}`,
	})

	proxies, err := proxyStore.GetAllProxies(context.Background())
	if err != nil {
		t.Error("unexpected err: ", err)
	}

	if len(proxies) != 2 {
		t.Error("expected len:2, len =", len(proxies))
	}

	// sort slices
	sort.Slice(proxies, func(i, j int) bool {
		return proxies[i].Address < proxies[j].Address
	})

	if proxies[0].Address != "192.168.1.1:1000" {
		t.Error("unexpected value on address: ", proxies[0].Address)
	}

	if proxies[0].Country != "AT" {
		t.Error("unexpected value on country: ", proxies[0].Country)
	}

	if proxies[0].LastStatus != model.ActiveStatus {
		t.Error("unexpected last status: ", proxies[0].LastStatus)
	}

	if proxies[1].Address != "192.168.1.1:2000" {
		t.Error("unexpected value on address: ", proxies[1].Address)
	}

	if proxies[1].Country != "NL" {
		t.Error("unexpected value on country: ", proxies[1].Country)
	}

	if proxies[1].LastStatus != model.InactiveStatus {
		t.Error("unexpected last status: ", proxies[1].LastStatus)
	}
}
