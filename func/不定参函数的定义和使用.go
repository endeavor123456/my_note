package main

import "fmt"

//虽然此处arr是个切片类型的变量 注意：不定参函数不是形参为切片类型的函数
func sum(arr ...int) {
	fmt.Println(arr)
}

//如果不定参函数调用时传递的参数为多个 不定参要写在其他参数后面
func test2(a int, arr ...int) {
	fmt.Println(a)
	fmt.Println(arr)
}
func main02() {
	sum()
	fmt.Println("------------------------")
	test2(1, 2, 3, 4, 5)
}

/*
形参为切片类型的函数
func add(s []int) {
	fmt.Println(s)
}
//不定参函数
func test3(arr ...int) {
	test4(arr[0], arr[1], arr[2], arr[3]) //上下两句等价 这一句是手动打散切片
	test4(arr[0:4]...)                    //左包右不包    这个是使用...运算符自动打散
}
func main04() {
	test3(1, 2, 3, 4)
	test := []int{1, 2, 3, 4}
	add(test)
}
*/
