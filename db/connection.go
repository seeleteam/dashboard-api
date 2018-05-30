/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/seeleteam/dashboard-api/db/pool"
)

// Connection to a remote InfluxDB server.
type Connection struct {
	host, database     string
	username, password string
	timeout            time.Duration
}

// Init init db
func Init() {
	if DBInitialSize <= 5 {
		DBInitialSize = 5
	}
	if DBMaxActive <= 20 {
		DBMaxActive = 20
	}
	if DBInitialSize <= 5 {
		DBInitialSize = 5
	}
	if int(DBIdleTimetout.Seconds()) <= 0 {
		DBIdleTimetout = 30 * time.Second
	}

	poolConfig := &pool.Config{
		// init conn size
		InitialSize: DBInitialSize,
		// the max conn size
		MaxActive: DBMaxActive,
		Factory:   factory,
		Close:     close,
		// the idle timeout
		IdleTimetout: DBIdleTimetout,
	}
	pool, err := pool.New(poolConfig)
	if err != nil {
		panic("dbpool init error")
	}
	dbPool = pool
}

func factory() (interface{}, error) {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     DBAddr,
		Username: DBUsername,
		Password: DBPassword,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer conn.Close()
	return conn, nil
}

func close(c interface{}) error {
	return c.(client.Client).Close()
}

// GetConn get influxdb client conn
func GetConn() client.Client {
	c, err := dbPool.Get()
	if err != nil {
		log.Fatal("get conn failed", err)
		return nil
	}
	return c.(client.Client)
}

// CloseConn release influxdb client conn
func CloseConn(conn client.Client) error {
	err := dbPool.Put(conn)
	if err != nil {
		log.Fatal("close conn failed", err)
		return err
	}
	return nil
}
