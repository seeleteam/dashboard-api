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

	EnableWebSocket bool

	EnableGraceful    bool
	DefaultHammerTime time.Duration

	WebSocketURL    string
	LimitConnection int

	// HTTPCors is the Cross-Origin Resource Sharing header to send to requesting
	// clients. Please be aware that CORS is a browser enforced security, it's fully
	// useless for custom HTTP clients.
	HTTPCors []string

	// HTTPHostFilter is the whitelist of hostnames which are allowed on incoming requests.
	HTTPWhiteHost []string

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
