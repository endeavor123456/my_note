package main

import "fmt"

//阶乘实现
var s int = 1 //这个地方必须赋初值为1 因为默认值是0 这样结果就成0了

func demo2(n int) {
	if n == 1 {
		return
	}
	s *= n
	demo2(n - 1)
}
func main() {
	demo2(5)
	fmt.Println(s)
}
