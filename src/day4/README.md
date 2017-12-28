# 第四天内容
***
*所有详细代码都能在code中有展示，可查看*

# Go切片
***

## 切片概念
- 切片是有数组构成的，它的**底层**就是**数组**，但不同于数组的是，切片的长度是不定的可变的。
- 切片有两大概念：**len长度**和**cap容量**，**长度**是指切片内**当前容纳**的**元素个数**，**容量**是指切片目前可容纳的**最多元素个数**，当给切片添加元素时，如果**长度超出容量**的时候，Go会再创建一个容量为当前容量**两倍**的新数组，并将数据拷贝进去

**初始化方式：**
```
s :=[] int {1,2,3} 
```
- 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

```
s := arr[:] 
```
- 初始化切片s,是数组arr的引用

```
s := arr[startIndex:endIndex] 
```
- 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

```
s := arr[startIndex:] 
```
- 缺省endIndex时将表示一直到arr的最后一个元素

```
s := arr[:endIndex] 
```
- 缺省startIndex时将表示从arr的第一个元素开始

```
s1 := s[startIndex:endIndex] 
```
- 通过切片s初始化切片s1

```
s :=make([]int,len,cap) 
```
- 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片



Demo:
```
slice1 := []string{"f", "u", "c", "k"}
fmt.Printf("slice1：%v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))
```

copy()和append():
- `append()`：可用来添加切片内部元素或者合并切片
- `copy()`：可将源切片的内容**复制到**目标切片，但**不会改变**目标切片的**长度以及容器**，**除非**目标切片为**Nil切片（即未做初始化）**


Demo：
```
slice1 := []string{"f", "u", "c", "k"}
slice6 := slice1[0:3]
slice11 := []string{"f1", "2u", "c3", "k"}


fmt.Printf("slice1：%v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))
slice1 = append(slice1, " ")
fmt.Printf("添加一个元素之后，slice1：%v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))

fmt.Printf("slice6：%v，长度：%d，容量：%d\n", slice6, len(slice6), cap(slice6))
fmt.Printf("slice11：%v，长度：%d，容量：%d\n", slice6, len(slice11), cap(slice11))
copy(slice6, slice11)
fmt.Printf("copy之后，slice6：%v，长度：%d，容量：%d\n", slice6, len(slice6), cap(slice6))
```

***

# Go映射（Map）
***
- Go中的字典或哈希类型，是一种键值映射的数据类型

demo: 
```
//声明方式1
mp1 := make(map[string]int)

//声明方式2
var mp2 map[string]string
mp2 = make(map[string]string)

mp1["one"] = 1
mp1["two"] = 2
mp1["three"] = 3

fmt.Printf("mp1 %s value is %d\n", "one", mp1["one"])
fmt.Printf("mp1 %s value is %d\n", "two", mp1["two"])
fmt.Printf("mp1 %s value is %d\n", "three", mp1["three"])

mp2["一"] = "one"
mp2["二"] = "two"
mp2["三"] = "three"

fmt.Printf("mp2 %s value is %s\n", "一", mp2["一"])
fmt.Printf("mp2 %s value is %s\n", "二", mp2["二"])
fmt.Printf("mp2 %s value is %s\n", "三", mp2["三"])

//删除键值
delete(mp2, "三")
fmt.Printf("after delete %s, mp2 %s value is %v\n", "三", "三", mp2["三"])
```
***


# Go范围（range）
***
- `range`可用来遍历数组、切片和映射，使用数组和切片，它返回项的索引为整数。使用映射则它返回下一个键值对的键。 

demo: 
```
//映射
mp := make(map[string]string)
mp["一"] = "one"
mp["二"] = "two"
mp["三"] = "three"
for chNum, enNum := range mp {
  fmt.Printf("mp %s value is %s\n", chNum, enNum)
}

//数组
arr := [5]int{1, 2, 3, 4, 5}
for i, num := range arr {
  fmt.Printf("arr[%d] value is %d\n", i, num)
}

//切片
slice := arr[0:3]
for i, num := range slice {
  fmt.Printf("slice[%d] value is %d\n", i, num)
}
```
**这里讲一个问题，demo如下：**
```
//数组
arr := [5]int{1, 2, 3, 4, 5}

map1 := make(map[int]*int)
for i, num := range arr {
  map1[i] = &num
}

for key, value := range map1 {
  fmt.Printf("map1[%v] value is %v\n", key, *value)
}
```

运行程序输出如下：
```
map1[0] value is 5
map1[1] value is 5
map1[2] value is 5
map1[3] value is 5
map1[4] value is 5
```

**与预期想的1，2，3，4，5不同，原因是：**
- for range时会把每个元素的值赋给代码的`num`中，然后`map1[i]`的指针指向`num`的地址，然而`num`只创建了一次，所以每个循环`num`的地址是不会变的，导致`map1`中保存指针全部指向`num`的地址，而最后`num`的地址所保存的至为5，所以导致了上述结果，
- 只要赋值是重新创建一个变量即可，这样每次赋值所指向的地址都是不同的

**正确方法如下：**

```
//数组
arr := [5]int{1, 2, 3, 4, 5}

map2 := make(map[int]*int)
for i, num := range arr {
  data := num
  map2[i] = &data
}

for key, value := range map2 {
		fmt.Printf("map2[%v] value is %v\n", key, *value)
	}
```

运行程序输出如下：
```
map2[0] value is 1
map2[1] value is 2
map2[2] value is 3
map2[3] value is 4
map2[4] value is 5
```
***

# Go递归
***
- Go语言与大部分的编程语言相同都支持递归调用，需要注意在函数中定义或设置一个退出条件，否则它会进入无限循环。

直接上代码，直接以*斐波那契数列*为Demo：
```
func recursionFunc() {
	fmt.Println("斐波那契数列: ")
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", feibonaqi(i))
	}
	fmt.Println("")
}

// 斐波那契数列
func feibonaqi(i int)int {
	if (i == 0 || i == 1) {
		return i
	}
	return feibonaqi(i-1) + feibonaqi(i-2)
}
```
***

# Go类型转换
***
- 是一种将变量从一种数据类型转换为另一种数据类型的方法
- 转换类型可通过**类型转换符**或者**strconv函数库**来进行，**strconv函数库**的功能更为强大

demo: 
```
var num1 int = 1
var num2 float32 = 2.2
var num3 string = "3"
var num4 int = 4
//string转换int
num3int, e := strconv.Atoi(num3)
//int转换string
num4str := strconv.Itoa(num4)
fmt.Printf("num1 type is %T\n", num1)
fmt.Printf("num2 type is %T\n", num2)
fmt.Printf("num3 type is %T\n", num3)
fmt.Printf("num4 type is %T\n", num4)
fmt.Printf("afte change type, num1 type is %T\n", float32(num1))
fmt.Printf("afte change type, num2 type is %T\n", int(num2))
fmt.Printf("afte change type, num3 type is %T\n", num3int)
fmt.Printf("afte change type, num4 type is %T\n", num4str)
```
***
