package main

import "fmt"

/*
如果在函数执行中遇到return，return后面的代码不会再被执行 函数会提前执行结束 如果函数执行中没遇到return则函数会在执行完之后再结束
*/
func test5(a int, b int) int {
	sum := a + b //这个地方的a和b已经在上面定义了a int, b int
	return sum   //return sum 表示返回变量sum里的值，此时变量sum里的值就是返回值 return表示函数的结束 如果在函数中return之后还有代码则不会执行
}

/*
与上面等价
func test5(a int, b int) (sum int) {
	sum = a + b //这个地方的a和b已经在上面定义了a int, b int
	return   //return表示函数的结束 如果在函数中return之后还有代码则不会执行
}
*/

func main05() {
	a := 10
	b := 20
	sum := test5(a, b)
	fmt.Println(sum)
}

//函数返回多个返回值
func test6(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

/*
与上面等价
func test6(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum,sub
}
*/
func main06() {
	a := 10
	b := 20
	sum, sub := test6(a, b)
	fmt.Println(sum, sub)
	sum, sub = test6(a, b) //如果两个返回值都不接收不能写成_,_ := test6(a, b)写成_,_ = test6(a, b)可以，不想接受返回值直接写test6(a, b)即可
	fmt.Println(sub)
}
