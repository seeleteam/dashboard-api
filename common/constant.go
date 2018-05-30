/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

var (
	// tempFolder used to store temp file, such as log files
	tempFolder string

	// defaultDataFolder used to store persistent data info, such as the database and keystore
	defaultDataFolder string

	// DisableConsoleColor disable the console color
	DisableConsoleColor = false
	// PrintLog default is false. If it is true, all logs will be printed in the console.
	PrintLog = false

	// WriteLog default is false. If it is true, all logs will be stored in the logfiles.
	WriteLog = false

	// LogDepth default is 8. It is used for callerHook.
	LogDepth = 8

	// WithCallerHook default is true. If it is true, log will print the log with file, line.
	WithCallerHook = true

	// LogLevel default is debug. If LogLevel is correct set, the log level will be use the value LogLevel, otherwise it will use the default LogLevel
	LogLevel = "debug"
)
