package main

import "fmt"

func test1(a int, b int) {
	fmt.Println(a + b)
}
func test(a, b int) {
	test1(a, b)
}
func main03() {
	a := 10
	b := 20
	test(a, b)
}
