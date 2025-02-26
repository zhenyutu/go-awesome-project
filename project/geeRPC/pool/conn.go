package pool

import (
	"container/list"
	"errors"
	"net"
	"sync"
	"time"
)

var (
	ErrSelect = errors.New("")
)

type PoolConn struct {
	conn       net.Conn
	pool       *ConnPool
	updateTime time.Time
}

func (p *ConnPool) Warp(c net.Conn) *PoolConn {
	return &PoolConn{
		conn:       c,
		pool:       p,
		updateTime: time.Now(),
	}
}

func (conn *PoolConn) TurnOff() error {
	conn.updateTime = time.Now()
	if err := conn.pool.Put(conn); err != nil {
		return conn.Close()
	}
	return nil
}

func (conn *PoolConn) Close() error {
	err := conn.conn.Close()
	if err != nil {
		return err
	}

	conn.pool.running--
	return nil
}

type ConnReq struct {
	conn *PoolConn
}

type ConnPool struct {
	addr     string
	mutx     sync.Mutex
	config   PoolConfig
	idle     *list.List
	connReqs []chan ConnReq
	running  int
}

func (c *ConnPool) Get() (interface{}, error) {
	c.mutx.Lock()
	defer c.mutx.Unlock()

	//判断是否有空闲链接，若存在直接返回
	for c.idle.Len() > 0 {
		pc := c.idle.Remove(c.idle.Front()).(*PoolConn)
		if time.Now().Sub(pc.updateTime) < c.config.MaxLifetime {
			return pc.conn, nil
		}

		go pc.TurnOff()
	}

	//无空闲链接，可以创建链接则创建返回
	if c.running < c.config.MaxCap {
		conn, err := c.config.Factory()
		if err != nil {
			return nil, err
		}

		c.running++
		pc := c.Warp(conn)
		return pc.conn, nil
	}

	//无空闲链接，不可创建链接则阻塞
	if c.running >= c.config.MaxCap {
		req := make(chan ConnReq, 1)
		c.connReqs = append(c.connReqs, req)
		c.mutx.Unlock()

		for {
			res, ok := <-req
			if !ok {
				return nil, errors.New("conn req select error")
			}
			pc := res.conn
			if time.Now().Sub(pc.updateTime) > c.config.MaxLifetime {
				pc.TurnOff()
				continue
			}
			return pc.conn, nil
		}
	}

	return nil, errors.New("conn get error")
}

func (c *ConnPool) Put(conn interface{}) error {
	c.mutx.Lock()
	defer c.mutx.Unlock()
	nc := conn.(net.Conn)

	//如果存在阻塞线程则处理通道请求
	if len(c.connReqs) > 0 {
		req := c.connReqs[0]
		c.connReqs = c.connReqs[1:]

		req <- ConnReq{
			conn: &PoolConn{
				conn:       nc,
				pool:       c,
				updateTime: time.Now(),
			},
		}
	}

	//无阻塞线程，判断空闲线程是否多于配置
	if c.idle.Len() > c.config.MaxIdle {
		nc.Close()
		return nil
	}

	c.idle.PushBack(conn)
	return nil
}

func NewConnPool(addr string, config PoolConfig) (Pool, error) {
	if config.InitialCap > config.MaxCap || config.Factory == nil {
		return nil, ErrConfigInvalid
	}

	connPool := &ConnPool{
		addr:   addr,
		config: config,
		mutx:   sync.Mutex{},
		idle:   list.New(),
	}

	return connPool, nil
}
