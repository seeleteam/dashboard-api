/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

import (
	"github.com/seeleteam/dashboard-api/common/query"
)

// GenerateQueryMeter generate influxdb query sql for meter
func GenerateQueryMeter(condition *query.Condition) string {
	// queryCount := fmt.Sprintf("select %s from \"%s\" where time >= now() - %s group by time(%s%s) %s fill(%s) limit %d %s;",

	return ""
}
