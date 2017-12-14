# 第二天内容
***

# Go运算符
***

## 关系运算符示例
`假设变量A的值为10，变量B的值为20`
| 运算符 | 描述 | 示例 |
| :---: | :-------------: | :---: |
| == | 检查两个操作数的值是否相等，如果相等，则条件为真。 | (A == B)结果为假 |
| != | 检查两个操作数的值是否相等，如果值不相等，则条件为真。 | (A != B)结果为真 |
| > | 检查左操作数的值是否大于右操作数的值，如果是，则条件为真。 | (A > B)结果为假 |
| < | 检查左操作数的值是否小于右操作数的值，如果是，则条件为真。 | (A < B)结果为真 |
| >= | 检查左操作数的值是否大于或等于右操作数的值，如果是，则条件为真。 | (A >= B)结果为假 |
| <= | 检查左操作数的值是否小于或等于右操作数的值，如果是，则条件为真。 | (A <= B)结果为真 |

## 逻辑运算符
`假设变量A的值为true，变量B的值为false`
| 运算符 | 描述 | 示例 |
| :---: | :-------------: | :---: |
| && | 逻辑AND运算符。如果两个操作数都为true，则条件为真，反之则假。 | (A && B)结果为假 |
| `||` | 逻辑OR运算符。如果两个操作数中的任何一个为true，则条件变为真。 | (A `||` B)结果为真 |
| ! | 逻辑非运算符。用于反转其操作数的逻辑状态。如果条件为真，则逻辑非运算符将为假。 | !(A && B)结果为真 |

## 位运算符
`假设变量A=60，并且变量B=13，A的二进制为00111100，B的二进制为00001101，`
| 运算符 | 描述 | 示例 |
| :---: | :-------------: | :-----: |
| &	| 如果两个操作数中都存在二进制AND运算符，则将其复制到结果。 | (A & B) 结果为 12, 也就是 0000 1100 |
| `|` | 二进制OR运算符复制一个位，如果它存在于任一操作数。 | (A `|` B) 结果为 61, 也就是 0011 1101 |
| ^	| 二进制XOR运算符复制位，如果它在一个操作数中设置，但不是在两个操作数中设置。 | (A ^ B) 结果为 49, 也就是 0011 0001 |
| << | 二进制左移位运算符。左操作数值向左移动由右操作数指定的位数。 | A << 2 结果为 240, 也就是 1111 0000 |
| >> | 二进制右移运算符。左操作数值向右移动由右操作数指定位数。 | A >> 2 结果为 15 , 也就是 0000 1111 |

## 赋值运算符
| 运算符 | 描述 | 示例 |
| :---: | :-------------: | :------: |
| = | 简单赋值操作符，将值从右侧操作数分配给左侧操作数 | C = A + B，就是将A + B的值赋给C |
| += | 相加和赋值运算符，向左操作数添加右操作数，并将结果赋给左操作数 | C += A 相当于 C = C + A |
| -= | 减去和赋值运算符，从左操作数中减去右操作数，并将结果赋给左操作数 | C -= A 相当于 C = C - A |
| *= | 乘法和赋值运算符，它将右操作数与左操作数相乘，并将结果赋给左操作数 | C *= A 相当于 C = C * A |
| /= | 除法和赋值运算符，它用右操作数划分左操作数，并将结果分配给左操作数 | C /= A 相当于 C = C / A |
| %= | 模数和赋值运算符，它使用两个操作数来取模，并将结果分配给左操作数 | C %= A 相当于 C = C % A |
| <<= | 左移和赋值运算符 | C << = 2相当于C = C << 2 |
| >>= | 右移和赋值运算符 | C >>= 2 相当于 C = C >> 2 |
| &= | 按位和赋值运算符 | C &= 2 相当于 C = C & 2 |
| ^= | 按位异或和赋值运算符 | C ^= 2 相当于 C = C ^ 2 |
| `|=` | 按位包含OR和赋值运算符 | C `|=` 2 相当于 C = C `|` 2 |

