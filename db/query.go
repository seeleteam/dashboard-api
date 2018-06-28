/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

import (
	"github.com/influxdata/influxdb/client/v2"
)

// queryDB convenient function to query the database
// The precision arguments can be empty strings if they are not needed for the query.
func queryDB(clnt client.Client, cmd, precision string) (res []client.Result, err error) {
	return queryDBWithParameters(clnt, cmd, precision, nil)
}

// queryDBWithParameters convenient function to query the database
// The precision arguments can be empty strings if they are not needed for the query. format in 'rfc3339|h|m|s|ms|u|ns'
// parameters is a map of the parameter names used in the command to their values.
func queryDBWithParameters(clnt client.Client, cmd, precision string, parameters map[string]interface{}) (res []client.Result, err error) {
	q := client.Query{
		Command:    cmd,
		Database:   DBNAME,
		Precision:  precision,
		Parameters: parameters,
	}
	defer CloseConn(clnt)
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

// Query with query and precision string
func Query(cmd string, precision string) (res []client.Result, err error) {
	client := GetConn()
	return queryDB(client, cmd, precision)
}
