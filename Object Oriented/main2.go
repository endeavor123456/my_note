package main

import "fmt"

type per struct {
	id   int
	name string
	age  int
	sex  string
}
type stu struct {
	*per
	name  string
	class int
	score int
}

func main03() {
	var stu stu
	stu.per = new(per)
	stu.id = 1          //(*stu.per).id = 1的简写 步骤：(*stu.per).id = 1 => stu.per.id=1 => stu.id = 1
	stu.per.name = "盖伦" //(*stu.per).name="盖伦"的简写 记住.的优先级比*高
	stu.age = 30
	stu.sex = "男"
	stu.class = 301
	stu.score = 90
	fmt.Println(stu.name)
	fmt.Println(stu.id)
	fmt.Println(stu.age)
	fmt.Println(stu.sex)
	fmt.Println(stu.class)
	fmt.Println(stu.score)
	//还有另一种写法 可以不用new函数	stu.per = &per{1001, "盖伦", 30, "男"}
}
