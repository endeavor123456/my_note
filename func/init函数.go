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
init函数 每个源文件中可以包含多个init函数
一、如果一个文件同时包含全局变量定义语句、init函数和main函数，则执行的流程是：先执行全局变量定义语句再执行init函数调用最后main函数调用【init函数、main函数是系统调用】调用函数就是执行函数
init 函数通常被用来：
	1)对全局变量进行初始化
	2)检查/修复程序的状态
	3)注册
	4)运行一次计算
二、Go要求非常严格，不允许引用不使用的包。但是有时你引用包只是为了调用init函数去做一些初始化工作。此时空标识符（也就是下划线）的作用就是为了解决这个问题
*/
func main() {
	fmt.Println("main()")

}
