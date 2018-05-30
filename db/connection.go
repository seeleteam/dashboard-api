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

var (
	// DBPool influxdb pool
	dbPool pool.Pool

	// MyDB influxdb db name
	MyDB = "influxdb"
	// Addr influxdb addr
	Addr = "http://localhost:8086"
	// username influxdb username
	username = "test"
	// password influxdb password
	password = "test123"

	// InitialSize ini pool size
	InitialSize int
	// MaxActive max active conn in pool
	MaxActive int
	// MinIdle the min idle conn
	MinIdle int
	// IdleTimetout the conn max idle time
	IdleTimetout time.Duration
)

func init() {
	if InitialSize <= 5 {
		InitialSize = 5
	}
	if MaxActive <= 20 {
		MaxActive = 20
	}
	if MinIdle <= 5 {
		MinIdle = 5
	}
	if InitialSize <= 5 {
		InitialSize = 5
	}
	if int(IdleTimetout.Seconds()) < 0 {
		IdleTimetout = 30 * time.Second
	}

	poolConfig := &pool.Config{
		// init conn size
		InitialSize: InitialSize,
		// the max conn size
		MaxActive: MaxActive,
		// the min idele conn size
		MinIdle: MinIdle,

		Factory: factory,
		Close:   close,
		// the idle timeout
		IdleTimetout: IdleTimetout,
	}
	pool, err := pool.New(poolConfig)
	if err != nil {
		panic("dbpool init error")
	}
	dbPool = pool
}

func factory() (interface{}, error) {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     Addr,
		Username: username,
		Password: password,
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
