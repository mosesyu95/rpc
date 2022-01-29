package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Args 定义参数
type Args struct {
	A int
	B int
}

// Answer 定义结果
type Answer struct {
	Sum int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := Args{9, 2}
	var res Answer

	err = conn.Call("Arith.Calc", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d + %d = %d \n", req.A, req.B, res.Sum)
}