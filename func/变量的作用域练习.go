package main

import "fmt"

var name = "tom"

func aa1() {
	fmt.Println(name) //如果函数内部用到某个变量 系统会先在当前函数内找 如果没有就去找全局变量
}
func aa2() {
	name := "jack"
	fmt.Println(name) //如果函数用到某个变量 系统会先在当前函数内找 如果没有当前函数内部没有声明此变量，就去找全局变量 就近原则
}

func main12() {
	fmt.Println("ok")
	fmt.Println(name)
	aa1()
	aa2()
	aa1()
}
