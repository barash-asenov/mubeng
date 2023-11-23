package model

import "time"

type ProxyStatus int32

const (
	_ ProxyStatus = iota
	ActiveStatus
	InactiveStatus
)

type Proxy struct {
	Address    string        `json:"address"`
	Country    string        `json:"country"`
	Latency    time.Duration `json:"latency"`
	LastStatus ProxyStatus   `json:"last_status"`
	Source     string        `json:"source"`
	CheckCount int           `json:"check_count"`
	FailCount  int           `json:"fail_count"`
}
