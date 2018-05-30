package log

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 2016-09-27 09:38:21.541541811 +0200 CEST
// 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700]
// "GET /apache_pb.gif HTTP/1.0" 200 2326
// "http://www.example.com/start.html"
// "Mozilla/4.08 [en] (Win98; I ;Nav)"

// var timeFormat = "02/Jan/2006:15:04:05 -0700"
var timeFormat = "20060102150405 -0700"
var timeFormat2 = "2006/01/02 - 15:04:05"

// Logger is the logrus logger handler
func Logger(globalLog *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latencyOrigin := time.Duration(int(math.Ceil(float64(stop.Nanoseconds()))))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknow"
		}
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		logrusFields := map[string]interface{}{
			"timestamp":  int(start.UnixNano() / 1000000.0),
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latencyOrigin, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		}

		_ = logrusFields

		entry := logrus.NewEntry(globalLog)

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s |%d |%s |%s \t|%s \t %s|%s |%d \"%s\" \"%s\"", time.Now().Format(timeFormat2), statusCode,
				clientIP, latencyOrigin, c.Request.Method, path, hostname, dataLength, referer, clientUserAgent)

			if statusCode > 499 {
				// you can output the format fields for debug, open the comment following code
				// entry.WithFields(logrusFields).Error(msg)

				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
