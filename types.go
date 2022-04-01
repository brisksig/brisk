// Copyright 2022 DomineCore.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package brisk

type HandleFunc func(c *Context)

type MiddleWare interface {
	process_request(c *Context)
	process_response(c *Context)
}
