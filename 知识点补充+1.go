package main

import "fmt"

type demo struct {
	age  int
	name string
}

func (demo *demo) test() int {
	return demo.age
}
func (demo *demo) test1() string {
	return demo.name
}
func (demo *demo) test2() *demo {
	return demo
}
func (demo *demo) test3() demo {
	return *demo
}
func main() {
	var demo = demo{10, "胡图图"}
	res := demo.test()
	res1 := demo.test1()
	res2 := demo.test2()
	res3 := demo.test3()
	fmt.Println("res的值为", res, "\n", "res1的值为", res1, "\n", "res2的值为", res2, "\n", "res3的值为", res3)
}
