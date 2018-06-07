/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package api

import (
	"golang.org/x/sync/errgroup"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/log"
)

var (
	eGroup errgroup.Group
)

// API api server config
type API struct {
	config     *Config
	log        *log.GlobalLog
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
	//run server with config and logs
	server := GetServer(api)
	return server.Run()
}
