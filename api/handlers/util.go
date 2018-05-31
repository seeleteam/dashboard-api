/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
)

// ResponseJSON response json
func ResponseJSON(c *gin.Context, responseData *common.ResponseData) {
	code := responseData.Code
	c.JSON(code, responseData)
}
