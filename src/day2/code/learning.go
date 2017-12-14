package code

import "fmt"

const a, b int = 60, 13
const A, B bool = true, false

//定义一个表示汽车的结构体
type Car struct {
	car_brand, car_color, car_type string
}

func Output() {
	fmt.Println("********************* 关系运算符: == != > < >= <= **********************")
	relationSymbol()

	fmt.Println("********************* 逻辑运算符: && || ! **********************")
	logicSymbol()

	fmt.Println("********************* 位运算符: & | ^ << >> **********************")
	bitSymbol()

	fmt.Println("********************* if条件语句: if else else if **********************")
	ifFunc()

	fmt.Println("********************* switch语句 **********************")
	switchFunc()

	fmt.Println("********************* select语句 **********************")
	selectFunc()

	fmt.Println("********************* for语句: Condition(条件) |  ( init; condition; increment ) | Range(范围) **********************")
	forFunc()

	fmt.Println("********************* 函数闭包 **********************")
	numAdd := closure()
	fmt.Printf("closure num is %d\n", numAdd())
	fmt.Printf("closure num is %d\n", numAdd())
	fmt.Printf("closure num is %d\n", numAdd())

	fmt.Println("********************* 函数方法 **********************")
	//初始化结构体
	car := Car{car_brand: "宝马", car_color: "白色", car_type: "SUV"}
	fmt.Printf("函数方法调用demo： %s \n", car.carDescribe())
}

//关系运算符
func relationSymbol() {	
	fmt.Printf("%d == %d : %t\n", a, b, a == b)
	fmt.Printf("%d != %d : %t\n", a, b, a != b)
	fmt.Printf("%d > %d : %t\n", a, b, a > b)
	fmt.Printf("%d < %d : %t\n", a, b, a < b)
	fmt.Printf("%d >= %d : %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d : %t\n", a, b, a <= b)
}

//逻辑运算符
func logicSymbol() {
	fmt.Printf("%t && %t : %t\n", A, B, A && B)
	fmt.Printf("%t || %t : %t\n", A, B, A || B)
	fmt.Printf("!(%t && %t) : %t\n", A, B, !(A && B))
}

//位运算符
func bitSymbol() {
	fmt.Printf("%d(%b) & %d(%b) : %d(%b)\n", a, a, b, b, a & b, a & b)
	fmt.Printf("%d(%b) | %d(%b) : %d(%b)\n", a, a, b, b, a | b, a | b)
	fmt.Printf("%d(%b) ^ %d(%b) : %d(%b)\n", a, a, b, b, a ^ b, a ^ b)
	fmt.Printf("%d(%b) << 1 : %d(%b)\n", a, a, a << 1, a << 1)
	fmt.Printf("%d(%b) >> 1 : %d(%b)\n", a, a, a >> 1, a >> 1)
}

//if条件语句
func ifFunc() {
	if (a < b) {
		fmt.Println("a < b")
	} else if (a == b) {
		fmt.Println("a == b")
	} else {
		fmt.Println("a > b")
	}
}

//switch语句
func switchFunc() {
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
}

//select语句
func selectFunc() {
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
}

//for循环语句
func forFunc() {
	var num, sum int = 5, 0;

	arr := [7]int{3, 2, 3} 

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
}

//函数闭包
func closure() func() int {
	var num int = 0
	return func() int {
		num++
		return num
	} 
}


//函数方法（描述汽车）
func (car Car) carDescribe() string {
	return "这是一辆" + car.car_color + "的" + car.car_brand + car.car_type
}