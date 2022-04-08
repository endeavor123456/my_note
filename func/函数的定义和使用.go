package main

import "fmt"

//函数声明格式
/*
func 函数名(形参变量名1 类型,形参变量名2 类型,形参变量名3 类型,.../形参变量名1,形参变量名2,... 类型)(返回值类型1,返回值类型2,.../返回值参数名 类型){
	代码体
	return 变量或者返回值或表达式/只写return/没有返回值(没有return)
}
返回值是值(数据)
函数名是一个特殊的全局变量 特殊在定义方式不一样
//返回值函数可以返回多个返回值
func result(a, b int) int {   //a, b int其实是var a, b int   传参其实就是初始化即var a int = 传过来的值
	return a + b  //return 表达式  表示返回表达式 a+b的结果
}
func result(a, b int) int {
	sum:= a + b
	return sum  //return 变量    //表示返回变量sum里的值，此时变量sum里的值就是返回值
}
func result(a, b int) (sum int) {//sum int其实是var sum int
	sum = a + b
	return //只写return    //表示返回返回值参数变量的值，此时变量sum里的值就是返回值
}
func result() int {
	return "狗子3号"  //return 返回值
}
func result(a, b int) string {
	if(a + b>3){
		return "大于3"  //return 返回值
	}else{
		return "小于3"
	}
	func main() {
	res := result(1, 5)
	fmt.Println(res)
   }
}
........
保存返回值的参数(返回值参数)也是一个变量
形参和实参都是变量，当然实参也可以是数据(值)
注意：函数不能在函数内声明(匿名函数除外) 但可以在函数内调用
函数名首字母小写 函数的作用域为当前包 函数名大写 函数的作用域为整个项目
函数名在当前包是唯一的 [注：函数在当前包不仅不能和函数同名，也不能和当前包的全局变量(结构体字段除外,结构体字段是一种特殊的全局变量)和结构体名同名]
*/
func add(s1 int, s2 int) {
	sum := s1 + s2
	fmt.Println(sum)
}
func result(a, b int) int {
	return a + b
}
func main01() {
	fmt.Println("Hello")
	a := 10
	b := 20
	add(a, b)
	add(10, 20)
	add(10, 20)
	add(1, 2)
	res := result(1, 2) //调用函数并将函数返回值保存到变量res里
	fmt.Println(res)
}
