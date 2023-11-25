package model

import (
	"encoding/json"
	"errors"
	"time"
)

type ProxyStatus int32

const (
	_ ProxyStatus = iota
	ActiveStatus
	InactiveStatus
)

type ProxyLatency struct {
	time.Duration
}

func (d ProxyLatency) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d ProxyLatency) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

type Proxy struct {
	Address    string       `json:"address"`
	Country    string       `json:"country"`
	Latency    ProxyLatency `json:"latency"`
	LastStatus ProxyStatus  `json:"last_status"`
	Source     string       `json:"source"`
	CheckCount int          `json:"check_count"`
	FailCount  int          `json:"fail_count"`
}
