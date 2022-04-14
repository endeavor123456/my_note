package main

import "fmt"

type Persons struct {
	age  int
	name string
}

func demo() Persons {
	return Persons{10, "二狗"}
}
func main() {
	fmt.Println(demo().age) //这句话的执行步骤是 第一步先调用函数demo，函数执行结束后返回返回值 第二步拿出返回值中的保存在字段age中的值 最后一步输出这个值
	test := demo()          //是var test Persons  test = demo()  先定义变量test 然后调用函数 函数执行结束后返回返回值 最后将返回值赋值给变量test
	fmt.Println(test)       //输出变量test里的数据
	fmt.Println("------------------------------------")
	var person = Persons{5, "胡图图"}
	fmt.Println(person)     //输出此变量里的所有数据
	fmt.Println(person.age) //执行步骤：第一步：从变量中拿出字段age 第二步：输出字段age里的所有的值
	fmt.Println("---------------------------------------------")
	fmt.Println(&person.name) //执行步骤：首先要注意.的优先级比&的优先级高 第一步：先拿出字段name  第二步：对字段name取地址 第三步：输出字段name的地址
	fmt.Printf("%p", &person)
}