***

# Go条件与循环
***
## if语句
`if ; if...else ; if...else if...else`
```
if (a < b) {
  fmt.Println("a < b")
} else if (a == b) {
  fmt.Println("a == b")
} else {
  fmt.Println("a > b")
}
```

## switch语句
### switch有两种类型：
- 表达式开关(switch)： 在表达式开关(switch)中，case包含与开关(switch)表达式的值进行比较的表达式。
- 类型开关(switch)： 在类型开关(switch)中，case包含与特殊注释的开关(switch)表达式的类型进行比较的类型，进行比较的类型时，switch接收的参数必须是经interface{}初始化的参数。
```
	switch a {
		case 1:
			fmt.Println("a == 1")
		case 2:
			fmt.Println("a == 2")
		case 60:
			fmt.Println("a == 60")
		default:
			fmt.Println("not find")
	}

	switch {
		case a == 1:
			fmt.Println("a == 1")
		case a == 2:
			fmt.Println("a == 2")
		case a == 60:
			fmt.Println("a == 60")
		default:
			fmt.Println("not find")
	}

	var x interface{}	//switch类型判断，接收的参数必须是经interface{}初始化的参数
	x = a
	switch x.(type) {
		case nil:
			fmt.Println("a.type == nil")
		case int:
			fmt.Println("a.type == int")
		case func(int) float64:
			fmt.Println("a.type == func(int) float64")
		case bool, string:
			fmt.Println("a.type == bool or string")
		default:
			fmt.Println("not find a.type")
	}
```

## select语句
- 与switch较为类似，与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个`IO操作`，确切的说，应该是一个面向channel的IO操作，暂时不理解没关系，后面会详细解释
```
  var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		case c2 <- i2:
			fmt.Printf("sent ", i2, " to c2\n")
		case i3, ok := (<-c3):  // same as: i3, ok := <-c3
			if ok {
					fmt.Printf("received ", i3, " from c3\n")
			} else {
					fmt.Printf("c3 is closed\n")
			}
		 default:
				fmt.Printf("no communication\n")
	}
```

## for循环语句
- 语法：
```
for [ Condition(条件) |  ( init; condition; increment ) | Range(范围)] {
  //逻辑代码
}
```
## 循环控制语句
| 控制语句 | 描述 |
| :---: | :-------------: |
| break语句 | 终止for循环或switch语句，并将执行转移到for循环或switch之后的语句。 |
| continue语句 | 循环跳过其主体的其余部分，执行下一个循环。 |
| goto语句 | 将控制转移到带标签的语句 |

- Demo展示:
```
	var num, sum int = 5, 0;
	arr := [4]int{3, 2, 3} 

	// Condition(条件)
	Loop: for sum < 50 {
		if (sum >= 45) {
			sum++
			fmt.Printf("goto Loop => %d\n", sum);
			goto Loop					//goto语句提供了从goto到相同函数中的带标签语句的无条件跳转。
		}
		sum += num
	}
	fmt.Printf("(for sum < 50) result => %d \n", sum);

	// init; condition; increment
	sum = 0
	for i := 0; i < 10; i++ {
		if (sum == 20) {
			break						//跳出当前循环
		}
		sum += num;
	}
	fmt.Printf("(for i := 0; i < 10; i++) result => %d \n", sum);

	// Range(范围)
	sum = 0
	for j,v:= range arr {
		if (v == 0) {
			continue				//执行下一个循环
		}
		fmt.Printf("value of x = %d at %d\n", v, j)
		sum += v;
	}
	fmt.Printf("(for x := range arr) result => %d \n", sum);
```
***


