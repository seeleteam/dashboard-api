/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package api

import (
	"time"
)

// Config config the api
type Config struct {
	Name    string
	Version string

	// HTTP Addr
	ListenAddr string

	LimitConnection int

	//RunMode, ex: debug,release,test
	RunMode string

	RootRouterPrefix string

	EnableHTTPS bool
	// HTTPS Addr
	HTTPSAddr string

	// TLS
	CertFile string
	KeyFile  string

	// http server config
	ServerConfig *ServerConfig
}

// ServerConfig http server config
type ServerConfig struct {
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	MaxHeaderBytes int
}
