/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package param

import (
	"testing"

	"github.com/seeleteam/dashboard-api/query"
)

func Test_New(t *testing.T) {
	condition := &query.Condition{
		Measurement:     "chain.block.insert.meter",
		Fields:          []string{"*"},
		Tags:            nil,
		TimeSince:       "5d",
		StartTime:       "",
		EndTime:         "",
		Limit:           3,
		Offset:          0,
		SLimit:          0,
		SOffset:         0,
		TimeZone:        "Asia/Shanghai",
		Intervals:       "30s",
		IntervalsOffset: "",
	}
	query, err := New(condition)
	if err != nil {
		t.Fatalf("error %v", err)
		return
	}
}
