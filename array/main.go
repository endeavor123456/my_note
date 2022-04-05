package main

import "fmt"

//数组是值类型
//虽然数组和切片不是一种东西 但二者在使用上有非常多的相同之处
func main() {
	fmt.Println("ok")
	//数组变量也是数组名
	//数组的长度为该数组的元素的个数
	/*
		var arr [3]int        //定义一个长度为3元素为int类型的数组的长度也是类型的一部分
		fmt.Printf("%T", arr) //变量arr的类型是[3]int 表示此变量只能保存3个int类型的数据
	*/
	/*
		var arr [4]int    //长度为4的数组
		fmt.Println(arr) //会输出默认值
	*/
	/*
		var arr1 [3]int
		arr1[0] = 23   //arr1[0]表示数组里下标为0的元素 元素也是一个特殊的变量，特殊在它是在数组内部自动定义的且不是用变量名来标识一块空间 而是用数组下标来标识 切片的元素也是一个特殊的变量，特殊在它是在切片指向的在堆中的空间的内部自动定义的且不是用变量名来标识一块空间 而是用切片下标来标识 还有切片的元素在堆里且切片元素没有默认值  map理解起来比较麻烦需要单独看
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
		//根据初始值的个数 自行推断数组长度
		var arr4 = [...]int{1, 2, 3}
		fmt.Println(len(arr4))
		fmt.Println("---------------------------------------")
		//初始化数组时指定索引 golang的数组的索引都是0-n
		arr5 := [...]int{1: 1, 0: 2}
		fmt.Println(arr5)
	*/
	/*
			//数组遍历
			//方法一：
			var arr1 = [3]int{23, 34, 5}
			for i := 0; i < len(arr1); i++ {
				fmt.Println(arr1[i])
			}
			//方法二：
			for k, v := range arr1 {
				fmt.Println(k, v)
			}
			//求出一个数组里面元素的值的和以及这些元素的平均值 分别用for和for-range实现
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

			//从数组中找出相加和为8的元素的值
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
		//多维数组
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
		//还有一点需要注意 多维数组只有第一层可以使用 ...来让编译器推导数组长度
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
		//只要for语句出现在数组、切片就必然是遍历[但是反过来说就不对了：切片、数组只要是遍历必然是for语句 这句话是错的 因为还可以用for-range语句]
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
				使用sort包进行升序排序
				对于int、float64和string数组或是切片的排序，go分别提供了sort.Ints()、sort.Float64s()和sort.Strings()函数，默认都是从小到大排列

			intList := [...]int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
			floatList := [...]float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
			strList := [...]string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
			fmt.Println("排序前")
			fmt.Println(intList)
			fmt.Println(floatList)
			fmt.Println(strList)
			sort.Ints(intList)
			sort.Float64s(floatList)
			sort.Strings(strList)
			fmt.Println("排序后")
			fmt.Println(intList)
			fmt.Println(floatList)
			fmt.Println(strList)

		// 使用sort包进行降序排序
		intList1 := [...]int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
		floatList1 := [...]float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
		strList1 := [...]string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
		fmt.Println("排序前")
		fmt.Println(intList1)
		fmt.Println(floatList1)
		fmt.Println(strList1)
		sort.Sort(sort.Reverse(sort.IntSlice(intList1)))
		sort.Sort(sort.Reverse(sort.Float64Slice(floatList1)))
		sort.Sort(sort.Reverse(sort.StringSlice(strList1)))
		fmt.Println("排序后")
		fmt.Println(intList1)
		fmt.Println(floatList1)
		fmt.Println(strList1)
	*/

}
