package main

import "fmt"

var a int

type res struct {
	age  int
	name string
	sex  string
}

func getName() string {
	name := "狗子2号"
	return name
}
func returnStruct() (returnStruct res) {
	return res{20, "狗子3号", "男"}
}
func main() {
	var a int = 20
	r := res{
		age:  a,         //把a的值赋到字段age
		name: getName(), //把getName函数的返回值赋值给字段name
		sex:  "男",
	}
	res1 := res{a, getName(), "男"} //分别把a的值，getName函数的返回值和"男"赋值给对应的字段【即：age name sex】
	fmt.Println(r)
	fmt.Println(res1)
	fmt.Println("-----------------------------------------------")
	fmt.Println(returnStruct())
	fmt.Println(returnStruct().name) //表示
}
