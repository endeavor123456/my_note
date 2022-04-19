package main

import "fmt"

//数组类型是值类型
//虽然数组类型和切片类型不是一种类型 但二者在使用上有非常多的相同之处
func main() {
	fmt.Println("ok")
	//数组变量也是数组名 数组变量即为数组类型的变量
	//数组的长度为该数组的元素的个数
	/*
		var arr [3]int        //定义一个长度为3元素为int类型的数组变量 长度也是数组类型的一部分
		fmt.Printf("%T", arr) //变量arr的类型是[3]int 表示此变量只能保存3个int类型的数据
	*/
	/*
		var arr [4]int    //长度为4的数组
		fmt.Println(arr) //会输出默认值
	*/
	/*
		var arr1 [3]int
		arr1[0] = 23   //arr1[0]表示数组变量里下标为0的元素 数组的元素和变量的区别在于它不需要定义它是在定义数组变量时在数组变量内部又自动开辟的一块空间且用数组下标来表示这一块空间 切片的元素和变量的区别在于它是在切片指向的在堆中的空间的内部自动或手动开辟的一块空间且用切片下标来表示这一块空间  map理解起来比较麻烦需要单独看
		arr1[1] = 10
		arr1[2] = 10
		fmt.Println(arr1)
		fmt.Println("-----------------------------------")
		var arr2 = [2]string{"php", "nodejs"}
		fmt.Println(arr2)
		fmt.Println("------------------------------------")
		arr3 := [2]string{"php", "nodejs"}
		fmt.Println(arr3)
		fmt.Println("--------------------------------------")
		//根据初始值的个数 自行推断数组变量长度【即数组变量里元素的个数】
		var arr4 = [...]int{1, 2, 3}
		fmt.Println(len(arr4))
		fmt.Println("---------------------------------------")
		//初始化数组时指定索引 golang的数组变量的索引都是0-n
		arr5 := [...]int{1: 1, 0: 2}
		fmt.Println(arr5)
	*/
	/*
			//遍历数组中的元素  遍历即把元素一个个的拿出来
			//方法一,使用for语句遍历：
			var arr1 = [3]int{23, 34, 5}
			//遍历并输出遍历出来的元素的值
			for i := 0; i < len(arr1); i++ {
				fmt.Println(arr1[i])
			}
			//方法二，使用for-range语句遍历：
			for k, v := range arr1 {
				fmt.Println(k, v)
			}
			//求出一个数组变量里面元素的值的和以及这些元素的平均值 分别用for和for-range实现
			var arr2 = [...]int{9, 4, 7, 2, 3, 5}
			var sum = 0
			for i := 0; i < len(arr2); i++ {
				sum += arr2[i] //求和
			}
			fmt.Println(sum)
			avg := float64(sum) / float64(len(arr2)) //求平均值
			fmt.Println(avg)
			fmt.Println("--------------------------------------------")
			//求最大值
			var max = arr2[0]
			var index = 0
			for i := 0; i < len(arr2); i++ {
				if max < arr2[i] {
					max = arr2[i]
					index = i
				}
			}
			fmt.Println("最大值为", max, "对应的索引为", index)

			//从数组变量中找出相加和为8的元素的值
		var arr = [...]int{1, 3, 5, 7, 8}
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i]+arr[j] == 8 {
					fmt.Println(i, j)
				}
			}
		}
	*/
	/*
		//定义多维数组变量
		var arr = [3][2]string{
			{"北京", "上海"},
			{"广州", "深圳"},
			{"成都", "重庆"}, //注意这个地方
		}
		fmt.Println(arr[0])
		fmt.Println(arr[0][0])
		arr2 := [3][2]string{
			{"北京", "上海"},
			{"广州", "深圳"},
			{"成都", "重庆"}, //注意这个地方
		}
		fmt.Println(arr2)
		var arr3 [2][2]int
		arr3[0][0] = 1
		arr3[0][1] = 2
		arr3[1][0] = 3
		arr3[1][1] = 4
		fmt.Println(arr3)

		for _, v1 := range arr {
			for _, v2 := range v1 {
				fmt.Println(v2)
			}
		}

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[i]); j++ {
				fmt.Println(arr[i][j])
			}
		}
		//还有一点需要注意 定义多维数组变量时只有第一层可以使用...    ...是让编译器推导数组长度
		//正确的写法
		a := [...][2]string{
			{"北京", "上海"},
			{"广州", "深圳"},
			{"成都", "重庆"},
		}

		//错误的写法
		b := [3][...]string{
			{"北京", "上海"},
			{"广州", "深圳"},
			{"成都", "重庆"},
		}
	*/
	/*
		//选择排序：进行从小到大排序
		//概念：通过比较 首先选出最小的数 放在第一个位置上 之后在其余的数中选出次小数放在第二个位置 以此类推
			var numSlice = [...]int{9, 8, 7, 6, 5, 4}
			//var numSlice = [...]string{"a", "d", "e", "g", "c"}
			for i := 0; i < len(numSlice); i++ {
				for j := i + 1; j < len(numSlice); j++ {
					if numSlice[i] > numSlice[j] {
						temp := numSlice[i]
						numSlice[i] = numSlice[j]
						numSlice[j] = temp
					}
				}
			}
			fmt.Println(numSlice)
	*/
	fmt.Println("-------------------------------------------")

	/*
		//数组和切片的很多操作都离不开for或for-range 就拿排序、求和来说，第一步都是先遍历 然后对遍历出来的元素进行判断等的操作
			//冒泡排序概念：从头到尾，比较相邻的两个元素的大小，如果符合交换条件，交换两个元素的位置
			//特点：每一轮比较中，都会选出一个最大的数，放在正确的位置

		//冒泡排序
		var numSlice = [...]int{9, 6, 5, 3, 8}
		//var numSlice = [...]string{"a", "d", "e", "g", "c"}
		for i := 0; i < len(numSlice); i++ {
			for j := 0; j < len(numSlice)-1-i; j++ {
				if numSlice[j] > numSlice[j+1] {
					temp := numSlice[j]
					numSlice[j] = numSlice[j+1]
					numSlice[j+1] = temp
				}
			}
		}
		fmt.Println(numSlice)

	*/

	/*
		数组是值类型,不能直接排序，必须转为切片
		 排序原理：---- ---------
		使用sort包进行升序排序
		对于int、float64和string切片的排序，go分别提供了sort.Ints()、sort.Float64s()和sort.Strings()函数，但这几个函数不能对数组排序
	*/

}
