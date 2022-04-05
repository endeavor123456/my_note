package main

import "fmt"

func test4(arr ...int) {
	fmt.Println(arr)
}

func test3(arr ...int) {
	test4(arr[0], arr[1], arr[2], arr[3]) //上下两句等价 这一句是手动打散切片
	test4(arr[0:4]...)                    //左包右不包 这个是使用...运算符自动打散
}
func main04() {
	test3(1, 2, 3, 4)
}

/*
func test1(arr ...float64) {
	fmt.Println(arr)
}
func test(arr ...float64) {
	// fmt.Println(arr)
	test1(arr[1])
}
func main() {
	test(1, 2, 3, 4)
}
*/
