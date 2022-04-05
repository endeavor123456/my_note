package main

import (
	"fmt"
)

//方法是特殊的函数 方法名和函数名都是函数类型的变量 该变量保存的都是一个地址
type student9 struct {
	name string
	age  int
	sex  string
	addr string
}

func test() {
	fmt.Println("你好我是test()")
}
func (s *student9) PrintInfo() { //PrintInf的类型为func()
	fmt.Println(*s)
}
func main10() {
	stu := student9{"甄姬", 32, "女", "许昌"}
	f := stu.PrintInfo  //PrintInf和test的类型是相同的
	f1 := test          //变量相互赋值必须二者的变量类型相同 f1 := test等价于var f1 func() f1=test
	fmt.Printf("%T", f) //输出为func()
	fmt.Println()
	fmt.Printf("%T", f1) //输出为func()
	fmt.Println()
	f()  //调用名为f的方法
	f1() //调用名为f1的函数
}
