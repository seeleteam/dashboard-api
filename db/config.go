/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

import (
	"time"

	"github.com/seeleteam/dashboard-api/db/pool"
)

// Config config the api
type Config struct {
	// NAME db name
	NAME string
	// Addr db addr
	Addr string
	// Username  db username
	Username string
	// Password db password
	Password string
	// InitialSize ini pool size
	InitialSize int
	// MaxActive max active conn in pool
	MaxActive int
	// IdleTimetout the conn max idle time
	IdleTimetout time.Duration
}

var (
	// DBPool influxdb pool
	dbPool pool.Pool

	// DBNAME db name
	DBNAME = "influxdb"
	// DBAddr db addr
	DBAddr = "http://localhost:8086"
	// DBUsername  db username
	DBUsername = "test"
	// DBPassword db password
	DBPassword = "test123"
	// DBInitialSize ini pool size
	DBInitialSize int
	// DBMaxActive max active conn in pool
	DBMaxActive int
	// DBIdleTimetout the conn max idle time
	DBIdleTimetout time.Duration
)
