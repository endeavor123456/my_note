package main

type Stu struct {
	Name    string
	Age     int
	Address string
}

func modifyUser(users map[string]map[string]string, name string, nickname string) {
	//先判断是否有
	//一般使用if v,ok := users[name];ok{....}   因为此处map的value是map类型的数据 所以此处还可因写成：if users[name] != nil{....}
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = nickname + name
	}
}

//map知识点补充
//map很多地方都挺怪
//map对长度不敏感
//map的key具有唯一性
func main() {
	//map变量的key的类型不能是切片、map、function类型
	/*
		map变量的定义格式之一，只列出这一个，其他的定义格式就不一一列举了
		var map名 = make(map[key的类型]value的类型)
	*/
	/*
		//map、切片类型的比较 详见golang中的比较注意事项.md
		var m = make(map[string]int) //map变量中的key-value对是无序的
		fmt.Println(m)
	*/
	/*
		   课堂练习
		   存放3个学生的信息，每个学生有max和sex信息，并遍历出来

		studentMap := make(map[string]map[string]string)
		studentMap["stu01"] = make(map[string]string)
		studentMap["stu01"]["name"] = "tom"
		studentMap["stu01"]["sex"] = "男"
		studentMap["stu01"]["address"] = "北京长安街"

		studentMap["stu02"] = make(map[string]string)
		studentMap["stu02"]["name"] = "mary"
		studentMap["stu02"]["sex"] = "女"
		studentMap["stu02"]["address"] = "上海黄浦江"
		fmt.Println(studentMap)
		fmt.Println(studentMap)
		fmt.Println(studentMap["stu02"])
		fmt.Println(studentMap["stu02"]["address"])
		for k1,v1 := range studentMap{  //遍历 map只能用for-range遍历  不能像切片和数组那样即可以用for-range可以用for
			fmt.Println("k1=",k1)
			for k2,v2:=range v1{
				fmt.Printf("\t k2=%v v2 = %v\n",k2,v2)
			}
			fmt.Println()
		}

	*/
	/*
		//如果map变量的某一个key已经存在，则为修改 不存在则为增加  例：cities["no3"]="上海" 如果"no3"这个key已经存在 则该语句为修改 如果不存在则为增加
		var kong map[string]int //此时没有在堆中开辟空间所以不能往里添加key-value对
		kong = m                //将变量m的值赋值给变量kong     变量m中存储的值是一个地址  这个地址指向堆中空间
		fmt.Println(m)
		fmt.Println(kong)
		//var m = make(map[string]int)不是空map var kong map[string]int才是 判定切片、map是否为空的标准是看看其变量的值是否为nil

	*/
	/*
		//删除map变量中的key-value对
		var m5 = map[string]string{
			"name": "狗子3号",
			"sex":  "男",
		}
		delete(m5, "sex")   //删除key为"sex"的kay-value对
		delete(m5, "calss") //删除key为"class"的kay-value对 当key为"class"的kay-value对不存在时 delete函数不会进行删除操作，也不会报错

	*/
	/*	删除map变量所有key-value对
		方法1：如果我们要删除map变量里的所有key-value，没有一个专门的方法一次性删除多个key-value对 可以遍历一下 逐个删除
		方法2：使用map变量=make(...),在堆中开辟一个新的空间，让原来的那个空间成为垃圾，被gc回收
		//方法1
		var d1 = map[string]string{
			"name": "狗子7号",
			"addr": "翻斗花园102室",
		}
		for k, _ := range d1 {
			delete(d1, k)
		}
		fmt.Println(d1)
		//方法2
		var d = map[string]string{
			"name": "狗子6号",
			"addr": "翻斗花园101室",
		}
		fmt.Println(d)
		d = make(map[string]string) //在堆中开辟一个新的空间，让原来的那个空间成为垃圾，被gc回收
		fmt.Println(d)
	*/
	/*
		//查找map中的key-value对
		var d3 = map[string]string{
			"name": "狗子8号",
			"addr": "翻斗花园103室",
		}
		//查找map中的key-value对的第一种写法
		val, ok := d3["age"] //如果d3中存在key为"age"的key-value对，ok的值就为true，如果不存在ok的值就为false    ？？？这个地方有些别扭
		if ok {
			fmt.Println(`d3中存在key为"age"的key-value对,value为：`, val)
		} else {
			fmt.Println("不存在")
		}
		//查找map中的key-value对的第二种写法
		if val, ok := d3["age"]; ok {
			fmt.Println(`d3中存在key为"age"的key-value对,value为：`, val)
		} else {
			fmt.Println("不存在")
		}
	*/
	/*
		map类型的切片  就如 int类型、int32类型、int64类型、.....统称为整数类型一样 map[string]string类型、map[int]string类型、....统称为map类型
		练习：

		var monsters []map[string]string
		monsters = make([]map[string]string, 2) //切片如果不设置容量则容量==长度
		if monsters[0] == nil {
			monsters[0] = make(map[string]string, 2)
			monsters[0]["name"] = "牛魔王"
			monsters[0]["age"] = "500"
		}
		if monsters[1] == nil {
			monsters[1] = make(map[string]string, 2)
			monsters[1]["name"] = "玉兔精"
			monsters[1]["age"] = "400"
		}
		newmonster := map[string]string{
			"name": "新的妖怪 火云邪神",
			"age":  "200",
		}
		fmt.Printf("%p", monsters)

		monsters = append(monsters, newmonster) //append函数插入的数据的数据类型必须与第一个实参【切片类型的变量/数据】的元素/元素的值的类型相同
		fmt.Println(monsters)
	*/
	/*
	   golang中map变量中的key-value对是无序的，所以在对map for-range遍历的时候也是无序的
	   对map变量中的key-value对排序


	*/
	/*
		//map的排序有点不一样，是按照key来对key-value对排序
		var m8 = map[string]string{
			"name": "胡图图",
			"addr": "翻斗花园",
		}
		var key []string
		for k := range m8 {
			key = append(key, k)
		}
		sort.Strings(key)
		for _, k := range key { //for-range 遍历的两个位置的变量分别用来接收不同的数据【这两个变量不用都写上，可以只写一个】 第一个位置的数据不需要变量接收，如果不放匿名变量顺序会乱
			fmt.Println(k, m8[k])
		}
	*/

	/*
		//map课堂练习：
		如果用户存在就将其密码修改为“888888”,如果不存在就增加这个用户信息(用户信息包括用户名、昵称、密码)
		使用函数实现
		users := make(map[string]map[string]string)
		users["狗子1号"] = make(map[string]string)
		users["狗子1号"]["pwd"] = "999999"
		users["狗子1号"]["nickname"] = "胡图图"
		modifyUser(users, "狗子2号", "胡英俊")
		modifyUser(users, "狗子3号", "张小丽")
		fmt.Println(users)

	*/

	/*
		map的value也经常使用结构体类型 更适合管理复杂的数据(比前面的value是map类型的更好)
		students := make(map[string]Stu)
		stu1 := Stu{"tom", 18, "北京"}
		stu2 := Stu{"mary", 18, "上海"}
		students["no1"] = stu1
		students["no2"] = stu2
		fmt.Println(students)

		for k, v := range students {
			fmt.Printf("学生的编号是%v\n", k)
			fmt.Printf("学生的名字是%v\n", v.Name)
			fmt.Printf("学生的编号是%v\n", v.Age)
			fmt.Printf("学生的编号是%v\n", v.Address)
			fmt.Println()
		}
	*/
}
