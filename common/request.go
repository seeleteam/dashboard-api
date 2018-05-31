/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

/**
* Base Select Section
*
* SELECT <field_key>[,<field_key>,<tag_key>] FROM <measurement_name>[,<measurement_name>]
*
 */
const (
	// RequestSelectFields columns without index in influxdb, if multi(array) separated by comma
	RequestFields = "fields"

	// RequestMeasurement table names in influxdb, if multi(array) separated by comma
	RequestMeasurement = "measurement"
)

/**
* Where Section
*
* SELECT_clause FROM_clause WHERE <conditional_expression> [(AND|OR) <conditional_expression> [...]]
* The WHERE clause supports conditional_expressions on fields, tags, and timestamps.
*
* field: field_key <operator> ['string' | boolean | float | integer]
* tags: tag_key <operator> ['tag_value']
* timestamps:
* 	For most SELECT statements, the default time range is between 1677-09-21 00:12:43.145224194 and 2262-04-11T23:47:16.854775806Z UTC.
* 	For SELECT statements with a GROUP BY time() clause, the default time range is between 1677-09-21 00:12:43.145224194 UTC and now()
*
 */
const (
	// RequestWhereExpressions where expression content in request
	// WHERE clause: =   equal to <> not equal to != not equal to =~ matches against !~ doesn’t match against
	// include all where conditional_expression, if multi(array) separated by comma
	RequestWhereExpressions = "whereExps"

	// RequestStartTime startTime field in request
	RequestStartTime = "startTime"

	// RequestEndTime endTime field in request
	RequestEndTime = "endTime"

	// RequestTimeSince timeSince field in request
	RequestTimeSince = "timeSince"

	// RequestIntervals intervals field in request
	RequestIntervals = "intervals"

	// RequestIntervalsOffset intervalsOffset field in request
	RequestIntervalsOffset = "intervalsOffset"
)

/*
* Group Section
*
* 1. Syntax
* SELECT_clause FROM_clause [WHERE_clause] GROUP BY [* | <tag_key>[,<tag_key]]
*
* 2. GROUP BY time intervals
* SELECT <function>(<field_key>) FROM_clause WHERE <time_range> GROUP BY time(<time_interval>),[tag_key] [fill(<fill_option>)]
*
* 3. Advanced GROUP BY time() Syntax
* SELECT <function>(<field_key>) FROM_clause WHERE <time_range> GROUP BY time(<time_interval>,<offset_interval>),[tag_key] [fill(<fill_option>)]
*
* 4. GROUP BY time intervals and fill()
* SELECT <function>(<field_key>) FROM_clause WHERE <time_range> GROUP BY time(time_interval,[<offset_interval])[,tag_key] [fill(<fill_option>)]
 */
const (
	// RequestGroupBys group by, if multi(array) separated by comma
	RequestGroupBys = "groups"
	// RequestTag groupTag field in request
	// warn: only support single tag or not
	RequestTags = "tags"

	// RequestFillOption fill field in request
	RequestFillOption = "fill"
)

/*
* Order By Section
*
* By default, InfluxDB returns results in ascending time order;
* the first point returned has the oldest timestamp and the last point returned has the most recent timestamp.
* ORDER BY time DESC reverses that order such that InfluxDB returns the points with the most recent timestamps first.
*
* SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] [GROUP_BY_clause] ORDER BY time DESC
*
 */
const (
	// RequestOrderBy order field in request
	RequestOrderBy = "order"
)

