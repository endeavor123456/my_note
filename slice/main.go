package main

/*
1)切片类型是引用类型 不可对数组变量、切片不存在的元素进行操作 append函数是先开辟空间 这样就有新的元素了 再将数据赋值到新的元素里
2)切片是一个拥有相同类型元素的可变长度的序列，只能通过append函数和“切片名[0:0]”来改变切片长度【使用append函数插入和删除元素，删除切片全部元素使用“切片名[0:0]”】。
3)切片只能通过append函数进行插入和扩容 append函数先扩容再追加
4)使用切片名、map名来分辨是不是同一切片、map
5）append函数返回值是切片类型的数据(切片类型的数据保存在堆中，切片类型的变量保存的是一个地址)，它会复制之前切片的数据并和我们要插入的数据一并作为返回值 用之前的切片名(切片变量)来保存它的地址会把之前的地址覆盖掉 这样就成了一个新切片
6)切片的索引(下标)为0-n
7)切片类型的变量保存的是一个地址 这个地址指向的是它在堆中存放数据的空间
虽然引用类型的变量中存放的是一个地址，但要输出它在堆中的数据 直接默认输出即可
a := make([]int, 10) //切片make之后才会输出默认值 这是一个长度为10 元素的值都是0的切片
fmt.Println(a)
只有对切片名、map名取地址、取值(不可以对切片和map类型的变量取值)以及传参还有切片名之间、map名之间相互赋值不是对切片名、map名里保存的地址指向的堆里的空间进行操作 其他对切片名和map名的操作就是对切片名、map名里保存的地址指向的堆里的空间进行操作 不要考虑对切片名和map名进行取值操作的任何事
*/
func main() {
	/*
				var slice []int   //定义一个[]int类型的变量【如果是宽泛的说法变量slice是一个切片类型的变量，如果是说详细了就是[]int类型的，就跟整数类型和int类型之间的关系一样】 此时变量slice里保存的值为nil(0X0) 此时切片没有长度【因为没有在堆中开辟空间所以没有长度，长度为0是因为实现切片类型的结构体的用来保存切片长度的字段没有接收到值所以是默认值】 容量为0 不能通过slice[0] = 1这样的方式 因为此时切片长度为0，只有长度不为0时才能通过slice[0] = 1这样的方式 此时输出slice[0]同理 此时得用make函数设置长度【就是用make函数在堆中开辟空间，注意：切片使用make函数是必须设置长度，容量可写可不写】 就可以用slice[0] = 1这样的方式了
				fmt.Println(slice) //空map是指堆中空间没有key-value，空切片是在堆中开辟一个长度为零的空间，var slice []int这不是一个空切片，因为它没有在堆中开辟空间，一个空间为空的前提是得有这个空间 而且切片一旦开辟空间每个元素就都有默认值了，也就不是空的了


				//定义一个空切片(两种方式)
				silce := make( []int , 0 )
				slice := []int{}


				var slice2 = []int{1, 2, 3, 4, 5}
				fmt.Println(slice2, len(slice2))

				var slice3 = []int{1: 2, 2: 4, 3: 6}
				fmt.Println(slice3)

				slice4 := []int{1, 2, 3, 4}
				fmt.Println(slice4)
				slice5 := []int{0: 1, 1: 2, 2: 3}
				fmt.Println(slice5)

				//基于数组变量定义切片(看看下面的讲解)
				a := [5]int{55, 56, 57, 58, 59}
				b := a[:] //获取数组变量里的所有的元素的值 必须得[下标1:下标2]【下标1 >=0/不写,下标2 >=0/不写,下标1 <= 下标2】
				fmt.Println(b)

				c := a[1:4] //获取数组变量里的指定的元素的值   左包右不包  必须得[下标1:下标2]【下标1 >=0/不写,下标2 >=0/不写,下标1 <= 下标2】
				fmt.Println(c)
				d := a[2:] //从数组变量索引(下标)为2的元素开始获取全部 必须得[下标1:下标2]【下标1 >=0/不写,下标2 >=0/不写,下标1 <= 下标2】
				fmt.Println(d)
				e := a[:3] //获取数组变量下标为3的元素前面的元素的值 必须得[下标1:下标2]【下标1 >=0/不写,下标2 >=0/不写,下标1 <= 下标2】
				fmt.Println(e)

				//基于切片定义切片(看看下面的讲解)
				a := []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
				b := a[1:] //从切片下标为1的元素开始获取全部
				fmt.Println(b)
					.........
		        //通过make函数定义切片  make([]T,size,cap)
				var sliceA = make([]int,4,8)   //长度为4 容量为8的空切片 对切片使用make函数时size必须写，cap可省略【切片如果不设置容量，容量==长度】 如果是对map使用make函数 size和cap都可省略
				fmt.Println(sliceA)    //会输出默认值
				sliceA[0] = 10
				sliceA[1] = 12
				sliceA[2] = 10
				sliceA[3] = 30
				fmt.Println(sliceA)

	*/
	/*
		sliceB := []string{"php","java","go"}
		sliceB[2] = "golang"
		fmt.Println(sliceB)    //输出["php","java","golang"]
	*/
	/*
		//append函数
		var sliceB []int
		var sliceE []int
		sliceB = append(sliceB,3)
		sliceE = append(sliceE,3,4,5,6,3)  //append函数可以追加1个或多个元素
		fmt.Println(len(sliceB))
		fmt.Println(len(sliceE))

		//append函数还可以和并切片
		sliceC := []string{"php","java"}
		sliceD := []string{"nodejs","go"}
		sliceC = append(sliceC,sliceD...)   //append函数插入的数据的数据类型必须与第一个实参【切片类型的变量/数据】的元素/元素的值的类型相同，这个地方是把sliceD打散 追加到sliceC里面的 和sliceC = append(sliceC,"nodejs","go")等价
		fmt.Println(sliceC)

		//切片的扩容策略【对于切片的扩容策略只需了解即可】
		var sliceE []int
		for i := 0;i<10;i++{
			sliceE = append(sliceE,i)
			fmt.Println(sliceE,"长度",len(sliceE),"容量",cap(sliceE))
		}
		//copy()函数
		sliceA := []int{1, 2, 3, 4, 5}
		sliceB := make([]int, 4, 4)
		copy(sliceB, sliceA) //copy函数的作用是让身为引用类型的切片变量也可以进行类似于值传递的操作
		sliceB[0] = 11
		fmt.Println(sliceA)
		fmt.Println(sliceB)

		//删除切片中的元素
		a := []int{1, 2, 3, 4, 5, 6, 7}
		a = append(a[:2], a[3:]...)//切片名[下标1:下标2]【下标1 >=0/不写,下标2 >=0/不写,下标1 <= 下标2】
		fmt.Println(a) //输出[1,2,4,5,6,7]

		//删除切片中的所有元素方法
		delslice := []int{1,2,3,4,5,6}
		fmt.Println(delslice)
		delslice = delslice[0:0]//删除切片中的所有元素 切片名[下标1:下标2]【下标1 >=0/不写,下标2 >=0/不写,下标1 <= 下标2】
		fmt.Println(delslice)
	*/
	/*

		//--------------这就是提到的“看看下面的讲解”--------------------------------------------------------
		slice := []int{1, 2, 3, 4, 5}
		fmt.Println(slice)
		slice1 := slice[0:2] //注意不要认为是：slice1 := []int{1,2}   其实slice1里保存的地址是指向slice变量里存储的地址指向的空间的某一个或几个元素
		slice1[1] = 8
		fmt.Printf("%p", slice)
		fmt.Println()
		fmt.Printf("%p", slice1)
		fmt.Println("----------------------------------------")
		var slice5 = []int{1, 2, 3, 4, 5, 6}
		slice5 = slice5[0:0]//清空所有元素
		fmt.Println(slice)
		//还有一点要注意 切片名里保存的地址和该地址指向的堆中空间的第一个元素的地址相同 重合了
		var array = [5]int{1, 2, 3, 4, 5}
		fmt.Println(array)
		var slice3 = array[0:2]
		fmt.Println(slice3)
		array[0] = 10
		fmt.Println(array)
		fmt.Println(slice3)
		//数组变量的地址和数组元素的地址相同 也是重合
		//--------------这就是提到的“看看下面的讲解”--------------------------------------------------------

	*/
	/*
		a, b, c := 1, 2, 3
		var slice = []int{a, b, c, a + b, arrayelements(1, 2), arrayelements1()} //将a,b,c的值，a+b的结果，函数arrayelements，arrayelements1的返回值分别赋值到对应的元素中
		fmt.Println(slice)
		fmt.Println([]int{a, b, c, a + b, arrayelements(1, 2), arrayelements1()})
		fmt.Println([]int{a, b, c, a + b, arrayelements(1, 2), arrayelements1()}[0]) //输出切片数据里保存在第一个元素的值
		var slice1 = []int{a, b, c, a + b, arrayelements(1, 2), arrayelements1()} //将a,b,c的值，a+b的结果，函数arrayelements，arrayelements1的返回值分别赋值到对应的元素中
		fmt.Println(slice1)
	*/

	/*
		//关于切片的长度和容量
		//切片的长度即为其所保存的元素个数
		//切片的容量需要重视
		//切片的容量是从它的第一个元素开始数到其底层切片元素末尾（包含末尾）的元素个数
		//以b := s[1:3]为例 切片b的第一个元素是下标为1的元素 此元素到切片s的末尾的元素的个数为5
		//切片变量就是切片名
		s := []int{2, 3, 5, 7, 11, 13}  //长度为6 容量为6
		fmt.Println("长度为", len(s), "容量为", cap(s))

		a := s[2:]
		fmt.Println("长度为", len(a), "容量为", cap(a))

		b := s[1:3]
		fmt.Println("长度为", len(b), "容量为", cap(b))
		e := s[4:]
		fmt.Println("容量为", cap(e))
	*/
	/*
		//切片的循环遍历 遍历即把元素一个个的拿出来
		var strSlice = []string{"济南", "北京", "上海", "南京", "深圳"}
		//遍历并输出遍历出来的元素的值
		for i := 0; i < len(strSlice); i++ {
			fmt.Println(strSlice[i])
		}
		for i, v := range strSlice {
			fmt.Println(i, v)
		}
	*/
	/*
		//选择排序
		var numSlice = []int{9, 8, 7, 6, 5, 4}
		//var numSlice = []string{"a", "d", "e", "g", "c"}
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
	/*
		//冒泡排序
		var numSlice = []int{9, 6, 5, 3, 8}
		//var numSlice = []string{"a", "d", "e", "g", "c"}
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
		                排序原理：---- ---------
						使用sort包进行升序排序
						对于int、float64和string切片的排序，go分别提供了sort.Ints()、sort.Float64s()和sort.Strings()函数，但这几个函数不能对数组排序，都是从小到大排列

					intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
					floatList := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
					strList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
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
				intList1 := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
				floatList1 := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
				strList1 := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
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
	/*
		//利用切片修改字符串数据
			s1 := "big"
			byteStr := []byte(s1)
			byteStr[0] = 'p'
			fmt.Println(string(byteStr))
	*/
}
