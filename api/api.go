/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package api

import (
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/log"
)

// error infos
var (
	ErrConfigIsNull       = errors.New("config info is null")
	ErrLogIsNull          = errors.New("APILogs is null")
	ErrNodeRunning        = errors.New("API is already running")
	ErrNodeStopped        = errors.New("API is not started")
	ErrServiceStartFailed = errors.New("API service start failed")
	ErrServiceStopFailed  = errors.New("API service stop failed")
)

var (
	eGroup errgroup.Group
)

// API api server config
type API struct {
	config *Config

	log  *log.GlobalLog
	lock sync.RWMutex

	ErrorGroup *errgroup.Group
}

// New generate API config
func New(conf *Config) (*API, error) {
	confCopy := *conf
	alog := log.GetLogger("api", common.PrintLog)
	return &API{
		config:     &confCopy,
		log:        alog,
		ErrorGroup: &eGroup,
	}, nil
}

// Start start the api server
func (api *API) Start() error {
	api.lock.Lock()
	defer api.lock.Unlock()

	//run server with config and logs
	server := GetServer(api)
	server.Run()
	return nil
}
