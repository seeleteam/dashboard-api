/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package api

import (
	"net/http"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/api/routers"
	"github.com/seeleteam/dashboard-api/log"
)

// EngineConfig engine config
type EngineConfig struct {
	middleware       []func(*gin.Context)
	log              *log.GlobalLog
	LimitConnections int
}

// initEngineConfig init engine config
func (config *EngineConfig) initEngineConfig() *gin.Engine {
	if config == nil {
		panic("engine config should not be nil")
	}

	// set gin mode release(hide handlers info)
	gin.SetMode(gin.ReleaseMode)

	e := gin.New()

	// gin api handlers, used for api log info
	e.Use(log.Logger(log.GetLoggerWithCaller("gin-handlers", true, false).GetLogger()))

	// use logs middleware logurs
	// e.Use(gin.Logger())

	// use recovery middleware
	e.Use(gin.Recovery())

	// By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded number of requests.
	// Limiting the number of simultaneous connections can sometimes greatly speed things up under load
	if config.LimitConnections > 0 {
		e.Use(limit.MaxAllowed(config.LimitConnections))
	}

	return e
}

// Init engine init
func (config *EngineConfig) Init() http.Handler {
	e := config.initEngineConfig()

	// here init the routers, need refactor
	routers.InitRouters(e)
	return e
}
