package brisk

type HandleFunc func(c *Context)

type MiddleWare interface {
	process_request(c *Context)
	process_response(c *Context)
}
