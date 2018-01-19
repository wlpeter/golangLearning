package code

import "fmt"

func Output() {
	fmt.Println("********************* 接口相关 **********************")
	interfaceFunc()
}

//狗接口
type Doger interface {
	//狗叫
	say() string
}

//泰迪犬
type Taidi struct {
	voice, name string
}

//斗牛犬
type Douniu struct {
	voice, name string
}

//泰迪犬叫
func (t Taidi) say() string {
	return "泰迪犬 " + t.name + " say " + t.voice
}

//斗牛犬叫
func (d Douniu) say() string {
	return "斗牛犬 " + d.name + " say " + d.voice
}

//打印狗叫
func dogSpeakPrint(dog Doger) {
	fmt.Println(dog.say())
}

type Simpler interface {
	Set(i int)
	Get() int
}

type Simple struct {
	num int
}

func (s *Simple) Set(i int) {
	s.num = i
}

func (s *Simple) Get() int {
	return s.num
}

//狗品种类型
func sayDogType(dog Doger) string {
	dogType := ""
	switch 	dog.(type) {
		case Douniu:
			dogType = "斗牛犬"
		case *Douniu:
			dogType = "*斗牛犬"
		case Taidi:
			dogType = "泰迪犬"
		case *Taidi:
			dogType = "*泰迪犬"
		default:
			dogType = "未知生物"
	}
	return dogType
}

//接口相关Demo
func interfaceFunc() {
	//泰迪小白
	taidi1 := Taidi{voice: "汪汪", name: "小白"}

	//泰迪小黄
	taidi2 := new(Taidi)
	taidi2.name = "小黄"
	taidi2.voice = "喔喔"

	//斗牛犬小白
	douniu := Douniu{voice: "旺旺", name: "小黑"}

	taidiList := []Doger{ taidi1, taidi2, douniu }
	for i, _ := range taidiList {
		dogSpeakPrint(taidiList[i])
	}

	simple1 := new(Simple)
	simple1.Set(2)
	fmt.Println(simple1.Get())
	simple1.Set(3)
	fmt.Println(simple1.Get())
	fmt.Println(simple1.num)

	var simplerIn Simpler
	simplerIn = simple1
	if v, ok := simplerIn.(*Simple); ok {
    fmt.Printf("Simple结构体 是 Simpler接口 的动态类型(%T)\n", v)
	} else {
		fmt.Printf("Simple结构体 不是 Simpler接口 的动态类型\n")
	}

	fmt.Printf("%s 是 %s \n", taidi1.name, sayDogType(taidi1))
	fmt.Printf("%s 是 %s \n", taidi2.name, sayDogType(taidi2))
	fmt.Printf("%s 是 %s \n", douniu.name, sayDogType(douniu))
}

