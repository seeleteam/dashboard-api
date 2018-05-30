package pool

import (
	"errors"
)

var (
	// ErrClosed is the error resulting if the pool is closed via pool.Close() or pool.Release().
	ErrClosed = errors.New("pool is closed")
)

// Pool dv conn pool
type Pool interface {
	// Get get conn from pool
	Get() (interface{}, error)

	// Put conn into pool
	Put(interface{}) error

	// Close the conn
	Close(interface{}) error

	// Release release all conn in pool
	Release()

	// Size the size of the pool conn
	Size() int
}
