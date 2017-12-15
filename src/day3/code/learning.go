package code

import "fmt"
import "strings"

func Output() {
	fmt.Println("********************* 字符串相关 **********************")
	stringfunc()

	fmt.Println("********************* 数组相关 **********************")
	listFunc()

	fmt.Println("********************* 指针相关 **********************")
	pointerFunc()

	fmt.Println("********************* 结构体相关 **********************")
	structFunc()
}

//字符串
func stringfunc () {
	str := "hello world!"
	fmt.Printf("%s\n", str);

	// 字节长度
	fmt.Printf("hello world! byte length is : %d\n", len(str))
	str = "hello world!\n"
	fmt.Printf("hello world!\\n byte length is : %d\n", len(str))
	str = "中国"
	fmt.Printf("中国 byte length is : %d\n", len(str))

	// 字符长度
	fmt.Printf("中国 character length is : %d\n", strings.Count(str, "") - 1)

	// 拼接
	str1, str2, str3, str4 := "f", "u", "c", "k"
	s := []string{str1, str2, str3, str4}
	fmt.Printf("字符串传拼接(strings.Join) : %s\n", strings.Join(s, ""))
	fmt.Printf("字符串传拼接(strings.Join) : %s\n", strings.Join(s, ","))
	fmt.Printf("字符串传拼接(+) : %s\n", str1 + str2 + str3 + str4)

	// 替换
	fmt.Println(strings.Replace("oink oink oink", "o", "ky", 1))
	fmt.Println(strings.Replace("oink oink oink", "o", "ky", -1))
}


//数组
func listFunc () {
	var list [] string = []string{"s", "q", "f", "y"}
	var list2 = [3]string{"t", "q", "z"}
	list3 := [][]string{ {"t1", "q1", "z1"}, {"t2", "q2", "z2"} }

	loopJob1(list)
	loopJob2(list2)
	loopJob3(list3)
}

//未知大小的数组作为参数
func loopJob1(list []string) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i] )
	}
}

//已知大小的数组作为参数
func loopJob2(list [3]string) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("list2[%d] = %s\n", i, list[i] )
	}
}

//二维数组作为参数
func loopJob3(list [][]string) {
	for i:=0; i < len(list); i++ {
		for j:=0; j < len(list[i]); j++ {
			fmt.Printf("list3[%d][%d] = %s\n", i, j, list[i][j] )
		}
	}
}

//指针
func pointerFunc() {
	var str string = "hello"
	var address *string = &str
	var address1 *string

	fmt.Printf("Address of num variable: %x\n", &str)
	fmt.Printf("Address stored in address variable: %x\n", address)
	fmt.Printf("value of address variable: %s\n", *address)
	fmt.Println(address != nil)
	fmt.Println(address1 == nil)


	//引用调用
	referenceFunc(address)
	fmt.Printf("after reference, value of address variable: %s\n", *address)
	fmt.Printf("after reference, value of str variable: %s\n", str)

	//值调用
	valueFunc(str)
	fmt.Printf("after valueFunc, value of str variable: %s\n", str)


	//指针数组
	list := []int{1, 2, 3}
	var alist [3]*int
	for i:=0; i< len(list); i++ {
		alist[i] = &list[i]
	}

	for i:=0; i< len(alist); i++ {
		fmt.Printf("Value of list[%d] = %d\n", i,*alist[i] )
	}

	//指针的指针
	var b string = "come on"
	var ip *string = &b
	var ipp **string = &ip

	fmt.Printf("Value of b = %s\n", b )
	fmt.Printf("Value available at *ptr = %s\n", *ip )
	fmt.Printf("Value available at **pptr = %s\n", **ipp)
}

//引用调用
func referenceFunc(d *string) {
	*d = "fuck you"
}

//值调用
func valueFunc(d string) {
	d = "hello world"
}


//人（结构体）
type People struct {
	name, sex string
	age int
}

//结构体
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

	desc := people2.peopleDesc()
	fmt.Println(desc)
}

//结构体方法
func (p People)peopleDesc() string {
	return p.name + "是一名" + string(p.age) + "的" + p.sex
}