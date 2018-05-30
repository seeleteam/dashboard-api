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

	// routerGroup API
	rootRouterPrefix := common.RootRouterPrefix
	if rootRouterPrefix == "" {
		rootRouterPrefix = "/api"
	}
	routerGroupAPI := e.Group(rootRouterPrefix)

	apiShowGroup := routerGroupAPI.Group("/show")
	apiShowGroup.GET("/databases", handlers.ShowDatabases())
	apiShowGroup.GET("/measurements", handlers.ShowMeasurements())
	apiShowGroup.GET("/tagKeys", handlers.ShowTagKeys())
	apiShowGroup.GET("/tagValues", handlers.ShowTagValues())
	apiShowGroup.GET("/fieldKeys", handlers.ShowFieldKeys())

	// base sql in group api
	apiSelectGroup := routerGroupAPI.Group("/query")
	apiSelectGroup.GET("/sqls", handlers.SelectBySQLs())
	apiSelectGroup.GET("/params", handlers.SelectWithParams())

}
