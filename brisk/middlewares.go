package brisk

import (
	"fmt"
	"time"
)

type LoggingMiddleware struct{}

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
