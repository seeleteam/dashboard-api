package query

// Condition meter condition for query
type Condition struct {
	fields string // columns for influxdb without index, separated by comma
	// tags        string // columns for influxdb with index, separated by comma

	measurement string // table name for influxdb

	// where time >= {startTime} and time <= {endTime}, format like 2015-08-18T00:00:00Z,
	// 2015-08-18 00:12:00, 1439856000000000000, 1439856000s, 24043524m,
	// now() - 30s, now() - 1m, now() - 1d, ...
	// UTC time
	startTime string
	// format like startTime
	endTime string

	// if timeSince set startTime and endTime will be disabled
	timeSince string

	intervals       string // required, default 30s ex: 5s, 5m, 5h..., group by time(5s)
	intervalsOffset string // should be same unit for intervals, if exist will append intervals with comma

	// at least one tag in here, multi separated by comma, usage: group by time(30s) {groupTags}
	// groupTags format like ,hosttimil
	groupTags string

	// fill options, in linear, none, null, previous
	// linear - Reports the results of linear interpolation for time intervals with no data.
	// none - Reports no timestamp and no value for time intervals with no data.
	// null -  Reports null for time intervals with no data but returns a timestamp. This is the same as the default behavior.
	// previous -   Reports the value from the previous time interval for time intervals with no data.
	fillOption string

	limit    int    // LIMIT <N> returns the first N points from the specified measurement.
	timeZone string // default use UTC
}
