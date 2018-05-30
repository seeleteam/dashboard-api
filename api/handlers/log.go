/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"github.com/seeleteam/dashboard-api/common"
	dlog "github.com/seeleteam/dashboard-api/log"
)

var (
	log *dlog.GlobalLog
)

// SetAPIHandlerLog set api handler logger
func SetAPIHandlerLog(name string, printLog bool) {
	log = dlog.GetLogger("api-handlers", common.PrintLog)
}
