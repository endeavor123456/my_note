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
init函数 每个源文件中可以包含多个init函数 同一个源文件内有多个init()函数时，执行顺序从上到下 但是对于可读性和以后的可维护性来说建议只写一个
一、如果一个文件同时包含全局变量定义语句、init函数和main函数，则执行的流程是：先执行全局变量定义语句再执行init函数调用最后main函数调用【init函数、main函数是系统调用】调用函数就是执行函数
init 函数通常被用来：
	1)对全局变量进行初始化
	2)检查/修复程序的状态
	3)注册
	4)实现sync.Once功能
二、Go要求非常严格，不允许引用不使用的包。但是有时你引用包只是为了调用init函数去做一些初始化工作。此时空标识符（也就是下划线）的作用就是为了解决这个问题
 _ "github.com/go-sql-driver/mysql"
三、init函数只能被系统调用 也不可赋值给函数变量 否则会报错
四、导入包不要出现循环依赖，否则会导致报错
五、init不应该依赖任何在main函数里定义的变量，因为init函数的执行顺序是在main函数之前的
六、init函数中也可以启动goroutine，也就是在初始化的同时启动新的goroutine，这并不会影响初始化顺序
七、不要依赖init函数执行的顺序
八、无论包被导入多少次，init函数只能被调用一次，也就是执行一次
_ "github.com/go-sql-driver/mysql"和"github.com/go-sql-driver/mysql"mysql包被导入了两次 而只里面的init函数只会执行一次 init函数不会因为其所在的包被导入两次就执行两次
九、不同包的init函数按照包导入的顺序来执行
*/
func main() {
	fmt.Println("main()")

}
