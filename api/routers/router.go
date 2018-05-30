/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/api/handlers"
	"github.com/seeleteam/dashboard-api/common"
)

// InitRouters init routers
func InitRouters(e *gin.Engine) {
	// set api handlers logger
	handlers.SetAPIHandlerLog("api-handlers", common.PrintLog)

	e.GET("/ping", handlers.Ping())
	e.GET("/pong", handlers.Pong())
	e.GET("/async", handlers.LongAsync())
}
