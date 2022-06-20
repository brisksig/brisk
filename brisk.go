// Copyright 2022 DomineCore.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package brisk

import (
	"fmt"
	"net/http"
	"strings"
)

type Brisk struct {
	Router *Router
	// Conf   *Conf
}

func New(configpath string) *Brisk {
	// new app
	brisk := &Brisk{Router: NewRouter()}
	// load config
	Config = *NewConf()
	brisk.LoadConfig(configpath)
	// load db
	// if usedb is true, execute loadDatabase
	if Config.GetBool("usedb") {
		brisk.LoadDataBase()
	}
	return brisk
}

func (b *Brisk) LoadConfig(configpath string) {
	// loading viper config
	configpathslice := strings.Split(configpath, "/")
	file := configpathslice[len(configpathslice)-1]
	fileslice := strings.Split(file, ".")
	name := fileslice[0]
	filetype := fileslice[1]
	path := strings.TrimRight(configpath, file)
	SetConf(path, name, filetype)
}

func (b *Brisk) LoadDataBase() {
	// loading database orm instance
	// build brisk.DB (gorm.DB)
	connector := NewDBConnector()
	connector.Connect()
}

func (b *Brisk) Post(pattern string, handler HandleFunc) {
	b.Router.Add(pattern, http.MethodPost, handler)
}

func (b *Brisk) Get(pattern string, handler HandleFunc) {
	b.Router.Add(pattern, http.MethodGet, handler)
}

func (b *Brisk) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 1、上下文构建
	c := NewContext(w, req)
	// 2、路由分发
	method := c.Method
	pattern := c.Path
	// --- 前置中间件
	for _, middleware := range b.Router.Middlewares {
		middleware.process_request(c)
	}
	handler, err := b.Router.Dispatch(method, pattern, c)
	if err != nil {
		c.WriteString(http.StatusNotFound, "404 Not Found")
	} else {
		// --- handler process
		handler(c)
	}
	// --- 后置中间件
	for _, middleware := range b.Router.Middlewares {
		middleware.process_response(c)
	}
}

func (b *Brisk) Run(addr string) (err error) {
	// load Conf
	// b.LoadConfig()
	// load DataBase
	// b.LoadDataBase()
	// Listen
	fmt.Printf("server running····\n")
	if strings.HasPrefix(addr, ":") {
		fmt.Printf("listen on http://localhost%s\n", addr)
	} else {
		fmt.Printf("listen on http://%s\n", addr)
	}
	return http.ListenAndServe(addr, b)
}
