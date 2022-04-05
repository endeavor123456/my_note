## golang中的比较注意事项

大家经常用比较运算符来比较两个变量。但是golang中的比较运算符有很多细节的地方，跟php是不一样的。很多时候不能直接用比较运算符来比较，编译器会直接报错。

## 一，基本类型的变量比较

golang中的基本类型

比较的两个变量类型必须相等。而且，golang没有隐式类型转换，比较的两个变量必须类型完全一样，类型别名也不行。如果要比较，先做类型转换再比较。

- 类型完全不一样的，不能比较
- `类型再定义关系，不能比较，可以强转之后再比较`
- `类型别名关系，可以比较`   类型名和类型别名是不一样的  type Int=int  var a Int   变量a还是int类型  类型别名是给类型名取一个别名 所以不属于用户用type关键字定义的类型



```go
    fmt.Println("2" == 2) //invalid operation: "2" == 2 (mismatched types string and int)

    type A int
    var a int = 1
    var b A = 1
    fmt.Println(a == b) //invalid operation: a == b (mismatched types int and A)
    fmt.Println(a == int(b)) //true

    type C = int
    var c C = 1
    fmt.Println(a == c) //true
```

## 二，复合类型的变量比较

复合类型是逐个字段，逐个元素比较的。需要注意的是，`array 或者struct中每个元素必须要是可比较的，如果某个array的元素 or struct的成员不能比较(比如是后面介绍的slice，map等)，则此复合类型也不能比较。`

#### 1，数组类型变量比较

- 数组的长度是类型的一部分，如果数组长度不同，无法比较
- 逐个元素比较类型和值。每个对应元素的比较遵循基本类型变量的比较规则。跟struct一样，如果item是不可比较的类型，则array也不能做比较。

#### 2，struct类型变量比较

逐个成员(字段)比较类型和值。每个对应成员的比较遵循基本类型变量的比较规则。



```go
    type Student struct {
        Name string
        Age  int
    }

    a := Student{"minping", 30}
    b := Student{"minping", 30}
    fmt.Println(a == b)   //true
    fmt.Println(&a == &b) //false
```

但是如果struct中有不可比较的成员类型时：



```go
type Student struct {
        Name string
        Age  int
        Info []string
    }

    a := Student{
        Name: "minping",
        Age:  30,
    }
    b := Student{
        Name: "minping",
        Age:  30,
    }
    fmt.Println(a == b)   //invalid operation: a == b (struct containing []string cannot be compared)
```

可以看到，struct中有slice这种不可比较的成员时，整个struct都不能做比较，即使没有对slice那个成员赋值(slice默认值为nil)也不能进行比较

## 三，引用类型的变量比较

slice和map的比较规则比较奇怪，我们先说普通的变量引用类型&val和channel的比较规则。

#### 1，指针类型的比较规则

指针类型变量的值是一个空间的地址。所以指针类型变量的比较，判断的是这两个指针类型存储的值是不是相同。

- 如果这两个指针类型的变量的值相等，用"=="判断为true              变量a == 变量b 如果该表达式的结果为true<=>变量a的值等于变量b的值
- 如果不是同一个地址，则这两个指针类型的变量的值不相等，"=="结果为false

#### 2，slice这种引用类型的比较

```
slice类型不可比较，只能与nil做比较。
```

```go
    a := []string{}
    b := []string{}
    fmt.Println(a == b)  //error：invalid operation: a == b (slice can only be compared to nil)
```

关于slice类型不可比较的原因，后面会专门写文章做讨论。

#### 3， map类型的比较

```
map类型和slice一样，不能比较，只能与nil做比较。
```

## 四，interface{}类型变量的比较

接口类型的变量，包含该接口变量存储的值和值的类型两部分组成，分别称为接口的动态类型和动态值。`只有动态类型和动态值都相同时，两个接口变量才相同:`



```go
type Person interface {
    getName() string
}

type Student struct {
    Name string
}

type Teacher struct {
    Name string
}

func (s Student) getName() string {
    return s.Name
}

func (t Teacher) getName() string {
    return t.Name
}

func compare(s, t Person) bool {
    return s == t
}

func main() {

    s1 := Student{"minping"}
    s2 := Student{"minping"}
    t := Teacher{"minping"}

    fmt.Println(compare(s1, s2)) //true
    fmt.Println(compare(s1, t))  //false,类型不同
}
```

而且接口的动态类型必须要是可比较的，如果不能比较(比如slice，map)，则运行时会报panic。因为编译器在编译时无法获取接口的动态类型，所以编译能通过，但是运行时直接panic:



