package main

import "fmt"

type Person struct {
	name string
}

func add(args ...int) { //此处为不定参 表示可以将任意个int类型的值(数据)传到[]int类型的形参变量args里
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	fmt.Println(sum)
}
func main33() {
	add([]int{1, 3, 7}...) //将此[]int类型的数据 打散 即把这一个数据的集合打散成一个个单个的数据
	fmt.Printf("%T", []int{1, 2, 3})
	/*
		[]int{1, 2, 3}是[]int类型的数据 []int型的数据是一组任意个int型元素的值的集合 不是切片元素的集合 切片名里保存的地址指向的堆中的那块空间才是切片元素的集合
		[3]int{1,2,3}  是[3]int{1,2,3}类型的数据 [3]int型的数据是3个int型元素的值的集合 不是数组元素的集合 数组名才是是元素的集合
		Person{"狗子"}为Person类型的数据 是一组字段的值的集合 不是字段的集合 结构体变量才是字段的集合
		fmt.Println([]any{1, 2, 3}[0])       //表示输出[]int{1, 2, 3}这个集合里保存在第一个元素里的值
		fmt.Println(res{10, "狗子", "男"}.name) //表示输出res{10, "狗子", "男"}这个集合里保存在字段name里的值
	*/
	var slice = []int{1, 2, 3}
	fmt.Println(slice)

	var per = Person{"狗子"}
	fmt.Printf("%T", per)
}
