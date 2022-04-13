<p align="center">
  <a href="" rel="noopener">
 <img width=250px height=100px src="docs/Brisk.png" alt="Project logo"></a>
</p>

<h4 align="center">高性能go语言web框架</h4>
<br>
<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![Go Reference](https://pkg.go.dev/badge/github.com/DomineCore/brisk.svg)](https://pkg.go.dev/github.com/DomineCore/brisk)

[![GitHub Issues](https://img.shields.io/github/issues/DomineCore/brisk.svg)](https://github.com/DomineCore/brisk/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/DomineCore/brisk.svg)](https://github.com/DomineCore/brisk/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)
</div>

# Brisk 高性能HTTP-web框架 
> Brisk，使用go标准库net/http构建的HTTP-web框架，采用前缀树路由系统，支持动态路由匹配。

<br>
Brisk可以用来构建一些简单的api，目前已支持动态路由、自定义中间件、分组路由；后续将会推出一系列内置中间件，支持项目级别的配置管理、集成go模板渲染库、信号与接收器、热更新等更加丰富的功能便于开发者使用。

<br>


## Getting Started
---
### 在项目中引入
```go
// main.go

package main

import "github.com/DomineCore/brisk"

```
在项目中引入后，请手动执行 `go mod tidy`

### 创建brisk实例 && 运行brisk应用
```go
// main.go

package main

import "github.com/DomineCore/brisk"

func main() {
  b := brisk.New()
  b.Run(":8000")
}

```

### 添加你的第一个API：Hello Brisk!
```go
// main.go

func main() {
  b :=brisk.New()
  b.Get("/", func (c *brisk.Context) {
    c.WriteString(http.StatusOk, "Hello Brisk!")
  })
  b.Run(":8000")
}

```
brisk提供了快捷的创建Get、Post请求的两个方法，这两个方法直接绑定在Brisk结构体上，便于开发者创建简单api。
```go
func (b *Brisk) Get(pattern string, handler HandleFunc){}
func (b *Brisk) Post(pattern string, handler HandleFunc){}
```
### 使用路由
Brisk提供了一个基于动态前缀树的路由结构体。使用该结构体可以实现添加API、分组路由、动态路径路由等功能。
<br>

#### 使用路由添加API
在Brisk结构体中包含一个Router，它将作为整个项目的总路由，下面我们演示用Brisk.Router来添加路由的方法
```go
func main() {
  ···
  b.Router.Add("/api/", http.MethodPost, func(c *brisk.Context){
    c.WriteString(http.StatusOk, "Hello Brisk!")
  })
  ···
}
```

#### 分组路由
Brisk.Router结构体提供了一个Include方法，用于连接子路由，借助Include方法可以方便地对我们的项目结构按照不同的路有前缀进行切分。借助分组路由，我们将获得更好的项目组织能力。
```go
func main() {
  ···
  // 创建一个子路由
  api_v1 := brisk.NewRouter()
  api_v1.Add("/hello/", http.MethodGet, func(c *brisk.Context){
    c.WriteString(http.StatusOk, "Hello Brisk!")
  })
  // 连接到主路由
  b.Router.Include("/api/v1/", api_v1)
  ···
}
```

#### 动态路径参数
Brisk.Router也支持开发者使用形如`api/v1/:id`的动态路径，例如：

`/api/v1/123/` 被解析后将会在上下文Context中新增一个PathParams参数id=123


### 使用中间件

中间件是用于处理api的公共逻辑，如跨域、日志、登录认证和鉴权等功能，通过中间件来实现是更好的选择。

Brisk内置一个Middleware接口，接口包含两个方法process_request和process_response，两个方法都将Context请求上下文作为参数，借助Context的能力来实现对请求和响应的公共处理逻辑。

#### 使用内置中间件
Brisk内置了两个中间件：LoggingMiddleware和CrosMiddleware分别用于打印访问日志、处理跨域。
```go
b.Router.Use(&brisk.LoggingMiddleware{}) // 使用Router.Use方法来应用中间件
b.Router.Use(&brisk.CrosMiddleware{})
```

#### 自定义中间件
Brisk实现的Middleware接口，是实现自定义中间件的标准。

实现自定义中间件我们需要创建一个自己的中间件结构体，并且在结构体上实现process_request和process_response两个方法用来实现具体逻辑，以访问日志中间件为例:
```go
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
```

### 管理配置
brisk 使用go知名开源库viper来管理配置，viper被注入在Birsk结构体中，这样我们可以通过app对象来访问配置。

#### 声明配置文件路径
```go
b.Conf.SetConfigFile('./config') // 配置文件所在路径
b.Conf.SetConfigName('settings') // 配置文件名
b.Conf.SetConfigType('json') // 配置文件类型
```


#### 获取配置项
```go
b.Conf.Get("key")
b.Conf.GetString("key") //返回key对应value的string
b.Conf.GetBool("key") //返回对应bool值
```

更多能力请访问viper<br>

<img src="https://pkg.go.dev/badge/mod/github.com/spf13/viper" alt="PkgGoDev">

----

Brisk仍在快速迭代中，敬请期待后续版本的优化👾。