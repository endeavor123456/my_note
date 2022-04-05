package main

import "fmt"

//递归函数 即函数自己调用自己
func demo1(n int) {
	if n == 0 {
		return
	}
	demo1(n - 1) //函数运行结束的时候会返回到调用它的地方 执行调用语句的下一行代码
	fmt.Println(n)
}
func main1() {
	demo1(4)
}

//最终输出 1 2 3 4
