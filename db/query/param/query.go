/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package param

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/db"
)

const (
	defaultIntervals = "30s"
	defaultTimeSince = "5m"
	defaultLimit     = 100
)

// Query metrics query with params
type Query struct {
	Stmt string
}

// New get Query for metrics by input params
func New(cond *db.Condition) (*Query, error) {
	stmt, err := generateQueryStmt(cond)
	if err != nil {
		return nil, err
	}
	return &Query{
		Stmt: stmt,
	}, nil
}

// Query query data from db for meter
func (m *Query) Query() (res []client.Result, err error) {
	if m.Stmt == "" {
		return nil, errors.New("error query stmt")
	}
	return db.Query(m.Stmt)
}

// generateQueryStmt generate query string with input params
func generateQueryStmt(condition *db.Condition) (stmt string, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select ")
	if condition == nil {
		return "", errors.New("condition empty")
	}

	// fields or columns
	fields := condition.Fields
	if fields == nil || len(fields) == 0 {
		return "", fmt.Errorf("%s blank", common.RequestFields)
	}

	var fieldsStr string
	for _, field := range fields {
		if field != "" {
			if strings.HasSuffix(field, ",") {
				fieldsStr += field
			} else {
				fieldsStr += field + ","
			}
		}
	}

	if strings.HasSuffix(fieldsStr, ",") {
		fieldsStr = strings.TrimRight(fieldsStr, ",")
	}
	buffer.WriteString(fieldsStr + " ")

	// measurement or table
	measurement := condition.Measurement
	if measurement == "" {
		return "", fmt.Errorf("%s blank", common.RequestMeasurement)
	}

	measurement = strings.Trim(measurement, "\"")
	if measurement == "" {
		return "", fmt.Errorf("%s trim blank", common.RequestMeasurement)
	}
	buffer.WriteString("from ")
	buffer.WriteString("\"" + measurement + "\" ")

	//where condition related time
	buffer.WriteString("where 1=1 ")
	// where condition, should not include time section
	var whereConditionBuffer bytes.Buffer
	whereExpressions := condition.WhereExpressions

	if whereExpressions != nil && len(whereExpressions) != 0 {
		// each must one condition, not support multi nest
		for _, whereExpression := range whereExpressions {
			if whereExpression != "" {
				// good addr="127.0.0.1"
				// good addr="127.0.0.1" and hostname="chain"
				// bad  addr="127" and
				if strings.HasSuffix(whereExpression, "and") {
					return "", fmt.Errorf("error where condition with suffix and")
				}
				// default use and join two conditions
				whereConditionBuffer.WriteString(fmt.Sprintf("and %s ", whereExpression))
			}
		}
		// append where condition
		buffer.WriteString(whereConditionBuffer.String())
		whereConditionBuffer.Reset()
	}

	// must exist time condition append the where
	var timeCondition bytes.Buffer
	timeStart := condition.StartTime
	timeEnd := condition.EndTime
	timeSince := condition.TimeSince
	if timeStart != "" && timeEnd != "" {
		timeCondition.WriteString(fmt.Sprintf("and time >= %s ", timeStart))
		timeCondition.WriteString(fmt.Sprintf("and time <= %s ", timeEnd))
	} else if timeStart != "" {
		timeCondition.WriteString(fmt.Sprintf("and time >= %s ", timeStart))
	} else if timeEnd != "" {
		timeCondition.WriteString(fmt.Sprintf("and time <= %s ", timeEnd))
	} else {
		if timeSince == "" {
			timeSince = defaultTimeSince
		}
		timeCondition.WriteString(fmt.Sprintf("and time >= now() - %s ", timeSince))
	}

	buffer.WriteString(timeCondition.String())
	timeCondition.Reset()

	var groupByTimeCondition bytes.Buffer
	// must required, group by time(30s,offsetInterval)
	intervals := condition.Intervals
	if intervals != "" {
		// intervals = defaultIntervals
		intervalsOffset := condition.IntervalsOffset
		if intervalsOffset != "" {
			groupByTimeCondition.WriteString(fmt.Sprintf("time(%s,%s) ", intervals, intervalsOffset))
		} else {
			groupByTimeCondition.WriteString(fmt.Sprintf("time(%s) ", intervals))
		}
	}

	var groupTagsStr string
	// group by time(),tag1,tag2...
	groupTags := condition.Tags
	if groupTags != nil && len(groupTags) != 0 {
		for _, groupTag := range groupTags {
			if groupTag != "" {
				if strings.HasPrefix(groupTag, ",") {
					groupTagsStr += groupTag
				} else {
					groupTagsStr += "," + groupTag
				}
			}
		}
	}

	if groupByTimeCondition.Len() != 0 {
		buffer.WriteString("group by ")
		buffer.WriteString(groupByTimeCondition.String() + " ")
		groupByTimeCondition.Reset()
		buffer.WriteString(groupTagsStr + " ")
		if condition.FillOption != "" {
			// group by judge, if not exist any elem, fill with *
			buffer.WriteString(fmt.Sprintf("fill(%s) ", condition.FillOption))
		}
	} else {
		groupTagsStr = strings.TrimLeft(groupTagsStr, ",")
		if groupTagsStr != "" {
			buffer.WriteString("group by ")
			buffer.WriteString(groupTagsStr + " ")
			if groupTagsStr != "" && condition.FillOption != "" {
				// group by judge, if not exist any elem, fill with *
				buffer.WriteString(fmt.Sprintf("fill(%s) ", condition.FillOption))
			}
		}
	}

	// order by default asc
	if strings.ToLower(condition.OrderBy) == "desc" {
		buffer.WriteString(fmt.Sprintf("order by desc "))
	}

	// limit
	limit := condition.Limit
	if limit <= 0 {
		limit = defaultLimit
	}
	buffer.WriteString(fmt.Sprintf("limit %d ", limit))

	offset := condition.Offset
	if offset > 0 {
		buffer.WriteString(fmt.Sprintf("offset %d ", offset))
	}

	slimit := condition.SLimit
	if slimit > 0 {
		buffer.WriteString(fmt.Sprintf("slimit %d ", slimit))
	}

	soffset := condition.SOffset
	if soffset > 0 {
		buffer.WriteString(fmt.Sprintf("soffset %d ", offset))
	}

	// format like Asia/SHanghai
	zoneVal := condition.TimeZone
	if zoneVal != "" {
		buffer.WriteString(fmt.Sprintf("tz('%s')", zoneVal))
	}
	return buffer.String(), nil
}
