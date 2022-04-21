package main

import "fmt"

var num = numtest() //全局变量定义语句 先执行赋值号右边再执行赋值号左边

func numtest() int {
	fmt.Println("test()")
	return 90
}
func init() {
	fmt.Println("init()")
}
func dd() {
	fmt.Println("Hello")
}

/*
init函数
1)如果一个文件同时包含全局变量定义语句、init函数和main函数，则执行的流程是：先执行全局变量定义语句再执行init函数调用最后main函数调用【init函数、main函数是系统调用】调用函数就是执行函数
*/
func main() {
	fmt.Println("main()")

}
