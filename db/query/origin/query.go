/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package origin

import (
	"errors"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/seeleteam/dashboard-api/db"
)

// Query metrics query
type Query struct {
	Stmt      string // query string
	Precision string // h,m,s,ms,u,ns and ""
}

// New get Query for metrics by input sql
func New(stmt string) *Query {
	return NewWithPrecision(stmt, "")
}

// NewWithPrecision get Query for metrics by input sql and precision
func NewWithPrecision(stmt string, precision string) *Query {
	return &Query{
		Stmt:      stmt,
		Precision: precision,
	}
}

// Query query data from db for meter
func (m *Query) Query() (res []client.Result, err error) {
	if m.Stmt == "" {
		return nil, errors.New("error query stmt")
	}
	return db.Query(m.Stmt, m.Precision)
}
