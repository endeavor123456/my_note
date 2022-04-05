package main

import "fmt"

//接口是一种类型 可以叫他接口、也可以叫他接口类型
type Producter interface {
	kaishi()
	zanting()
}
type Tv struct {
	Name string
}
type Computer struct {
}

func (p Tv) kaishi() {
	fmt.Println(p.Name, "启动")
}
func (p Tv) zanting() {
	fmt.Println(p.Name)
}
func (c Computer) work(pro Producter) {
	pro.kaishi() //调用接口中的方法
	pro.zanting()
}
func main12() {
	var computer = Computer{}
	var tv = Tv{
		Name: "小米电视",
	}
	computer.work(tv)
}
