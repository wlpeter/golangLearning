# 第三天内容
***
*所有详细代码都能在code中有展示，可查看*

# Go字符串类型与派生类型
***

## 1、字符串类型
- 表示字符的集合。它的值是一个字节序列，字符串的编码格式为*UTF-8*
```
var str string = "hello"
```
- `len()`可用来计算字符串的*字节*长度
- `strings.Count(str, "") - 1`可用来计算字符串的*字符*长度
- 大部分中文字在*utf-8*的编码格式下，一个*中文字*表示*3个字节*，而*英文数字及相关符号*表示*1个字节*
- `strings.Replace(s, old, new string, n int)`可用来替换字符串，当n < 0时，就替换所有，当n > 0时，就替换n个字符
- 更多的关于*字符串类型*的*操作*可以参考`strings`标准库文档[进入入口](http://docscn.studygolang.com/pkg/strings/)
```
import "strings"

//....

str := "hello world!"
fmt.Printf("%s\n", str);                                      //hello world!

// 字节长度
fmt.Printf("hello world! byte length is : %d\n", len(str))    //hello world! length is : 12
str = "hello world!\n"
fmt.Printf("hello world!\\n byte length is : %d\n", len(str)) //hello world!\n length is : 13
str = "中国"
fmt.Printf("中国 byte length is : %d\n", len(str))            //中国 byte length is : 6

// 字符长度
fmt.Printf("中国 character length is : %d\n", strings.Count(str, "") - 1) //中国 character length is : 2

// 字符串拼接
str1, str2, str3, str4 := "f", "u", "c", "k"
s := []string{str1, str2, str3, str4}                                      //字符串数组
fmt.Printf("字符串传拼接(strings.Join) : %s\n", strings.Join(s, ""))        //字符串传拼接(strings.Join) : fuck
fmt.Printf("字符串传拼接(strings.Join) : %s\n", strings.Join(s, ","))       //字符串传拼接(strings.Join) : f,u,c,k
fmt.Printf("字符串传拼接(+) : %s\n", str1 + str2 + str3 + str4)             //字符串传拼接(+) : fuck

// 替换
fmt.Println(strings.Replace("oink oink oink", "o", "ky", 1))              //kyink oink oink
fmt.Println(strings.Replace("oink oink oink", "o", "ky", -1))             //kyink kyink kyink
```

## 2、数组类型
- 用来存储类型相同元素且有固定顺序的数据集合

```
// 声明
var 变量名 [数组长度] 数组元素类型

// 赋值
变量名 = [数组长度]数组元素类型{ 元素参数1, 元素参数2 }
```
- *备注：*大括号`{}`中的值的数量不能大于在方括号`[]`中为数组声明指定的元素数量，如果`[]`中省略数组的大小，则只创建一个足够容纳初始化的数组

demo:
```
var list [] string = []string{"s", "q", "f"}
var list2 = [3]string{"t", "q", "z"}
list3 := []string{"t", "q", "z"}

//循环遍历数组
for i := 0; i < len(list); i++ {
  fmt.Printf("list[%d] = %s\n", i, list[i] )
}
```

*当数组作为参数传入函数时，有两种情况：*
- 形式参数作为一个未知道大小的数组时，`func myFunction(param []int)`
- 形式参数作为一个已知大小的数组时，`func myFunction(param [5]int)`

### 2.1、多维数组 
```
// 声明
var 变量名 [数组1长度][数组2长度]...[数组n长度] 数组元素类型

// 赋值
变量名 = [数组长度][数组2长度]数组元素类型{ 
  { 数组1元素1, 数组1元素2, 数组1元素3... },
  { 数组2元素1, 数组2元素2, 数组2元素3... }
  ...
}
```
多维数组demo: 
```
list := [][]string{ {"t1", "q1", "z1"}, {"t2", "q2", "z2"} }

for i:=0; i < len(list); i++ {
  for j:=0; j < len(list[i]); j++ {
    fmt.Printf("list3[%d][%d] = %s\n", i, j, list[i][j] )
  }
}
```

## 3、指针类型

- 学过C/C++的，应该对指针这个概念比较了解，指针在C中是一大神器
- 用来存储其他变量内存地址的变量，使用方法：要先声明一个指针，然后才能使用它来存储任何变量地址
- 通常指针都用星号（`*`）表示，用（`&`）来获取地址
```
var 指针变量名 *指针的基类型
```

demo:
```
var str string = "hello"
var address *string = &str
var address1 *string

fmt.Printf("Address of num variable: %x\n", &str)                 //Address of num variable: c42000e300
fmt.Printf("Address stored in address variable: %x\n", address)   //Address stored in address variable: c42000e300
fmt.Printf("value of address variable: %s\n", *address)           //value of address variable: hello
fmt.Println(address != nil)                                       //true
fmt.Println(address1 == nil)                                      //true
```
- 指针创建时会首先指向`nil`值，以防止无地址分配

顺便看下函数`值调用`和`引用调用`的Demo:
```
var str string = "hello"
var address *string = &str

//引用调用
referenceFunc(address)
fmt.Printf("after reference, value of address variable: %s\n", *address)      //after reference, value of address variable: fuck you
fmt.Printf("after reference, value of str variable: %s\n", str)               //after reference, value of str variable: fuck you

//值调用
valueFunc(str)
fmt.Printf("after valueFunc, value of str variable: %s\n", str)               //after valueFunc, value of str variable: fuck you

...

//引用调用
func referenceFunc(d *string) {
	*d = "fuck you"
}

//值调用
func valueFunc(d string) {
	d = "hello world"
}
```

### 3.1、数组指针
```
//指针数组
list := []int{1, 2, 3}
var alist [3]*int
for i:=0; i< len(list); i++ {
  alist[i] = &list[i]
}

for i:=0; i< len(alist); i++ {
  fmt.Printf("Value of list[%d] = %d\n", i,*alist[i] )
}
```

### 3.2、指针的指针
- 顾名思义，存储其他指针的地址的指针
- `var s **int`，可以拆分着看，这样看比较容易懂，分为`*`和`*int`，第一个*代表指针*，第二个代表*指针类型*
```
var b string = "come on"
var ip *string = &b
var ipp **string = &ip

fmt.Printf("Value of b = %s\n", b )                     //Value of b = come on
fmt.Printf("Value available at *ptr = %s\n", *ip )      //Value available at *ptr = come on
fmt.Printf("Value available at **pptr = %s\n", **ipp)   //Value available at **pptr = come on
```

## 4、结构体
- 学过C/C++的应该对这个概念比较了解，结构体是一个可以在其内部*定义多个不同类型变量*的变量类型，有点点类似于*面向对象语言*中的*类*
```
type 结构体名称 struct {
   成员变量1 变量类型1
   成员变量2 变量类型2
   ...
   成员变量n 变量类型n
}
```
demo：
```

//人（结构体）
type People struct {
	name, sex string
	age int
}

func structFunc() {
	var people1, people2 People

	// 赋值方法1
	people1 = People{ name: "张三", sex: "男性", age: 25 }

	// 赋值方法2
	people2.name = "李四"
	people2.sex = "女性"
	people2.age = 23

	fmt.Printf( "people1 name : %s\n", people1.name)
	fmt.Printf( "people1 sex : %s\n", people1.sex)
	fmt.Printf( "people1 age : %d\n", people1.age)

	fmt.Printf( "people2 name : %s\n", people2.name)
	fmt.Printf( "people2 sex : %s\n", people2.sex)
	fmt.Printf( "people2 age : %d\n", people2.age)
}
```

### 4.1、结构体与函数方法的使用
- 函数方法可接收结构体内部参数进行调用
Demo(demo中使用了`string()`进行了数据类型转换):
```
//人（结构体）
type People struct {
	name, sex string
	age int
}

//结构体方法
func (p People)peopleDesc() string {
	return p.name + "是一名" + string(p.age) + "的" + p.sex
}

func main() {
  var people2 People

  people2.name = "李四"
	people2.sex = "女性"
	people2.age = 23

  fmt.Println(people2.peopleDesc())
}
```
***