# Go的函数详解
***
- 函数是一组执行任务的语句。 每个Go程序至少有一个函数，也就是main()函数，所有最小的程序可以定义附加的函数
### 函数形式：
```
func 函数名( [参数的列表] ) [返回类型列表]
{
   // 函数体
}
```
### 函数详解：
- `func`开始一个函数的声明。
- *函数名称：*这是函数的名称。函数名和参数列表一起构成函数签名。
- *参数：*参数类似于占位符。 调用函数时，将一个值传递给参数。 此值称为实际参数或参数。 参数列表指的是函数的参数的类型，顺序和数量。 参数是可选的; 即，函数可以不用包含参数。
- *返回类型：*函数可以有返回值列表。返回类型列表是函数返回的值的数据类型的列表。某些函数执行所需的操作，而不用(无)返回值。在这种情况下，返回类型就不是必需定义的。
- *函数体：*函数体包含定义函数的功能的语句集合。
- 函数名称如果是`以大写字母开头`的，则该函数被`导出`，可被`其他go文件程序调用`，类似于`Java`中的`public`，反之`小写开头的`则只能该Go文件`内部调用`，类似于`Java`中的`private`

### Demo(day1的代码中已经有使用):
```
//无返回类型
func job1() {
	//...
}

//单个参数单个返回类型
func job2(x string) string {
	return x
}

//多个参数多个返回类型
func job3(x, y int) (int, int) {
	return y, x
}
z, f := job3(x, y)

//命名返回值
func job4(x, y int) (r int) {
	r = x + y
	return
}
```

## 参数传递给函数的两种方式：
| 控制语句 | 描述 |
| :---: | :-------------: |
| `按值调用` | 此方法将参数的实际值复制到函数的形式参数中。在这种情况下，对函数中的参数所做的更改不会影响参数的值，退出函数后参数的值还是原值。 |
| `按引用调用` | 此方法将参数的地址复制到形式参数中。在函数内部，地址用于访问在调用中使用的实际参数。这意味着对参数所做的更改会影响参数的值，退出函数后参数的值是函数内部被修改后的值。 |
*`具体Demo关于引用调用会在后面讲指针的时候进行展示`*

## 函数闭包
- 闭包，即函数内部参数的外部引用
```
//闭包
func closure() func() int {
	var num int = 0
	return func() int {
		num++
		return num
	} 
}

numAdd := closure()
fmt.Printf("closure num is %d\n", numAdd())
fmt.Printf("closure num is %d\n", numAdd())
fmt.Printf("closure num is %d\n", numAdd())
```

## 函数方法
-	方法，也属于函数，相当与包含了接收者的函数，有点类似于`面向对象语言`中的`成员函数`
### Demo(demo中包含结构体的知识，会在后面详解):
```
//定义一个表示汽车的结构体
type Car struct {
	car_brand, car_color, car_type string
}

//方法（描述汽车）
func (car Car) carDescribe() string {
	return "这是一辆" + car.car_color + "的" + car.car_brand + car.car_type
}

car := Car{car_brand: "宝马", car_color: "白色", car_type: "SUV"}
fmt.Printf("函数方法调用demo： %s \n", car.carDescribe())	//函数方法调用demo结果： 这是一辆白色的宝马SUV
```
***

# Go作用域
***
## 作用域：
- *局部变量：* 在函数或块中声明的变量，只能由在该函数或该代码块中的语句调用
- *全局变量：* 在函数之外定义的变量，可以在程序定义的任何函数内访问到它们
```
package main

import "fmt"

//全局变量
var j int = 20

func main() {
	//局部变量
	var i int = 10

	fmt.Printf("i 是局部变量，值为： %d \n", i)
	fmt.Printf("j 是全局变量，值为： %d \n", j)
}
```
### *备注：* 如果`全局变量`以`大写字母开头`的，则该变量被`导出`，可被`其他go文件程序调用`，类似于`Java`中的`public`，反之`小写开头的`则只能该Go文件`内部调用`，类似于`Java`中的`private`
### *备注：* `局部变量`和`全局变量`的`名称`可以`相同`，但在`函数内部`的`局部变量`的值将`优先`与`全局变量`

***
## 