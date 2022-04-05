package main

import "fmt"

//定义一个函数 可以传入任意类型的数据 然后根据不同的数据类型实现不同的功能  第一种实现方式
func MyPrint(x interface{}) { //interface{}类型的变量.(Type)
	if _, ok := x.(string); ok {
		fmt.Println("我是string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("我是int类型")
	} else if _, ok := x.(bool); ok {
		fmt.Println("我是bool类型")
	}
}

// 定义一个函数 可以传入任意类型的数据 然后根据不同的数据类型实现不同的功能  第二种实现方式
//interface{}类型的变量.(type)只能在Switch中使用    注意：变量.(type)中的type不是Type 别搞混了
func MyPrint2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("我还是int类型")
	case string:
		fmt.Println("我还是string")
	case bool:
		fmt.Println("我是bool类型")
	default:
		fmt.Println("错误")
	}
}

//类型断言
func main() {
	var a interface{}
	a = "你好，golang"
	v, ok := a.(string) //变量ok是一个bool类型的变量
	if ok {
		fmt.Println("a就是一个string类型，值是：", v)
	} else {
		fmt.Println("类型断言失败")
	}

	MyPrint("www")
	MyPrint2(3)
}
