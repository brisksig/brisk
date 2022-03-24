package gee

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	// 原始结构封装
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	// 请求信息
	Path   string
	Params map[string]string
	Method string
	// 响应信息
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:        r,
		ResponseWriter: w,
		Path:           r.URL.Path,
		Method:         r.Method,
	}
}

//Get Request Data Methods

func (c *Context) PathParams(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) QueryParams(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Form(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) JsonBind(obj interface{}) error {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, obj)
}

// Set Response Info

func (c *Context) SetStatusCode(code int) {
	c.ResponseWriter.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.ResponseWriter.Header().Set(key, value)
}

// Write Response

func (c *Context) WriteJSON(code int, obj interface{}) {
	// set code
	c.SetStatusCode(code)
	// set Header
	c.SetHeader("Content-type", "application/json")
	// json
	encoder := json.NewEncoder(c.ResponseWriter)
	if err := encoder.Encode(obj); err != nil {
		// internal server error
		http.Error(c.ResponseWriter, err.Error(), 500)
	}
}

func (c *Context) WriteString(code int, format string, values ...interface{}) {
	c.SetStatusCode(code)
	c.SetHeader("Content-type", "text/plain")
	c.ResponseWriter.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) HTML(code int, html string) {
	c.SetStatusCode(code)
	c.SetHeader("Content-type", "text/html")
	c.ResponseWriter.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	// set unknow content-type response
	c.SetStatusCode(code)
	c.ResponseWriter.Write(data)
}
