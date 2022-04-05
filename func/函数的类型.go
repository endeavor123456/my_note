package main

import (
	"fmt"
)

//函数也是一种数据类型  以下面的函数为例 其实函数名demo是func(int,int)类型的变量 该变量保存的是此函数在代码区的地址
func demo(a int, b int) {
	fmt.Println(a + b)
}
func main07() {
	/*
		//函数的名字是一个地址 函数在代码区的地址(对于这个地方先不深究)
		fmt.Print("%p", demo)
		fmt.Printf("%T", demo)
	*/
	f := demo //这个地方等价于 var f func(int,int)   f=demo
	//对于函数名相互赋值的前提是二者的变量类型相同 变量f和变量demo都是func(int,int)类型的 所以可以相互赋值
	f(10, 20)
	fmt.Printf("%T", f)

}
