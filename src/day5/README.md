# 第五天内容
***
*所有详细代码都能在code中有展示，可查看*

# Go接口类型
***
- GO语言没有类和继承的概念，但是任然有**接口**这个概念，接口类型是对其它类型行为的抽象和概括
- Go中的接口与其他语言的接口比较相似，接口定义了一组方法（方法集），但是这些方法不包含（实现）代码即**抽象**的，接口里也**不能包含变量**。

定义接口格式：
```
type 接口名称 interface {
    方法名称1(方法参数集) 参数返回类型
    方法名称2(方法参数集) 参数返回类型
    ...
}
```

Demo:
```
//狗接口
type Doger interface {
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

//接口相关Demo
func main() {
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
}
```
输出：
```
泰迪犬 小白 say 汪汪
泰迪犬 小黄 say 喔喔
斗牛犬 小黑 say 旺旺
```

- 一个接口类型的变量可以包含任何类型的值，可以通过以下方式来检测它的**动态**类型：
```
_, ok := 接口变量.(动态类型)
//ok为true则表示包括该动态类型，反之则不包括
```

Demo: 
```
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

func main() {
  simple1 := new(Simple)
	simple1.Set(2)

	var simplerIn Simpler
	simplerIn = simple1
	if v, ok := simplerIn.(*Simple); ok {
    fmt.Printf("Simple结构体 是 Simpler接口 的动态类型(%T)\n", v)
	} else {
		fmt.Printf("Simple结构体 不是 Simpler接口 的动态类型\n")
	}
}
```

输出：
```
Simple结构体 是 Simpler接口 的动态类型(*code.Simple)
```
