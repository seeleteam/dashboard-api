/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Ping default for test
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Ping() time now %v", time.Now())
		c.JSON(200, gin.H{
			"message": "ping" + c.Request.URL.Path,
		})
	}
}

// Pong default for test
func Pong() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Pong() time now %v", time.Now())
		c.JSON(204, gin.H{
			"message": "pong" + c.Request.URL.Path,
		})
	}
}

// Kong default for test
func Kong() gin.HandlerFunc {
	return func(c *gin.Context) {
		code, err := c.Writer.WriteString("Kong" + c.Request.URL.Path)
		if err != nil {
			log.Error("error is %v %v", code, err)
			log.Errorln("error is ", code, err)

		} else {
			log.Info("info is %v and time now %v", code, time.Now())
			log.Infoln("info is ", code, time.Now().Nanosecond())
		}
	}
}

// LongAsync async task with goroutine!
func LongAsync() gin.HandlerFunc {
	return func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Printf("Done! in path %v", cCp.Request.URL.Path)
		}()
	}
}
