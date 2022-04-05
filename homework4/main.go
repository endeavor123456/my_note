package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		//变量stu的地址是不变的 变得是stu里的数据
		//循环体每执行一次就把上一次的stu里保存的数据覆盖掉 最后只剩下最后一条数据
		//此时对&stu进行取值操作 取出的值就是{name: "博客", age: 28} (*v).name就为博客
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	/*
				修改结果
		m := make(map[string]student)
			stus := []student{
				{name: "pprof.cn", age: 18},
				{name: "测试", age: 23},
				{name: "博客", age: 28},
			}

			for _, stu := range stus {
				m[stu.name] = stu
			}
			for k, v := range m {
				fmt.Println(k, "=>", v.name)
			}
	*/
}
