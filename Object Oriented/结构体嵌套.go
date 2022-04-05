package main

import "fmt"

//golang的类就是结构体   结构体继承和嵌套是有区别的
//结构体嵌套其实就是一个结构体的字段的类型是另一个结构体的结构体名或者是与另一个结构体名相关的类型 注意：结构体不能与他本身嵌套 即结构体不能有与该结构体名类型相关的类型的字段
type Person1 struct {
	id   int
	name string
	age  int
	sex  string
}
type Student1 struct {
	p Person1 //嵌套
	//Person1 匿名结构体类型的字段 继承
	class int
	score int
}

func main0() {
	//p Person 实现嵌套
	var stu Student1
	stu.p.name = "张三"
	//不可以使用 stu.name = "张三" 这种的方式
	stu.p.age = 18
	stu.p.sex = "女"
	stu.class = 302
	stu.score = 99
	fmt.Println(stu)
}
