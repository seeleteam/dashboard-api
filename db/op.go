package db

import (
	"github.com/influxdata/influxdb/client/v2"
)

// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// Query with query string
func Query(cmd string) (res []client.Result, err error) {
	client := GetConn()
	defer CloseConn(client)
	return queryDB(client, cmd)
}
