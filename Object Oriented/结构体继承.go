package main

import "fmt"

//golang中的继承就是结构体继承
//golang的类就是结构体   结构体继承和嵌套是有区别的
//继承即为一个结构体的没有名字的字段(没有名字的变量)的字段类型是另一个结构体的结构体名
type Person struct {
	id   int
	name string
	age  int
	sex  string
}
type Student struct {
	Person //匿名结构体类型的字段 继承
	class  int
	score  int
}
type person2 struct {
	name string
	age  int
	sex  string
}
type person3 struct {
	id   int
	addr string
}
type student2 struct {
	person2
	person3
	class int
	score int
}

func main04() {
	/*
		//匿名字段 实现继承  匿名字段 是没有名字的变量(不是指匿名变量)
		var stu1 Student
		stu1.name = "张三"
		stu1.age = 18
		stu1.sex = "女"
		stu1.class = 302
		stu1.score = 99
		fmt.Println(stu1)
		var stu Student = Student{Person{101, "东方不败", 35, "不详"}, 302, 100}
		//var stu1 Student = Student{Person:Person{101, "东方不败", 35, "不详"}, class:302, score:100} 这种写起来比较麻烦
		fmt.Println(stu)
	*/
	//结构体多重继承 即一个类有多个父类
	var Stu = student2{person2{"劫", 12, "男"}, person3{1002, "召唤师峡谷"}, 302, 100}
	fmt.Println(Stu)
}
