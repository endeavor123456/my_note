# golang使用for语句实现while和do while语句功能

**golang中是没有while和do while语句的但是可以用for语句实现while和do while语句功能**

```go
//输出10次hello,world(使用类似while循环形式，先判断后做)
func jobWhileMoni() {//当然不用非得封装成函数 但是封装成函数会更方便使用
    var count = 0
    for {
        if count >= 10 {
            break //如果count>=10则退出
        }
        fmt.Println("hello,world", count)
        count++
    }
}
//模拟do……while实现输出10次hello,world（先做后判断）
func jobDowhileMoni(){
    var i = 0
    for{
        fmt.Println("hello,world",i)
        i++
        if(i>=10){
            break
        }
    }
}
```

