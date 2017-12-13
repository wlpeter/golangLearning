package helloWorld

import "fmt"

func Output() {
	var r1, r2 int
	r1, r2 = funcManyResult(5, 2)
	fmt.Println("********************* 多值返回:funcManyResult(x, y int) (int, int) **********************")
	fmt.Printf("funcManyResult: 5 + 2 = %d . \n", r1)
	fmt.Printf("funcManyResult: 5 - 2 = %d . \n", r2)

	fmt.Println("********************* 函数中命名返回值:funcResultNamed() (r string) **********************")
	fmt.Printf("%s . \n", funcResultNamed())

	fmt.Println("********************* 静态类型声明: var x float64 = 20.0 **********************")
	paramStatic()

	fmt.Println("********************* 动态类型声明: x := 20; var y = 23.3 **********************")
	paramActive()

	fmt.Println("********************* 混合变量声明: var x, y, s = 23.3, 1, \"seaeaeda\" **********************")
	paramBlend()

	fmt.Println("********************* 常量声明: const WIDTH int = 10; const LENGTH = 20 **********************")
	constParam()

	fmt.Println("********************* 算术运算符: ++ -- **********************")
	operation()
}

// 多值返回
func funcManyResult(x, y int) (int, int){
	return (x + y), (x - y)
}

// 函数中命名返回值
func funcResultNamed() (r string){
	r = "函数中命名返回值为r"
	return
}

//静态类型声明
func paramStatic() {
	var x float64
	x = 20
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
}

//动态类型声明
func paramActive() {
	x := 20
	var y = 23.3
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
}

//混合变量声明
func paramBlend() {
	var x, y, s = 23.3, 1, "seaeaeda"
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("s is of type %T\n", s)
}

//常量声明
func constParam() {
	const LENGTH = 10
	const WIDTH int = 20
	fmt.Printf("LENGTH const ：%d\n", LENGTH)
	fmt.Printf("WIDTH const ：%d\n", WIDTH)
}

//常量声明
func operation() {
	var a int = 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}