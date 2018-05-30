/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

// DB service for database
type DB interface {
	// query data from db
	Query() (res interface{}, err error)
}
