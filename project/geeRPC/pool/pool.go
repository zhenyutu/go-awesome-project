package pool

import (
	"errors"
	"net"
	"time"
)

var (
	ErrClosed        = errors.New("pool is closed")
	ErrConfigInvalid = errors.New("config is invalid")
)

type Pool interface {
	Get() (interface{}, error)
	Put(interface{}) error
}

type PoolConfig struct {
	InitialCap  int
	MaxCap      int
	MaxIdle     int
	IdleTime    time.Duration
	MaxLifetime time.Duration
	Factory     func() (net.Conn, error)
}
