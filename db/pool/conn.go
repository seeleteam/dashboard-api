package pool

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Config the connection pool config
type Config struct {
	// init conn size
	InitialSize int
	// the max conn size
	MaxActive int
	// the min idele conn size
	MinIdle int

	Factory func() (interface{}, error)
	Close   func(interface{}) error
	// the idle timeout
	IdleTimetout time.Duration
}

type connPool struct {
	mu          sync.Mutex
	conns       chan *idleConn
	factory     func() (interface{}, error)
	close       func(interface{}) error
	idleTimeout time.Duration
}

type idleConn struct {
	conn interface{}
	t    time.Time
}

// New create a new Pool config
func New(c *Config) (Pool, error) {
	if c.InitialSize < 0 || c.MaxActive <= 0 || c.InitialSize > c.MaxActive {
		return nil, errors.New("invalid size config")
	}

	connPool := &connPool{
		conns:       make(chan *idleConn, c.MaxActive),
		factory:     c.Factory,
		close:       c.Close,
		idleTimeout: c.IdleTimetout,
	}

	for i := 0; i < c.InitialSize; i++ {
		conn, err := connPool.factory()
		if err != nil {
			connPool.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		connPool.conns <- &idleConn{conn: conn, t: time.Now()}
	}

	return connPool, nil

}

//getConns get all connection
func (c *connPool) getConns() chan *idleConn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

//Get get one connection from the conn pool
func (c *connPool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, ErrClosed
	}
	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, ErrClosed
			}
			// if timeout, drop the conn
			if timeout := c.idleTimeout; timeout > 0 {
				if wrapConn.t.Add(timeout).Before(time.Now()) {
					// drop and close the conn
					c.Close(wrapConn.conn)
					continue
				}
			}
			return wrapConn.conn, nil
		default:
			conn, err := c.factory()
			if err != nil {
				return nil, err
			}

			return conn, nil
		}
	}
}

//Put put the conn into pool again
func (c *connPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conns == nil {
		return c.Close(conn)
	}

	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		return nil
	default:
		// conn pool is full, drop and close the conn
		return c.Close(conn)
	}
}

//Close close the current conn
func (c *connPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.close(conn)
}

//Release release all conns in the pool
func (c *connPool) Release() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.factory = nil
	closeFun := c.close
	c.close = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		closeFun(wrapConn.conn)
	}
}

// Size the count of conn in the pool
func (c *connPool) Size() int {
	return len(c.getConns())
}
