package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.WithFields(logrus.Fields{
			"header":              c.Request.Header,
			"body":                c.Request.Body,
			"query":               c.Request.URL.Query(),
			"path":                c.Request.URL.Path,
			"ip":                  c.ClientIP(),
			"agent":               c.Request.UserAgent(),
			"request_id":          c.Request.Header.Get("X-Request-ID"),
			"request_time":        c.Request.Header.Get("X-Request-Time"),
			"request_size":        c.Request.ContentLength,
			"request_protocol":    c.Request.Proto,
			"request_remote_addr": c.Request.RemoteAddr,
			"request_host":        c.Request.Host,
			"realm":               "middleware/logrus",
		}).Infof("Request: %s %s", c.Request.Method, c.Request.URL.String())
		c.Next()
	}
}
