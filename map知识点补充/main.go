package main

import "fmt"

type Int = int

//map知识点补充
//map很多地方都挺怪
//map对长度不敏感
func main() {
	//map的key的类型不能是切片、map、function类型
	//map、切片类型的比较 详见golang中的比较注意事项.md
	var m = make(map[string]int) //该map的key是string类型的，value是int类型的 map中的key-value对是无序的
	fmt.Println(m)
	/*
		   课堂练习
		   存放3个学生的信息，每个学生有max和sex信息

		studentMap := make(map[string]map[string]string)
		studentMap["stu01"] = make(map[string]string)
		studentMap["stu01"]["name"] = "tom"
		studentMap["stu01"]["sex"] = "男"
		studentMap["stu01"]["address"] = "北京长安街"

		studentMap["stu02"] = make(map[string]string)
		studentMap["stu02"]["name"] = "mary"
		studentMap["stu02"]["sex"] = "女"
		studentMap["stu02"]["address"] = "上海黄浦江"
		fmt.Println(studentMap)
		fmt.Println(studentMap)
		fmt.Println(studentMap["stu02"])
		fmt.Println(studentMap["stu02"]["address"])
	*/
	//如果map的某一个key已经存在，则为修改 不存在则为增加  例：cities["no3"]="上海" 如果"no3"这个key已经存在 则该语句为修改 如果不存在则为增加
	var kong map[string]int //此时没有在堆中开辟空间所以不能往里添加key-value对
	kong = m                //将变量m的值赋值给变量kong     变量m中存储的值是一个地址  这个地址指向堆中空间
	fmt.Println(m)
	fmt.Println(kong)
	//var m = make(map[string]int)不是空map var kong map[string]int才是 判定切片、map是否为空的标准是看看其变量的值是否为nil

	//map删除key-value对
	var m5 = map[string]string{
		"name": "狗子3号",
		"sex":  "男",
	}
	delete(m5, "sex")   //删除key为"sex"的kay-value对
	delete(m5, "calss") //删除key为"class"的kay-value对 当key为"class"的kay-value对不存在时 delete函数不会进行删除操作，也不会报错

	/*	删除map变量所有key-value对
		方法1：如果我们要删除map变量里的所有key-value，没有一个专门的方法一次性删除多个key-value对 可以遍历一下 逐个删除
		方法2：使用map变量=make(...),在堆中开辟一个新的空间，让原来的那个空间成为垃圾，被gc回收
		//方法1
		var d1 = map[string]string{
			"name": "狗子7号",
			"addr": "翻斗花园102室",
		}
		for k, _ := range d1 {
			delete(d1, k)
		}
		fmt.Println(d1)
		//方法2
		var d = map[string]string{
			"name": "狗子6号",
			"addr": "翻斗花园101室",
		}
		fmt.Println(d)
		d = make(map[string]string) //在堆中开辟一个新的空间，让原来的那个空间成为垃圾，被gc回收
		fmt.Println(d)
	*/

}
