/**
* sql ref: https://docs.influxdata.com/influxdb/v1.5/query_language/data_exploration
*
* @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

// Condition metrics condition for db ops
type Condition struct {
	Fields []string // columns for influxdb without index, separated by comma

	Measurement string // table name for influxdb

	// append time condition with and
	WhereExpressions []string

	// where time >= {startTime} and time <= {endTime}, format like 2015-08-18T00:00:00Z,
	// 2015-08-18 00:12:00, 1439856000000000000, 1439856000s, 24043524m,
	// now() - 30s, now() - 1m, now() - 1d, ...
	// UTC time
	StartTime string
	// format like startTime
	EndTime string

	// if timeSince set startTime and endTime will be disabled
	TimeSince string

	Intervals       string // required, default 30s ex: 5s, 5m, 5h..., group by time(5s)
	IntervalsOffset string // should be same unit for intervals, if exist will append intervals with comma

	// format like "host"  "email", if multi separated by comma
	// follows the group by ...
	Tags []string

	// fill options, in linear, none, null, previous
	// linear - Reports the results of linear interpolation for time intervals with no data.
	// none - Reports no timestamp and no value for time intervals with no data.
	// null -  Reports null for time intervals with no data but returns a timestamp. This is the same as the default behavior.
	// previous -   Reports the value from the previous time interval for time intervals with no data.
	// number
	// // follows the group by ...
	FillOption string

	// default is asc, you can use desc replace it
	OrderBy string

	Limit  int // LIMIT <N> returns the first N points from the specified measurement.
	Offset int // OFFSET <N> paginates N points in the query results.

	SLimit  int // SLIMIT <N> returns every point from <N> series in the specified measurement.
	SOffset int // SOFFSET <N> paginates N series in the query results.

	TimeZone string // default use UTC
}
