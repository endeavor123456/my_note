package main

import (
	"fmt"
)

type Person struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func swap(a *int, b *int) {
	temp := *a //temp是int类型 因为对*Type类型的指针变量取值 取出的值为Type类型的值(数据)
	*a = *b
	*b = temp
	fmt.Println(*a, *b)
}
func test(s *[]int) {
	*s = append(*s, 4, 5, 6)
}
func main() {
	/*
		&有类型的可以改变值的空间【也可以是某些类型的数据 例如：&[]int{1, 2, 3}、&Person{"狗子"}等】   *指针变量
		nil和0x0是一个东西 nil和0x0都是一个没有指向任何空间的指针
		var a *int
		//当指针变量(注意是指针变量)的值为nil或0x0时，对该指针变量取值会报错 这时候需要使用到new() 这样就可以对该变量取值了
		fmt.Println(a)
		a = new(int)    //new函数【格式为：new(Type)】会开辟一个没有名字的Type类型保存的值是可以改变的空间【它不是变量，他和变量的区别仅仅是没有名字，我给它命名为匿名非变量】 并返回该空间的地址，地址为*Type类型
		fmt.Println(*a) //结果为0，其实是int类型的默认值
		fmt.Println("------------------------------")
		var b *int
		b = new(int)
		*b = 10
		fmt.Printf("%T", *b)
		fmt.Println("-----------------------------")
		var a1 int
		var c *int = &a1 //&a1是取变量a1的地址
		fmt.Println(*c)  //*c就是变量a1
		var c1 = c   //*c1也是变量a1
		fmt.Println(*c1)
		fmt.Println("------------------------------")
	*/
	/*
		//切片指针[即切片类型的指针类型]变量
			slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
			p := &slice
			//p[1] = 222    //err 记住此处不能这样写即可
			(*p)[1] = 222 //这才对
	*/
	/*
		//指针切片[即为元素为指针类型的切片类型]变量 所以只能对元素取值，要和切片类型的指针类型区分开
		e := 10
		f := 20
		g := 30
		var slice = []*int{&e, &f, &g}
		fmt.Println(&slice)
		for i := 0; i < len(slice); i++ {
			fmt.Println(*slice[i])
		}
	*/
	/*
		//数组指针[即数组类型的指针类型]变量
		var arr [3]int = [3]int{1, 2, 3}
		var p *[3]int
		p = &arr
		(*p)[1] = 222 //此处可以简写为p[1] = 222系统内部优化的
		fmt.Println(arr)
		p[1] = 333
		fmt.Println(arr)
	*/
	/*
		//指针数组[即为元素为指针类型的数组,所以只能对元素取值，要和数组类型的指针区分开]
		e := 10
		f := 20
		g := 30
		var arr = [3]*int{&e, &f, &g}
		fmt.Println(&arr[0])
		for i := 0; i < len(arr); i++ {
			fmt.Println(*arr[i])       //arr[i]表示数组arr里名为i的元素
		}

	*/
	/*
		//指针变量作为函数参数
		a2 := 10
		b2 := 20
		swap(&a2, &b2)      //因为传递的是地址 此时形参可以改变实参的值
		fmt.Println(a2, b2) //输出 20 10     //详情请看./演示.png
	*/

	/*
		----------------------------------------
		-        -        -          -         -
		-        -        -          -         -
		-        -        -          -         -
		----------------------------------------
	*/
	/*
		//这个地方多看几遍
		h := [3]int{1, 2, 3}
		i := [3]int{4, 5, 6}
		j := [3]int{7, 8, 9}
		//这是一个长度为3，元素类型为*[3]int的数组
		var arr1 [3]*[3]int = [3]*[3]int{&h, &i, &j}
		fmt.Println(arr1)
		(*arr1[1])[1] = 555
		fmt.Println(arr1)
		//遍历此数组
		for i := 0; i < len(arr1); i++ {
			for j := 0; j < len(*arr1[i]); j++ {
				fmt.Printf("%d", (*arr1[i])[j])
			}
			fmt.Println("\t")
		}
	*/
	/*
		s := []int{1, 2, 3}
		fmt.Println(s)
		test(&s)
		fmt.Println(s)
	*/
	/*
		var per Person = Person{101, "张三疯", "男", 180, "武当三"}
		var p *Person = &per
		(*p).name = "张君宝" //可简写成p.name 虽然他的格式和per.name是一样的 但是表示的意思是不一样的 不要弄错了
	*/
	/*
		//[3]Person类型的数组
		arr := [3]Person{Person{101, "钢铁侠", "男", 180, "斯塔克公司"},
			Person{102, "绿巨人", "男", 180, "实验室"},
			Person{103,"黑寡妇","女", 180,"前苏联"}}
		p := &arr
		fmt.Printf("%p", p)
		for i := 0; i < len(p); i++ { //数组：v中的元素数（与len（v）相同）。

			//指向数组的指针：*v中的元素数（与len（v）相同）
			fmt.Println(p[i])
		}
	*/

	/*
		//[]Person类型的切片
		Slice := []Person{Person{101, "钢铁侠", "男", 180, "斯塔克公司"},
			Person{102, "绿巨人", "男", 180, "实验室"},
			Person{103, "黑寡妇", "女", 180, "前苏联"}}
		p := Slice
		fmt.Printf("%p", p)

		for i := 0; i < len(p); i++ { //数组：v中的元素数（与len（v）相同）。

			//指向数组的指针：*v中的元素数（与len（v）相同）
			fmt.Println(p[i])
		}
	*/

	/*
		//这段代码要多看看
		m := make(map[int]*[3]Person)
		m[1] = new([3]Person)
		m[1] = &[3]Person{Person{101, "钢铁侠", "男", 180, "斯塔克公司"},
			Person{102, "绿巨人", "男", 180, "实验室"},
			Person{103, "黑寡妇", "女", 180, "前苏联"}}
		m[2] = &[3]Person{Person{101, "索尔", "男", 180, "斯塔克公司"},
			Person{102, "绿巨人", "男", 180, "实验室"},
			Person{103, "黑寡妇", "女", 180, "前苏联"}}
		for k, v := range m {
			fmt.Println(k, *v)
		}
	*/
	/*
		//多重指针
		a := 10
		p := &a           //p为*int类型的变量
		var pp **int = &p //pp保存的是一个指向变量p的地址
		**pp = 100
		fmt.Println(a)
		//三级指针
		ppp := &pp
		fmt.Println(ppp)
	*/
}

//程序的销毁顺序是和运行顺序是相反的
//*Type类型的变量保存的是Type类型的空间(变量或者是没有名字的变量【他是一个变量，只不过没名字，但说的不是匿名变量】)
//常量的地址不允许被访问
