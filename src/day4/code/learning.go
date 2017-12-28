package code

import "fmt"
import "strconv"

func Output() {
	fmt.Println("********************* 切片相关 **********************")
	sliceFunc()

	fmt.Println("********************* 映射map相关 **********************")
	mapFunc()

	fmt.Println("********************* range相关 **********************")
	rangeFunc()

	fmt.Println("********************* 递归调用相关 **********************")
	recursionFunc()

	fmt.Println("********************* 类型转换相关 **********************")
	typeTransformFunc()
}

func sliceFunc() {
	testArr := [5]int{1, 2, 3, 4, 5}

	slice1 := []string{"f", "u", "c", "k"}
	slice11 := []string{"f1", "2u", "c3", "k"}
	slice2 := make([]string, 5, 5)
	slice3 := testArr[0:2]
	slice4 := testArr[:3]
	slice5 := testArr[2:]
	slice6 := slice1[0:3]

	fmt.Printf("数组：%v\n", testArr)
	fmt.Println("初始化：")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)

	fmt.Println("长度与容量：")
	fmt.Printf("slice1：%v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))


	slice1 = append(slice1, " ")
	fmt.Printf("添加一个元素之后，slice1：%v，长度：%d，容量：%d\n", slice1, len(slice1), cap(slice1))

	fmt.Printf("slice6：%v，长度：%d，容量：%d\n", slice6, len(slice6), cap(slice6))
	fmt.Printf("slice11：%v，长度：%d，容量：%d\n", slice6, len(slice11), cap(slice11))
	copy(slice6, slice11)
	fmt.Printf("copy之后，slice6：%v，长度：%d，容量：%d\n", slice6, len(slice6), cap(slice6))
}

func mapFunc() {
	mp1 := make(map[string]int)

	mp1["one"] = 1
	mp1["two"] = 2
	mp1["three"] = 3

	fmt.Printf("mp1 %s value is %d\n", "one", mp1["one"])
	fmt.Printf("mp1 %s value is %d\n", "two", mp1["two"])
	fmt.Printf("mp1 %s value is %d\n", "three", mp1["three"])

	var mp2 map[string]string
	mp2 = make(map[string]string)

	mp2["一"] = "one"
	mp2["二"] = "two"
	mp2["三"] = "three"

	fmt.Printf("mp2 %s value is %s\n", "一", mp2["一"])
	fmt.Printf("mp2 %s value is %s\n", "二", mp2["二"])
	fmt.Printf("mp2 %s value is %s\n", "三", mp2["三"])

	delete(mp2, "三")
	fmt.Printf("after delete %s, mp2 %s value is %v\n", "三", "三", mp2["三"])
}

func rangeFunc() {
	mp := make(map[string]string)
	mp["一"] = "one"
	mp["二"] = "two"
	mp["三"] = "three"
	for chNum, enNum := range mp {
		fmt.Printf("mp %s value is %s\n", chNum, enNum)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	for i, num := range arr {
		fmt.Printf("arr[%d] value is %d\n", i, num)
	}

	slice := arr[0:3]
	for i, num := range slice {
		fmt.Printf("slice[%d] value is %d\n", i, num)
	}

	fmt.Println("映射指针时：")

	map1 := make(map[int]*int)
	for i, num := range arr {
		map1[i] = &num
		fmt.Println(&num)
	}
	fmt.Println(map1)
	for key, value := range map1 {
		fmt.Printf("map1[%v] value is %v\n", key, *value)
	}

	map2 := make(map[int]*int)
	for i, num := range arr {
		data := num
		map2[i] = &data
		fmt.Println(&data)
	}
	fmt.Println(map2)
	for key, value := range map2 {
		fmt.Printf("map2[%v] value is %v\n", key, *value)
	}
}

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

func typeTransformFunc() {
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
	fmt.Println(e)
}