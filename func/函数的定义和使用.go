package main

import "fmt"

//函数声明的几种格式
/*
func 函数名() 返回值类型 {
		代码体
		return 变量或者值(数据)或表达式
	}
func 函数名() (返回值类型,返回值类型,....){
		代码体
		return 变量或者值(数据)或表达式
	}
func 函数名(形参变量名1 类型,形参变量名2 类型,形参变量名3 类型,.../形参变量名1,形参变量名2,... 类型) 返回值类型 {
				代码体
				return 变量或者值(数据)或表达式
	}
func 函数名(形参变量名1 类型,形参变量名2 类型,形参变量名3 类型,.../形参变量名1,形参变量名2,... 类型) (返回值类型,返回值类型,....){
		代码体
		return 变量或者值(数据)或表达式
	}
func 函数名(形参变量名1 类型,形参变量名2 类型,形参变量名3 类型,.../形参变量名1,形参变量名2,... 类型) (返回值参数名 类型){
			代码体
	        return
	}
func 函数名(形参变量名1 类型,形参变量名2 类型,形参变量名3 类型,.../形参变量名1,形参变量名2,... 类型)(返回值参数名1 类型,返回值参数名2 类型,...){
				代码体
		        return
	}

返回值是值(数据)
函数名是一个特殊的全局变量 特殊在定义方式不一样
//返回值函数可以返回多个返回值
func result(a, b int) int {   //a, b int其实是var a, b int   传参其实就是初始化即var a int = 传过来的值
	return a + b  //return 表达式  表示返回表达式 a+b的结果
}
func result(a, b int) int {
	sum:= a + b
	return sum  //return 变量    //表示结束函数result并返回变量sum里的值，此时变量sum里的值就是返回值
}
func result(a, b int) (int,int) {//多个返回值 必须加括号
	sum:= a + b
	sum1 := a-b
	return sum,sum1  //return 变量    //表示结束函数result并返回变量sum,sum1里的值，此时变量sum,sum1里的值就是返回值
}
func result(a, b int) (sum int) {//sum int其实是var sum int   格式为：返回值参数名 类型 的必须加括号
	sum = a + b
	return //只写return    //表示结束函数result并返回返回值参数变量的值，此时变量sum里的值就是返回值
}
func result(a, b int) (sum int,sum1 int) {
	sum = a + b
	sum1 = a-b
	return //只写return    //表示结束函数result并返回返回值参数变量的值，此时变量sum里的值就是返回值
}
func result() int {
	return "狗子3号"  //return 返回值
}
func result()(int,int,string){//多个返回值 必须加括号
	return 1,2,"狗子"
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
保存返回值的参数(返回值参数)也是变量
形参变量就是变量，只不过加了限制条件：没有默认值且只能一种定义方法。实参既可以是变量也可以是数据(值)还可以是表达式，切片、数组元素、map元素、函数返回值、结构体字段
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
