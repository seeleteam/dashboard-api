/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package log

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/seeleteam/dashboard-api/common"
)

const (
	defaultLogDir      = "dashboard-api-logs"
	defaultLogFilePrex = ""
	defaultLogFile     = "all.logs"
)

// GlobalLog struct
type GlobalLog struct {
	log *logrus.Logger
}

// GetLogger convert GlobalLog to *logrus.Logger
func (m *GlobalLog) GetLogger() *logrus.Logger {
	return m.log
}

var logMap map[string]*GlobalLog
var getLogMutex sync.Mutex

// Debug wrapper Debug logger
func (m *GlobalLog) Debug(f interface{}, args ...interface{}) {
	m.log.Debug(formatLog(f, args...))
}

// Info wrapper Info logger
func (m *GlobalLog) Info(f interface{}, args ...interface{}) {
	m.log.Info(formatLog(f, args...))
}

// Warn wrapper Warn logger
func (m *GlobalLog) Warn(f interface{}, args ...interface{}) {
	m.log.Warn(formatLog(f, args...))
}

// Printf wrapper Printf logger
func (m *GlobalLog) Printf(f interface{}, args ...interface{}) {
	m.log.Print(formatLog(f, args...))
}

// Panic wrapper Panic logger
func (m *GlobalLog) Panic(f interface{}, args ...interface{}) {
	m.log.Panic(formatLog(f, args...))
}

// Fatal wrapper Fatal logger
func (m *GlobalLog) Fatal(f interface{}, args ...interface{}) {
	m.log.Fatal(formatLog(f, args...))
}

// Error wrapper Error logger
func (m *GlobalLog) Error(f interface{}, args ...interface{}) {
	m.log.Error(formatLog(f, args...))
}

// Debugln wrapper Debugln logger
func (m *GlobalLog) Debugln(v ...interface{}) {
	m.log.Debug(fmt.Sprintln(v...))
}

// Infoln wrapper Infoln logger
func (m *GlobalLog) Infoln(args ...interface{}) {
	m.log.Info(fmt.Sprintln(args...))
}

// Warnln wrapper Warnln logger
func (m *GlobalLog) Warnln(args ...interface{}) {
	m.log.Warn(fmt.Sprintln(args...))
}

// Printfln wrapper Printfln logger
func (m *GlobalLog) Printfln(args ...interface{}) {
	m.log.Print(fmt.Sprintln(args...))
}

// Panicln wrapper Panicln logger
func (m *GlobalLog) Panicln(args ...interface{}) {
	m.log.Panic(fmt.Sprintln(args...))
}

// Fatalln wrapper Fatalln logger
func (m *GlobalLog) Fatalln(args ...interface{}) {
	m.log.Fatal(fmt.Sprintln(args...))
}

// Errorln wrapper Errorln logger
func (m *GlobalLog) Errorln(args ...interface{}) {
	m.log.Error(fmt.Sprintln(args...))
}

// Options config the logger
type Options struct {
	PrintLog       *bool
	WriteLog       *bool
	LogLevel       *logrus.Level
	Depth          *int
	WithCallerHook *bool
}

// GetLoggerWithOptions with options config
func GetLoggerWithOptions(logName string, options *Options) *GlobalLog {
	getLogMutex.Lock()
	defer getLogMutex.Unlock()
	if logMap == nil {
		logMap = make(map[string]*GlobalLog)
	}
	curLog, ok := logMap[logName]
	if ok {
		return curLog
	}

	log := logrus.New()

	// get logLevel
	logLevel := common.GetLogLevel(common.LogLevel)

	// logger config set
	printLog := common.PrintLog
	writeLog := common.WriteLog
	depth := common.LogDepth
	withCallerHook := common.WithCallerHook

	// options set
	cpOptions := *options
	if &cpOptions != nil {
		if cpOptions.PrintLog != nil {
			printLog = *cpOptions.PrintLog
		}
		if cpOptions.WriteLog != nil {
			writeLog = *cpOptions.WriteLog
		}
		if cpOptions.LogLevel != nil {
			logLevel = *cpOptions.LogLevel
		}
		if cpOptions.Depth != nil {
			depth = *cpOptions.Depth
		}
		if cpOptions.WithCallerHook != nil {
			withCallerHook = *cpOptions.WithCallerHook
		}
	}

	log.SetLevel(logLevel)

	if writeLog {
		err := os.MkdirAll(defaultLogDir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("creating log file failed: %s", err.Error()))
		}

		path := defaultLogDir + string(os.PathSeparator) + defaultLogFile
		writer, err := rotatelogs.New(
			path+".%Y%m%d%H%M%S",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(time.Duration(60*60*24*7)*time.Second),  // 24 hours
			rotatelogs.WithRotationTime(time.Duration(86400)*time.Second), // 1 days
		)
		if err != nil {
			panic(fmt.Sprintf("rotatelogs log failed: %s", err.Error()))
		}

		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: writer,
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.FatalLevel: writer,
			},
			&logrus.TextFormatter{},
		))

		pathMap := lfshook.PathMap{
			logrus.DebugLevel: fmt.Sprintf("%s/%sdebug.log", defaultLogDir, defaultLogFilePrex),
			logrus.InfoLevel:  fmt.Sprintf("%s/%sinfo.log", defaultLogDir, defaultLogFilePrex),
			logrus.WarnLevel:  fmt.Sprintf("%s/%swarn.log", defaultLogDir, defaultLogFilePrex),
			logrus.ErrorLevel: fmt.Sprintf("%s/%serror.log", defaultLogDir, defaultLogFilePrex),
			logrus.FatalLevel: fmt.Sprintf("%s/%sfatal.log", defaultLogDir, defaultLogFilePrex),
		}
		log.AddHook(lfshook.NewHook(
			pathMap,
			&logrus.TextFormatter{},
		))
	} else {
		if printLog {
			log.Out = os.Stdout
		}
	}

	if withCallerHook {
		log.AddHook(&CallerHook{depth: depth}) // add caller hook to print caller's file and line number
	}
	curLog = &GlobalLog{
		log: log,
	}
	logMap[logName] = curLog
	return curLog
}

// GetLogger gets logrus.Logger object according to logName
// each module can have its own logger
func GetLogger(logName string, printConsole bool) *GlobalLog {
	// writeFile flag use global config
	return GetLoggerWithOptions(logName, &Options{PrintLog: &printConsole})
}

// GetLoggerWithCaller gets logrus.Logger object according to logName
// with paramters printConsole and withCaller
// Used for caller control
func GetLoggerWithCaller(logName string, printConsole bool, withCaller bool) *GlobalLog {
	// writeFile flag use global config
	return GetLoggerWithOptions(logName, &Options{PrintLog: &printConsole, WithCallerHook: &withCaller})
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
