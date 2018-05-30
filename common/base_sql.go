/**
*  ref: https://docs.influxdata.com/influxdb/v1.5/query_language/schema_exploration
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

const (
	// ShowDatabases SHOW DATABASES
	ShowDatabases = "SHOW DATABASES"

	// ShowRetentionPolices show retention polices
	// ref: SHOW RETENTION POLICIES [ON <database_name>]
	ShowRetentionPolices = "SHOW RETENTION POLICIES"

	// ShowSeries show series
	// ref: SHOW SERIES [ON <database_name>] [FROM_clause] [WHERE <tag_key> <operator> [ '<tag_value>' | <regular_expression>]] [LIMIT_clause] [OFFSET_clause]
	// WHERE clause: =   equal to <> not equal to != not equal to =~ matches against !~ doesn’t match against
	// Regular expressions are surrounded by / characters
	// here we only support one where condition
	ShowSeries = "SHOW SERIES"

	// ShowMeasurements show measurements like table
	// ref: SHOW MEASUREMENTS [ON <database_name>] [WITH MEASUREMENT <regular_expression>] [WHERE <tag_key> <operator> ['<tag_value>' | <regular_expression>]] [LIMIT_clause] [OFFSET_clause]
	// here we only support one where condition
	ShowMeasurements = "SHOW MEASUREMENTS"

	// ShowTagKeys show tag key(columns with index)
	// ref: SHOW TAG KEYS [ON <database_name>] [FROM_clause] [WHERE <tag_key> <operator> ['<tag_value>' | <regular_expression>]] [LIMIT_clause] [OFFSET_clause]
	ShowTagKeys = "SHOW TAG KEYS"

	// ShowTagValues show tag values
	// ref: SHOW TAG VALUES [ON <database_name>][FROM_clause] WITH KEY [ [<operator> "<tag_key>" | <regular_expression>] | [IN ("<tag_key1>","<tag_key2")]] [WHERE <tag_key> <operator> ['<tag_value>' | <regular_expression>]] [LIMIT_clause] [OFFSET_clause]
	// SHOW TAG VALUES [ON <database_name>][FROM_clause] WITH KEY [ [<operator> "<tag_key>" | <regular_expression>] | [IN ("<tag_key1>","<tag_key2")]] [WHERE <tag_key> <operator> ['<tag_value>' | <regular_expression>]] [LIMIT_clause] [OFFSET_clause]
	// Supported operators in the WITH and WHERE clauses: =   equal to <> not equal to != not equal to =~ matches against !~ doesn’t match against
	ShowTagValues = "SHOW TAG VALUES"

	// ShowFieldKeys show field keys
	// ref: SHOW FIELD KEYS [ON <database_name>] [FROM <measurement_name>]
	ShowFieldKeys = "SHOW FIELD KEYS"
)