/*
* LIMIT, SLIMIT Section
*
* LIMIT and SLIMIT limit the number of points and the number of series returned per query.
*
* 1. LIMIT <N> returns the first N points from the specified measurement.
*	SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] [GROUP_BY_clause] [ORDER_BY_clause] LIMIT <N>
* N specifies the number of points to return from the specified measurement. If N is greater than the number of points in a measurement,
* InfluxDB returns all points from that measurement.
*
* 2. SLIMIT <N> returns every point from <N> series in the specified measurement.
*	SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] GROUP BY *[,time(<time_interval>)] [ORDER_BY_clause] SLIMIT <N>
* N specifies the number of series to return from the specified measurement. If N is greater than the number of series in a measurement,
* InfluxDB returns all series from that measurement.
*
* 3. LIMIT and SLIMIT
*	LIMIT <N> followed by SLIMIT <N> returns the first <N> points from <N> series in the specified measurement.
* SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] GROUP BY *[,time(<time_interval>)] [ORDER_BY_clause] LIMIT <N1> SLIMIT <N2>
*
* N1 specifies the number of points to return per measurement. If N1 is greater than the number of points in a measurement,
* InfluxDB returns all points from that measurement.
* N2 specifies the number of series to return from the specified measurement. If N2 is greater than the number of series in a measurement,
* InfluxDB returns all series from that measurement.
*
 */
const (
	// RequestLimit limit field in request
	RequestLimit = "limit"

	// RequestLimit limit field in request
	RequestSLimit = "slimit"
)

/*
* OFFSET and SOFFSET Section
*
* OFFSET and SOFFSET paginates points and series returned.
*
* 1. OFFSET <N> paginates N points in the query results.
*	SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] [GROUP_BY_clause] [ORDER_BY_clause] LIMIT_clause OFFSET <N> [SLIMIT_clause]
* N specifies the number of points to paginate. The OFFSET clause requires a LIMIT clause.
* Using the OFFSET clause without a LIMIT clause can cause inconsistent query results.
*
* 2. SOFFSET <N> paginates N series in the query results.
*	SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] GROUP BY *[,time(time_interval)] [ORDER_BY_clause] [LIMIT_clause] [OFFSET_clause]
* SLIMIT_clause SOFFSET <N>
* N specifies the number of series to paginate. The SOFFSET clause requires an SLIMIT clause. Using the SOFFSET clause without an SLIMIT clause can cause inconsistent query results.
* There is an ongoing issue that requires queries with SLIMIT to include GROUP BY *
*
 */
const (
	// RequestOffset offset field in request
	RequestOffset = "offset"

	// RequestSOffset soffset field in request
	RequestSOffset = "soffset"
)

/*
* Time Zone Section
*
* The tz() clause returns the UTC offset for the specified timezone.
*
* SELECT_clause [INTO_clause] FROM_clause [WHERE_clause] [GROUP_BY_clause] [ORDER_BY_clause] [LIMIT_clause] [OFFSET_clause] [SLIMIT_clause] [SOFFSET_clause] tz('<time_zone>')
* Description of Syntax
*
* 1. Absolute Time
* SELECT_clause FROM_clause WHERE time <operator> ['<rfc3339_date_time_string>' | '<rfc3339_like_date_time_string>'
* | <epoch_time>] [AND ['<rfc3339_date_time_string>' | '<rfc3339_like_date_time_string>' | <epoch_time>] [...]]
*
* Supported operators
* =   equal to <> not equal to != not equal to >   greater than >= greater than or equal to <   less than <= less than or equal to
* Currently, InfluxDB does not support using OR with absolute time in the WHERE clause.
*
* 2. Relative Time
* Use now() to query data with timestamps relative to the server’s current timestamp.
* SELECT_clause FROM_clause WHERE time <operator> now() [[ - | + ] <duration_literal>] [(AND|OR) now() [...]]
*
* now() is the Unix time of the server at the time the query is executed on that server. The whitespace between - or + and the duration literal is required.
*
* Supported operators
* =   equal to <> not equal to != not equal to >   greater than >= greater than or equal to <   less than <= less than or equal to
*
* duration_literal
* u or µ microseconds ms       milliseconds s      seconds m      minutes h      hours d      days w      weeks

 */

const (
	// RequestTimeZone timeZone field in request
	RequestTimeZone = "tz"
)

// Common Section
const (
	// RequestDataBase database name
	RequestDataBase = "db"

	// RequestWithExpression WITH MEASUREMENT <regular_expression>
	// used for show clause
	RequestWithExpression = "withExp"

	// RequestSQL the sql or sql array field in request
	RequestSQL = "sql"
)
