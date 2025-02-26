package pool

import (
	"log"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestNewConnPool(t *testing.T) {
	_, err := NewConnPool(PoolConfig{
		MaxCap:      10,
		MaxIdle:     10,
		IdleTime:    time.Minute,
		MaxLifetime: time.Hour,
		Factory: func() (net.Conn, error) {
			return net.Dial("tcp", "127.0.0.1:8080")
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestConnPool_Get(t *testing.T) {
	pool, err := NewConnPool(PoolConfig{
		MaxCap:      5,
		MaxIdle:     5,
		IdleTime:    time.Minute,
		MaxLifetime: time.Hour,
		Factory: func() (net.Conn, error) {
			return net.Dial("tcp", "127.0.0.1:8080")
		},
	})
	if err != nil {
		t.Error(err)
	}

	conn, err := pool.Get()
	if err != nil {
		t.Error(err)
	}
	log.Println(reflect.TypeOf(conn))
}

func TestConnPool_Put(t *testing.T) {
	pool, err := NewConnPool(PoolConfig{
		MaxCap:      5,
		MaxIdle:     5,
		IdleTime:    time.Minute,
		MaxLifetime: time.Hour,
		Factory: func() (net.Conn, error) {
			return net.Dial("tcp", "127.0.0.1:8080")
		},
	})
	if err != nil {
		t.Error(err)
	}

	conn, err := pool.Get()
	if err != nil {
		t.Error(err)
	}
	log.Println(reflect.TypeOf(conn))
	err = pool.Put(conn)
	if err != nil {
		t.Error(err)
	}
}

func TestConnPool(t *testing.T) {
	pool, err := NewConnPool(PoolConfig{
		MaxCap:      5,
		MaxIdle:     5,
		IdleTime:    time.Minute,
		MaxLifetime: time.Hour,
		Factory: func() (net.Conn, error) {
			return net.Dial("tcp", "127.0.0.1:8080")
		},
	})
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 20; i++ {
		go func(id int) {
			conn, err := pool.Get()
			if err != nil {
				return
			}
			time.Sleep(time.Millisecond * 500)
			defer pool.Put(conn)

			log.Println("pool get conn,do something...", id)
		}(i)
	}

	time.Sleep(10 * time.Second)
}
