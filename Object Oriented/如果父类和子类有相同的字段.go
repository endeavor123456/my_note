package main

import "fmt"

type person struct {
	id   int
	name string
	age  int
	sex  string
}
type student struct {
	person
	name  string
	class int
	score int
}

func main02() {
	//子类和父类有相同的字段name 默认赋值给子类 可以通过stu.person.name = "李四"来给父类中的name字段赋值
	var stu student
	stu.score = 100
	stu.class = 301
	stu.id = 1002
	stu.name = "李四"
	fmt.Println(stu) //输出结果为{{1002 空串 0} 李四 301 100}
}
