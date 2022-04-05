package main

import "fmt"

//空接口   interface{}就是空接口 因为空接口没有任何方法所以任意的类型都可以实现空接口
func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)
}
func main13() {
	var aa interface{}
	var str = "你好，golang"
	aa = str //让字符串类型实现空接口
	fmt.Println(aa)

	var a interface{} //interface{}类型的变量可以保存任意类型的数据   后面会慢慢理解
	a = 20
	fmt.Printf("值：%v 类型:%T\n", a, a)
	a = "你好，golang"
	fmt.Printf("值：%v 类型：%T\n", a, a)
	a = true
	fmt.Printf("值:%v 类型:%T\n", a, a)
	show(20)
	show("你好golang")
	slice := []int{1, 2, 3, 4}
	show(slice)
	var s1 = []interface{}{1, 2, "你好", true}
	fmt.Println(s1)
	var m1 = make(map[string]interface{})
	m1["username"] = "张三"
	m1["age"] = 20
	m1["married"] = true
	fmt.Println(m1)
}

/*
var a interface{} //interface{}类型的变量 可以理解成js里的变量
	a = "string"
	fmt.Printf("%T\n", a) //输出为string
	a = 10
	fmt.Printf("%T\n", a) //输出为int
	a = 1.25
	fmt.Printf("%T\n", a) //输出为float64
	var b int32
	a = b
	fmt.Printf("%T\n", a) //输出为int32
*/
