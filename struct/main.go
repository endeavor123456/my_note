package main

import "fmt"

//golang中的类型按照定义的角色不同可分为用户自定义类型和系统类型 用户自定义类型即用户使用type关键字定义的类型 系统类型即系统使用type关键字定义的类型
//结构体是一种用户自定义类型【用户使用type关键字定义的类型】 是使用type和struct关键字定义的【使用type和struct关键字定义的类型统称为结构体类型】 可以叫他结构体、也可以叫他结构体类型
//结构体是值类型
//type关键字在函数外部使用 golang中的关键字在vscode中显示为蓝色 类型定义之后就用类型名代表该真正的类型 所以类型名不同则类型不同 对于类型名不同 真正的类型相同可以理解成虽然二者的真正的类型是相同的但不是同一个个体 可以类比一下两个相同的硬币虽然两个硬币相同但是二者不是同一个硬币
//类型名、类型别名不能与当前包已经存在的类型名重复
//类型名小写只能在当前包里使用 大写则能在整个项目里使用
//类型名大写时，字段名首字母大写该字段才能被别的包使用,否则只能在当前包使用
/*
这个地方先这样理解，后期再改！！！！！！
结构体字段是一种特殊的全局变量 第一个特殊之处是可以和函数名同名
第二个特殊之处就是声明方式不同且只有这一种声明方式
第三个特殊之处是使用方式不同 拿结构体变量为例：该结构体类型的变量.字段 如果是结构体类型的切片：slice[0].字段
*/
/*
type 类型名 struct{
	字段名1 字段类型  //在结构体中字段名必须唯一
	字段名2 字段类型
}
关于结构体类型变量的强制类型转换 类型Student类型和Student1类型虽然不是同一类型但是真正的类型相同 所以Student和Student1类型的变量可以进行强制数据类型转换
type Student struct {
	id    int
	name  string
	age   int
	sex   string
}
type Student1 struct{
	id int
	name string
	age int
	sex string
}
func main(){
var Stu Student
	var Stu1 Student1
	Stu.age=10
	Stu.id=1
	Stu.name="狗子5号"
	Stu.sex="男"
	fmt.Println(Stu)
	Stu1 = Student1(Stu)
	fmt.Println(Stu1)
*/
type Student struct { //定义一个类型名为Student的结构体类型
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}
type Int int //定义一个类型名为Int的整数类型
type Person struct {
	name string
	age  int
	sex  string
}
type Monster struct {
	Name string `json:"姓名"`
	Age  int    `json:"年龄"`
}

func main() {
	fmt.Println("ok")
	/*
		实例化结构体(即定义一个对象 即：一个该结构体类型的变量(结构体变量)并赋值)
		var p1 Person //声明一个Person类型的变量(结构体变量)   var p2 []Person  定义一个结构体切片类型的变量
		p1.name = "狗子1号"
		p1.sex = "男"
		p1.age = 10
		fmt.Println(p1)

		var p4 = Person{
			name: "狗子4号",
			age:  20,
			sex:  "女",
		}
		fmt.Println(p4)

		p5 := Person{
			name: "狗子4号",
			age:  20,
			sex:  "女",
		}
		fmt.Println(p5)

		var p6 = Person{} //会输出字段默认值 {空串 0 空串}
		fmt.Println(p6)

		var p2 = new(Person) //new(Type)返回的是一个Type类型的指针 所以p2保存的是一个Person类型的指针
		p2.name = "狗子2号"
		p2.sex = "女"
		p2.age = 11
		fmt.Println(p2)
		fmt.Println((*p2).age) //在这个语境中可以简写成p2.age

		var p3 = &Person{} //p3也是一个指针类型变量 结构体专门的一种实例化方法和&p3等价
		p3.name = "狗子3号"
		p3.age = 12    //(*p3).age
		p3.sex = "男"
		fmt.Println(p3.age)   //(*p3).age 可以简写成p3.age  这是系统优化的
		fmt.Println(p3)     //注意：(*p3).age 可以简写成p3.age 但是(*p3)不可以简写成p3

		var p6 = &Person{ //p6还是指针类型的变量 结构体专门的一种实例化方法和&p3等价
			name: "狗子5号",
			age:  21,
			sex:  "男",
		}
		fmt.Println(p6)

		var p7 = &Person{ //p7还是一个指针类型
			"狗子6号",
			20,
			"女",
		}
		fmt.Println(p7.age)
	*/
	/*
		//结构体数组
		var arr [5]Student = [5]Student{
			Student{101, "唐三藏", 32, "男", 100, "北京"},
			//{101, "唐三藏", 32, "男", 100, "北京"}, 也可以这样写
			Student{102, "孙悟空", 532, "男", 100, "花果山"},
			Student{103, "猪悟能", 600, "男", 100, "高老庄"},
			Student{104, "sah", 800, "男", 100, "六山河"},
			Student{105, "sshsh", 400, "男", 100, "济南"}}
		arr[1].score = 60
		arr[2].score = 10
		fmt.Println(arr)
		fmt.Println(arr[2].age)
	*/
	/*
		//结构体数组作为函数参数
		// func Sort (arr [5]Studnet){
		// 	.........
		// }
	*/
	/*
		//结构体切片
		var slice []Student = []Student{
			Student{101, "唐三藏", 32, "男", 100, "北京"},
			//{101, "唐三藏", 32, "男", 100, "北京"}, 也可以这样写
			Student{102, "孙悟空", 532, "男", 100, "花果山"},
			Student{103, "猪悟能", 600, "男", 100, "高老庄"},
			Student{104, "sah", 800, "男", 100, "六山河"},
			Student{105, "sshsh", 400, "男", 100, "济南"}}
		slice = append(slice, Student{106, "狗子", 999, "男", 100, "北京"})
		fmt.Println(slice)
	*/
	/*
		//结构体转json字符串
		monster := Monster{"牛魔王", 500}
		jsonStr, err := json.Marshal(monster)
		if err != nil {
			fmt.Println("json 处理错误", err)
		}
		fmt.Println("jsonStr", string(jsonStr))

		//json字符串转结构体
		var str = `{"姓名":"牛魔王","年龄":500}`
		var s1 Monster
		error := json.Unmarshal([]byte(str), &s1)
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(s1)
		fmt.Println(s1.Name)
	*/
}

// 和该结构体类型相关的类型的变量即：该结构体类型的切片变量、结构体类型的数组变量 结构体类型的map变量
