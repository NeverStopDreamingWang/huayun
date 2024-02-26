package middleware

import (
	"github.com/NeverStopDreamingWang/goi"
	"huayun/huayun"
)

func init() {
	// 注册中间件
	// 注册请求中间件
	huayun.Server.MiddleWares.BeforeRequest(RequestMiddleWare)
	// 注册视图中间件
	// Server.MiddleWares.BeforeView()
	// 注册响应中间件
	// Server.MiddleWares.BeforeResponse()
}

// 请求中间件
func RequestMiddleWare(request *goi.Request) any {
	// fmt.Println("请求中间件", request.Object.URL)
	return nil
}

// 请求中间件
func ViewMiddleWare(request *goi.Request) any {
	// fmt.Println("请求中间件", request.Object.URL)
	return nil
}

// 请求中间件
func ResponseMiddleWare(request *goi.Request) any {
	// fmt.Println("请求中间件", request.Object.URL)
	return nil
}
