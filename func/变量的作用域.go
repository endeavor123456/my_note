package main

import "fmt"

/*
1)函数内部声明的变量叫局部变量 作用域仅限于函数内部 无论首字母大小写都只能在函数内部使用
2)函数外部声明的变量叫全局变量首字母小写作用域在当前包有效，如果 其首字母为大写，则作用域在整个项目有效
3)如果变量是在一个代码块，比如for/if，那么这个变量的作用域就在该代码块
4)变量作用域遵循就近原则
*/
var age int = 50         //只在当前包有效
var Name string = "jack" //在整个项目有效	var Name string = "jack"为变量初始化
var height int

func people() { //函数内部声明的变量 age,Name的作用域只在函数people()内部
	age := 10
	Name := "tom"
	fmt.Println(age)
	fmt.Println(Name)
}
func people1(age int, Name string) { //函数内部声明的变量 age,Name的作用域只在函数people1()内部
	fmt.Println(age)
	fmt.Println(Name)
}

//除了函数名，这个函数的里的变量都是函数内部定义的变量
func main11() {
	fmt.Println("ok")
	fmt.Println(age)
	fmt.Println(Name)
	fmt.Println(height)
	fmt.Println("-----------------------------------------")
	/*
		for i := 0; i <= 10; i++ {
			fmt.Println("i=", i)
		}
		fmt.Println("i=", i) //这个地方会报错 因为变量i没有被声明 因为for循环里面的变量不能被外部使用
	*/
	var i int //变量i可以在for语句内部使用 因为变量i是局部变量 可以在当前的函数的内部使用
	for i = 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
	/*
		var i int = 10
		for i := 0;i<= 10;i++{  当for语句需要使用某个变量时 如果for语句内部有这个变量
			                    就直接使用for语句内部的 如果for语句没有定义这个变量就使用for语句外面的 就近原则
			fmt.Println(i)
		}
		fmt.Println(i)
	*/
	//if等语句就不测试了
}
