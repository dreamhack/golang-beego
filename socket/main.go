package main

import (
	"fmt"
	"math/rand"
)

//一个查询结构体
type query struct {
	//参数Channel
	sql chan string
	//结果Channel
	result chan string
}

//执行Query
func execQuery(q query) {
	//启动协程
	go func() {
		//获取输入
		sql := <-q.sql
		//访问数据库，输出结果通道
		q.result <- "get " + sql
	}()

}

func main() {
	//初始化Query
	q := query{make(chan string, 1), make(chan string, 1)}
	//执行Query，注意执行的时候无需准备参数
	execQuery(q)

	//准备参数
	q.sql <- "select * from table"
	//获取结果
	fmt.Println(<-q.result)
}
