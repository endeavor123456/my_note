package main

import "fmt"

//父类
type person7 struct {
	id   int
	name string
	age  int
	sex  string
}

//子类
type student7 struct {
	person7
	class int
	score int
	addr  string
}

func (p *person7) Print() {
	fmt.Printf("编号：%d\n", p.id) //(*p).id的简写
	fmt.Printf("姓名：%s\n", p.name)
	fmt.Printf("年龄：%d\n", p.age)
	fmt.Printf("性别：%s\n", p.sex)
}

func main01() {
	// p := person7{1001, "孙尚香", 32, "女"}
	// p.Print()
	s := student7{person7{1002, "糜夫人", 32, "女"}, 302, 100, "巴蜀"}
	s.Print() //子类也可以继承父类的方法  注意输出结果 多看看
}
