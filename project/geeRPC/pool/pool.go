package pool

import (
	"container/list"
	"errors"
	"net"
	"sync"
	"time"
)

var (
	ErrConfigInvalid = errors.New("config is invalid")
)

type Pool interface {
	Get() (interface{}, error)
	Put(interface{}) error
}

type PoolConfig struct {
	MaxCap      int
	MaxIdle     int
	IdleTime    time.Duration
	MaxLifetime time.Duration
	Factory     func() (net.Conn, error)
}

func NewConnPool(config PoolConfig) (Pool, error) {
	if config.MaxIdle > config.MaxCap || config.Factory == nil {
		return nil, ErrConfigInvalid
	}

	connPool := &ConnPool{
		config: config,
		mutx:   sync.Mutex{},
		idle:   list.New(),
	}

	return connPool, nil
}
