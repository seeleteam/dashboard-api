/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */
package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
)

// Ping return without content
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseData := common.NewResponseData(200, "server is ok", nil, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