```go
type Person interface {
    getName() string
}

type Student map[string]string

type Teacher map[string]string

func (s Student) getName() string {
    return s["name"]
}

func (t Teacher) getName() string {
    return t["name"]
}

func compare(s, t Person) bool {
    return s == t
}

func main() {

    s1 := Student{}
    s1["name"] = "minping"
    s2 := Student{}
    s2["name"] = "minping"

    fmt.Println(compare(s1, s2)) //runtime error: comparing uncomparable type main.Student

}
```

## 五，函数类型的比较

golang的func作为一等公民，也是一种类型，而且不可比较



```go
f := func(int) int { return 1 }
g := func(int) int { return 2 }
f == g
```

## 六，slice和map的特殊比较

上面说过，map和slice是不可比较类型，但是有没有特殊的方法来对slice和map做比较呢，有

#### 1，[]byte类型的变量，使用工具包byte提供的函数就可以做比较



```go
s1 := []byte{'f', 'o', 'o'}
s2 := []byte{'f', 'o', 'o'}
fmt.Println(bytes.Equal(s1, s2)) // true
s2 = []byte{'b', 'a', 'r'}
fmt.Println(bytes.Equal(s1, s2)) // false
s2 = []byte{'f', 'O', 'O'}
fmt.Println(bytes.EqualFold(s1, s2)) // true
s1 = []byte("źdźbło")
s2 = []byte("źdŹbŁO")
fmt.Println(bytes.EqualFold(s1, s2)) // true
s1 = []byte{}
s2 = nil
fmt.Println(bytes.Equal(s1, s2)) // true
```

#### 2，使用反射

reflect.DeepEqual函数可以用来比较两个任意类型的变量



```go
func DeepEqual(x, y interface{})
```

对map类型做比较：



```go
m1 := map[string]int{"foo": 1, "bar": 2}
m2 := map[string]int{"foo": 1, "bar": 2}
// fmt.Println(m1 == m2) // map can only be compared to nil
fmt.Println(reflect.DeepEqual(m1, m2)) // true
m2 = map[string]int{"foo": 1, "bar": 3}
fmt.Println(reflect.DeepEqual(m1, m2)) // false
m3 := map[string]interface{}{"foo": [2]int{1,2}}
m4 := map[string]interface{}{"foo": [2]int{1,2}}
fmt.Println(reflect.DeepEqual(m3, m4)) // true
var m5 map[float64]string
fmt.Println(reflect.DeepEqual(m5, nil)) // false
fmt.Println(m5 == nil) // true
```

对slice类型做比较：



```go
s := []string{"foo"}
fmt.Println(reflect.DeepEqual(s, []string{"foo"})) // true
fmt.Println(reflect.DeepEqual(s, []string{"bar"})) // false
s = nil
fmt.Println(reflect.DeepEqual(s, []string{})) // false
s = []string{}
fmt.Println(reflect.DeepEqual(s, []string{})) // true
```

对struct类型做比较：



```go
type T struct {
    name string
    Age  int
}
func main() {
    t := T{"foo", 10}
    fmt.Println(reflect.DeepEqual(t, T{"bar", 20})) // false
    fmt.Println(reflect.DeepEqual(t, T{"bar", 10})) // false
    fmt.Println(reflect.DeepEqual(t, T{"foo", 10})) // true
}
```

可以发现，只要变量的类型和值相同的话，reflect.DeepEqual比较的结果就为true

#### 2，使用google的cmp包

直接看用例：



```go
import (
    "fmt"
    "github.com/google/go-cmp/cmp"
)
type T struct {
    Name string
    Age  int
    City string
}
func main() {
    x := T{"Michał", 99, "London"}
    y := T{"Adam", 88, "London"}
    if diff := cmp.Diff(x, y); diff != "" {
        fmt.Println(diff)
    }
}
```

结果为：



```bash
 main.T{
-       Name: "Michał",
+       Name: "Adam",
-       Age:  99,
+       Age:  88,
        City: "London",
  }
```

## 五，总结

- 1，复合类型，只有每个元素(成员)可比较，而且类型和值都相等时，两个复合元素才相等
- 2，slice，map不可比较，但是可以用reflect或者cmp包来比较
- 3，func作为golnag的一等公民，也是一个类型，也不能比较。
- 4，引用类型的比较是看指向的是不是同一个变量
- 5，类型再定义(type A string)不可比较，是两种不同的类型
- 6，类型别名(type A = string)可比较，是同一种类型。

## 六，拓展知识

1， [golang的类型再定义和类型别名](https://www.jianshu.com/p/3f5b66aa216e)
 2，golang的slice和map为什么不可以比较



作者：舒小贱
链接：https://www.jianshu.com/p/a982807819fa
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。