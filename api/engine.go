/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package api

import (
	"net/http"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/api/routers"
	"github.com/seeleteam/dashboard-api/db"
	"github.com/seeleteam/dashboard-api/log"
)

// EngineConfig engine config
type EngineConfig struct {
	middleware       []func(*gin.Context)
	log              *log.GlobalLog
	LimitConnections int
	RunMode          string // runMode, ex: debug,release,test
	RootRouterPrefix string // root router, default ""
}

// initEngineConfig init engine config
func (config *EngineConfig) initEngineConfig() *gin.Engine {
	if config == nil {
		panic("engine config should not be nil")
	}

	// set gin mode release(hide handlers info)
	gin.SetMode(config.RunMode)

	e := gin.New()

	e.Use(gzip.Gzip(gzip.DefaultCompression))

	// gin api handlers, used for api log info
	e.Use(log.Logger(log.GetLoggerWithCaller("gin-handlers", true, false).GetLogger()))

	// use logs middleware logurs
	// e.Use(gin.Logger())

	// use recovery middleware
	e.Use(gin.Recovery())

	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	corsConfig.AllowAllOrigins = true
	e.Use(cors.New(corsConfig))

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

	// init the db
	db.Init()
	// here init the routers, need refactor
	routers.InitRouters(e)
	return e
}
