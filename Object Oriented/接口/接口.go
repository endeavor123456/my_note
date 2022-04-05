package main

import "fmt"

//先定义接口 再根据接口去定义方法
//接口是一种类型
// 接口是一组方法的集合(0-n个方法的集合) 是一个规范 它规定必须有哪些名字的方法、有无形参若是由的话得是什么类型的、有无返回值若是有的话得是什么类型
// 接口中所有的方法都没有方法体 接口定义了一个对象的行为规范 该对象绑定的方法得按照接口里的方法格式来 只定义规范不实现。接口体现了程序设计的多态和高内聚低耦合的思想
//如果一种类型实现了某个接口，那么这种类型的对象能够赋值给这个接口类型的变量。

/*
type 接口名 interface{
	方法名1(形参变量类型)返回值类型
	方法名2(形参变量类型)返回值类型
	......
}
*/
type Usber interface { //定义一个接口
	start()
	stop()
}
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name)
}
func main11() {
	p := Phone{Name: "手机"} //变量p是Phone类的一个对象
	var p1 Usber
	p1 = p     //这里没报错说明Phone这个类型实现了这个接口
	p1.start() //调用接口中的方法
	// fmt.Println(p1)

}
