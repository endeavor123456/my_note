package main

import "fmt"

//map是引用类型
func main() {
	fmt.Println("ok")
	/*

		//map也可以在声明的时候填充
			var userInfo1 = map[string]string{ //变量userInfo1的变量类型为map[string]string
				"username": "张三",
				"age":      "20",
				"sex":      "男", //注意这个地方也有逗号
			}
			fmt.Println(userInfo1)
			fmt.Println(userInfo1["sex"])

			//上面那段代码是此段代码的简写
			var userInfo1 map[string]string = map[string]string{
					"username": "张三",
					"age":      "20",
					"sex":      "男", //注意这个地方也有逗号
			}
			fmt.Println(userInfo1)
			fmt.Println(userInfo1["sex"])

			//使用make()创建map   数组变量可以用“var arr [2]int”这种方式方式定义、切片变量可以用“var slice []int”这种方式定义，map变量虽然也可以“var userInfo map[string]string”这样定义 但是没啥意义 因为除会输出map[]之外啥也做不了
			var userInfo = make(map[string]string) //map对长度没要求
			userInfo["username"] = "张三"
			userInfo["age"] = "20"
			userInfo["sex"] = "男"   //注意这个地方都没有逗号
			fmt.Println(userInfo)
			fmt.Println(userInfo["username"])


			userInfo2 := map[string]string{
				"username": "张三",
				"age":      "20",
				"sex":      "男",
			}
			fmt.Println(userInfo2)
			fmt.Println(userInfo2["age"])
	*/
	/*
		//for-range 遍历map
		for k, v := range userInfo {
			fmt.Println(k, v)
		}
	*/
	/*
		//删除map里的某个键值对
		var userInfo3 = map[string]string{
			"username": "张三",
			"age":      "20",
			"sex":      "男",
			"height":   "188cm",
		}
		fmt.Println(userInfo3)
		delete(userInfo3, "sex")
		fmt.Println(userInfo3) //输出map[age:20 height:188cm username:张三]
	*/
	/*
		//我们想在切片里存放一系列用户的的信息，这时候我们就可以定义一个map类型的切片
		//创建map类型的切片(元素为map[string]string类型的切片)
		var userInfo = make([]map[string]string, 2, 3)
		userInfo[0] = make(map[string]string)
		userInfo[0]["username"] = "张三"
		userInfo[0]["age"] = "20"
		userInfo[0]["height"] = "180cm"

		userInfo[1] = make(map[string]string)
		userInfo[1]["username"] = "李四"
		userInfo[1]["age"] = "15"
		userInfo[1]["height"] = "175cm"
		fmt.Println(userInfo)
	*/
	/*
		//遍历map类型的切片
		var userInfo = make([]map[string]string, 2, 3)
		userInfo[0] = make(map[string]string)
		userInfo[0]["username"] = "张三"
		userInfo[0]["age"] = "20"
		userInfo[0]["height"] = "180cm"

		userInfo[1] = make(map[string]string)
		userInfo[1]["username"] = "李四"
		userInfo[1]["age"] = "15"
		userInfo[1]["height"] = "175cm"
		for _, v := range userInfo {
			for key, value := range v {
				fmt.Println(key, value)
			}
		}
	*/
	//value为切片类型的map
	var userInfo6 = make(map[string][]string)
	userInfo6["hobby"] = []string{
		"吃饭",
		"睡觉",
		"打豆豆",
	}
	fmt.Println(userInfo6)
	//map长度问题不必纠结
}
