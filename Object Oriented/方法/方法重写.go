package main

import (
	"fmt"
)

//多练 多练 多练 多练 ！！！！！
//如果子类和父类有相同的方法名 叫方法重写 那么该方法名被称为重写方法名
//父类
type person8 struct {
	id   int
	name string
	age  int
	sex  string
}

//子类
type student8 struct {
	person8
	class int
	score int
	addr  string
}

//父类的方法
func (p *person8) SayHello() {
	fmt.Printf("大家好，我是：%s，我今年%d岁，我是%s生\n", p.name, p.age, p.sex)
}

//子类的方法
func (s *student8) SayHello() {
	fmt.Printf("大家好，我是：%s，我今年%d岁，我是%s生，我的班级是%d，我的成绩是%d\n", s.name, s.age, s.sex, s.class, s.score)
}
func main09() {
	//方法重写 只需要定义子类(student8)类型的变量 父类的方法调用格式：该变量.父类.重写方法名
	stu := student8{person8{1001, "甘夫人", 32, "女"}, 112, 80, "巴蜀"}
	//调用子类的方法
	stu.SayHello()
	//调用父类的方法
	stu.person8.SayHello()
}
