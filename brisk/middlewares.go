package brisk

import (
	"fmt"
	"net/http"
	"time"
)

type LoggingMiddleware struct{}
type CrosMiddleware struct{}

func (l *LoggingMiddleware) process_request(c *Context) {
	method := c.Method
	path := c.Path
	time := time.Now()
	timestr := time.Format("2006-01-02 15:04")
	useragent := c.Request.UserAgent()
	loggingstr := fmt.Sprintf("*Request:\t【method:%s; path:%s】\t %s\t from：%s\t", method, path, timestr, useragent)
	println(loggingstr)
}

func (l *LoggingMiddleware) process_response(c *Context) {
	status := c.StatusCode
	path := c.Path
	time := time.Now()
	timestr := time.Format("2006-01-02 15:04")
	loggingstr := fmt.Sprintf("*Response:\t【status:%d; path:%s】\t %s\t", status, path, timestr)
	println(loggingstr)
}

// CROS
func (cros *CrosMiddleware) process_response(c *Context) {
	if c.Method == http.MethodOptions {
		c.WriteString(http.StatusNoContent, "")
	}
}

func (cros *CrosMiddleware) process_request(c *Context) {
	c.AddHeader("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	c.SetHeader("Access-Control-Allow-Origin", "*")
	c.AddHeader("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.AddHeader("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.SetHeader("Access-Control-Allow-Credentials", "true")

	if c.Method == "OPTIONS" {
		c.SetStatusCode(http.StatusNoContent)
	}
}
