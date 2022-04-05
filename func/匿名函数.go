package main

import "fmt"

func main2() {
	fmt.Println("ok")
	f := func(a int, b int) { //等价于var f func(int,int)  f = func(a int,b int)
		fmt.Println(a + b)
	}
	/*
		匿名函数是特殊的函数 特殊在......
		f := func(a int, b int) {
			fmt.Println(a + b)
		}
		匿名函数就是没有函数名的函数 这段代码就是给匿名函数起了个函数名
		这一段代码可以理解为func f(a int,b int){
			fmt.Println(a+b)
		}  但是二者不等价
	*/
	fmt.Println(f)
	f(1, 5)
	fmt.Println("-------------------------------------------")
	var m func(int, int) //声明一个func(int,int)类型的变量m
	m = f
	m(1, 2)

	fmt.Println("------------------------------")
	c := 10
	d := 20
	//此处变量v保存的是此匿名函数的返回值 所以得是有参数和返回值才行
	v := func(c int, d int) int {
		return c + d
	}(c, d) //如果声明匿名函数时{}后面有()表示函数调用 后面的()用来传参
	/*
		与上面等价
		v := func(c int, d int) int {
			return c + d
		}(10, 20)
	*/
	fmt.Println(v)
	fmt.Println("--------------------------------")
	e := 100
	g := 200
	h := func() int {
		return e + g
	}
	e = 100
	g = 200
	i := h()
	fmt.Println(i)
}
