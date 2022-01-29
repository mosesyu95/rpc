package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Arith 定义算数算法服务
type Arith int

func (a Arith) Calc(req Args, res *Answer) error {
	log.Printf("receive %+v \n", req)
	if req.B == 0 {
		return errors.New("除数不能为0！")
	}
	res.Sum = req.A + req.B
	return nil
}

// Args 定义参数
type Args struct {
	A, B int
}

// Answer 定义结果
type Answer struct {
	Sum int
}

func main() {
	var err error
	// 注册rpc服务
	err = rpc.Register(new(Arith))
	if err != nil {
		log.Fatalln(err)
	}
	// http为rpc载体
	rpc.HandleHTTP()
	// 设置监听
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// 开启http服务
	go http.Serve(lis, nil)
	log.Println("server is running ")
	select {}
}
