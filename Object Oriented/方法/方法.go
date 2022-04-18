package main

import "fmt"

type Student3 struct {
	Age  int
	Name string
}
type Person2 struct {
	Name string
	Age  int
}
type Person3 struct {
	name string
	age  int
	sex  string
}

func (p Person2) getname() {
	fmt.Println(p.Name)
}
func (p Person2) getage(a, b int) {
	fmt.Println(p.Age)
	fmt.Println(a + b)
}
func (p Person3) PrintInfo() {
	fmt.Println("姓名", p.name, "年龄", p.age)
}
func (p *Person3) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

/*
String方法的格式：
		func (接收者变量 接收者变量类型) String() string{
			return ......
		}
*/
//String方法
func (Stu Student3) String() string {
	s := fmt.Sprintf("Name= [%v] Age = [%v]", Stu.Name, Stu.Age)
	return s
}
func main00() {
	/*
		var p1 = Person2{
		name: "张三",
		age:  20,
		sex:  "男",
		}
		p1.getname()       //调用方法
		p1.getage(10, 34) //调用方法
		fmt.Println("-----------------------------------")
		//接收者类型是用户用type关键字定义的类型的指针类型的方法有以下五种调用方式 *Type类型的变量里保存的是Type类型的可变空间的地址
		var p2 =Person3{
			name: "李四",
			age:  20,
			sex:  "女",
		}
		p2.PrintInfo()
		p2.SetInfo("狗子2号",30)

		var p2 =&Person3{
			name: "李四",
			age:  20,
			sex:  "女",
		}
		p2.PrintInfo()
		p2.SetInfo("狗子2号",30)

		var p2 =Person3{
			name: "李四",
			age:  20,
			sex:  "女",
		}
		var pp = &p2
		pp.PrintInfo()
		pp.SetInfo("狗子2号",30)

		var p2 = new(Person3)
		p.name="李四"
		p.age = 20
		p.sex="女"
		p2.PrintInfo()
		p2.SetInfo("狗子2号",30)

		var p2 Person3
		p.name="李四"
		p.age = 20
		p.sex="女"
		p2.PrintInfo()
		p2.SetInfo("狗子2号",30)
	*/
	/*
		var Stu = Student3{
			Name: "tom",
			Age:  20,
		}
		fmt.Println(&stu) 如果一个类型实现了String()这个方法，那么fmt.Println默认这个变量的String()进行输出
		//上面的代码的第二种写法
		var Stu = &Student3{
			Name: "tom",
			Age:  20,
		}
		fmt.Println(Stu)

	*/
}

/*
type test struct{  //类型其实是类型名 而它真正的类型是struct{ age int}
	age int
}
type int int  //类型其实是类型名 而它真正的类型是 int 类型也就是类型名的性质(存储结构、默认值、有无符号、大小、表数范围等等)是由它真正的类型决定的
*/
/*方法
在go语言中，没有类的概念但是可以给用户用type关键字定义的类型定义方法且接收者类型只能是用户用type关键字定义的类型或该类型的指针类型。
所谓方法就是定义了接收者的函数(方法是特殊的函数特殊在可以定义接收者且方法名可以同名【只有方法重写和接收者变量不同时方法才能重名】还有一个特殊之处是调用方法不同，要想使用方法首先得声明一个该方法接收者型的变量，然后以该变量.方法 来调用方法
例如：p1.PrintInfo()、p1.SetInfo("李四", 34)格式调用)。接收者的概念就类似于其他语言中的this或者self
当接收者的类型是用户用type关键字定义的类型和当接收者的类型是用户用type关键字定义的类型的指针类型时即能1.先声明一个该方法接收者型的指针类型的变量来调用 2.上面的方式调用
方法的定义(声明)格式
func (接收者变量 接收者类型) 方法名(形参列表)(返回参数列表){
	函数体
}
方法的接收者类型只可以是用户用type关键字定义的类型及其指针类型
结构体描述的是某一个体的属性(就是特征) 而结构体方法描述的该个体的行为
*/
