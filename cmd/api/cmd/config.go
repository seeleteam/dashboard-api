/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/api"
	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/db"
)

// Config aggregates all configs exposed to users
// Note to add enough comments for every field
type Config struct {
	api.Config
	DB *db.Config

	DisableConsoleColor bool

	// If LogLevel is set and corret it will be LogLevel, otherwise use DebugLevel
	LogLevel string

	// If PrintLog is true, all logs will be printed in the console, otherwise they will be stored in the file.
	PrintLog bool

	WriteLog bool

	LogDepth int
}

// GetConfigFromFile unmarshals the config from the given file
func GetConfigFromFile(filepath string) (Config, error) {
	var config Config
	buff, err := ioutil.ReadFile(filepath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(buff, &config)
	return config, err
}

// LoadConfigFromFile gets node config from the given file
func LoadConfigFromFile(configFile string) (*api.Config, error) {
	config, err := GetConfigFromFile(configFile)
	if err != nil {
		return nil, err
	}

	// common config
	common.DisableConsoleColor = config.DisableConsoleColor
	common.LogLevel = config.LogLevel
	common.PrintLog = config.PrintLog
	common.RootRouterPrefix = config.RootRouterPrefix
	common.WriteLog = config.WriteLog
	if config.LogDepth > 0 {
		common.LogDepth = config.LogDepth
	}

	// api config
	apiConfig := new(api.Config)

	apiConfig.Name = config.Name
	apiConfig.Version = config.Version

	apiConfig.ListenAddr = config.ListenAddr

	runMode := config.RunMode
	if runMode != gin.ReleaseMode && runMode != gin.DebugMode && runMode != gin.TestMode {
		runMode = gin.ReleaseMode
	}
	apiConfig.RunMode = runMode

	apiConfig.LimitConnection = config.LimitConnection

	// server config
	serverConfig := config.ServerConfig
	serverConfig.IdleTimeout = config.ServerConfig.IdleTimeout
	serverConfig.MaxHeaderBytes = config.ServerConfig.MaxHeaderBytes
	serverConfig.ReadTimeout = config.ServerConfig.ReadTimeout
	serverConfig.WriteTimeout = config.ServerConfig.WriteTimeout

	apiConfig.ServerConfig = serverConfig

	// db config
	dbConfig := config.DB
	if dbConfig != nil {
		db.DBNAME = dbConfig.NAME
		db.DBAddr = dbConfig.Addr
		db.DBUsername = dbConfig.Username
		db.DBPassword = dbConfig.Password
		db.DBInitialSize = dbConfig.InitialSize
		db.DBMaxActive = dbConfig.MaxActive
		db.DBIdleTimetout = dbConfig.IdleTimetout
	}
	return apiConfig, nil
}
