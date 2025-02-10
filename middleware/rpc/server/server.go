package main

import (
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Res int
}

type Cal int

// CalSquare 定义服务提供方法
func (cal *Cal) CalSquare(num int, result *Result) error {
	result.Num = num
	result.Res = num * num

	return nil
}

func main() {
	//发布 Cal 中满足 RPC 注册条件的方法
	rpc.Register(new(Cal))
	//注册用于处理 RPC 消息的 HTTP Handler
	rpc.HandleHTTP()

	http.ListenAndServe(":1234", nil)

}
