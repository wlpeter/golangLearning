# 第六天内容
***
*所有详细代码都能在code中有展示，可查看*


# Go垃圾回收（GC）
***
Go语言不像C/C++须要自行回收内存，Go语言有独立的进程去自行操作垃圾回收，即**垃圾回收器（GC）**，它会自行搜索不再使用的变量然后释放它们的内存
Go语言使用的垃圾回收算法是**标记清扫算法**，而且由于Go能识别每个变量的类型，所以可以实现**精确的垃圾回收**，运行时分为两部分：
- 1、标记阶段，从**root区域**出发，扫描所有root区域的对象直接或间接引用到的对象，将这些对上全部加上标记
- 2、清扫阶段，扫描整个堆区，对所有无标记的对象进行回收


# Go读写（I/O）
***

## 1 读取用户的输入
- 从键盘和标准输入可以使用`os`库的`Stdin`读取输入或者使用`fmt`的`Scan`和`Sscan`开头的函数
- `fmt`的`Sscan`是通过首个string参数中读取数据，`Scan`则通过键盘输入读取
- `os.Stdin`为创建的读取器可获取键盘输入内容

demo:
```
import "fmt"

func main() {
	var i, j string
	fmt.Scanln(&i, &j)
	fmt.Println(i + " and " + j)              //输出 i(值)与' and '与j(值)的字符串拼接结果

	var num1, num2 int
	fmt.Scanf("%d + %d", &num1, &num2)        //输入 num1(值) + num2(值)
	fmt.Printf("%v\n", num1 + num2)           //输出 num1(值) + num2(值)的相加结果

	var d int
	var s string
	input := "56 and Go"
	format := "%d and %s"
	fmt.Sscanf(input, format,  &d, &s)                      
  fmt.Println("From the string we read: ", d, s)          //输出：From the string we read: 56 Go

  var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)         //NewReader(rd io.Reader) *Reader 会将一个读取器与标准输入绑定，返回一个带缓冲的 io.Reader 对象
	str, err := inputReader.ReadString('\n')        //ReadString(delim byte),该方法从输入中读取内容，直到碰到 delim 指定的字符,返回string类型
	//btyeData, err := inputReader.ReadBytes('\n')  //ReadBytes(delim byte),该方法从输入中读取内容，直到碰到 delim 指定的字符,返回btye类型
	if err == nil {
		fmt.Printf("The input was: %s\n", str)    //输出 The input was: （输入的参数）
	} else {
		fmt.Println(err)
	}
}
```

## 2 文件读写
在Go中，文件使用指向`os.File`类型的**指针**来表示的，也叫做**文件句柄**。
一般会通过`os`标准库的`Open`或`OpenFile`方法来打开文件（即获取文件句柄）
- `Open`方法只能以**只读模式**打开文件
- `OpenFile`可以通过参数指定任意一种模式打开文件，模式如下：
- - os.O_RDONLY：只读
- - os.O_WRONLY：只写
- - os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
- - os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。
- 使用`OpenFile`时，第三个参数指代文件的**权限**。**读文件**时，第三个参数可以用`0`。而在**写文件**时，不管是 Unix 还是 Windows，都需要使用`0666`。


### 2.1 文件读取
- 可直接使用`io/ioutil`标准库的方法`ioutil.ReadFile(文件名)`直接读取文件所有内容
- 也可通过`bufio`标准库来创建**读取器**，来进行带缓冲的读取文件数据

demo：
```
//整个文件直接读取
fmt.Println("vvvv整个文件直接读取vvvv")
buf, error := ioutil.ReadFile("file/test1.txt")
fmt.Printf("%s\n", string(buf))
if error != nil {
	fmt.Fprintf(os.Stderr, "File Error: %s\n", error)
	return
}


//整个文件通过缓冲区读取
fmt.Println("vvvv整个文件通过缓冲区读取vvvv")
//获取文件句柄
inputFile, err := os.Open("file/test1.txt")
if err != nil {
	fmt.Println(err)
	return
}
//关闭该文件
defer inputFile.Close()
inputReader := bufio.NewReader(inputFile)
buff := make([]byte, 1024)
btyeNum, readError := inputReader.Read(buff)
if readError != nil {
	fmt.Fprintf(os.Stderr, "File Error: %s\n", readError)
	return
}
fmt.Printf("读取字节数：%d\n", btyeNum)
fmt.Printf("读取文件内容：%s\n", string(buf))


//通过缓冲区逐行读取
fmt.Println("vvvv通过缓冲区逐行读取vvvv")
//获取文件句柄
inputFile1, err := os.Open("file/test1.txt")
if err != nil {
	fmt.Println(err)
	return
}
//关闭该文件
defer inputFile1.Close()
inputReader2 := bufio.NewReader(inputFile1)
for i := 1; true; i++ {
	str, _, readLineError := inputReader2.ReadLine()
	if readLineError == io.EOF {
		return
	}
	fmt.Printf("The %d line was: %s\n", i, str)
}
```

### 2.2 文件写入
- 可直接使用`io/ioutil`标准库的方法`ioutil.WriteFile(文件名, []byte类型的数据)`直接写入文件所有内容
- 可直接使用通过`os.OpenFile`获取的**文件句柄**执行`Write`系列方法写入数据
- 由于通过`os.OpenFile`获取的类型也是`io.Writer`,所以也可使用`fmt.Fprintf`来写入文件
- 也可通过`bufio`标准库来创建**写入器**，来进行带缓冲的写入文件数据，最后将缓存上的所有数据写入到底层的`io.Writer`中
demo：
```
//直接写文件
ioutil.WriteFile("file/outTest1.txt", []byte("qqqqqqq\nqqweq"), 0666)

//获取文件句柄，以只写模式打开文件，如果不存在就创建文件
outputFile, outputError := os.OpenFile("file/outTest2.txt", os.O_WRONLY|os.O_CREATE, 0666)
if outputError != nil {
	fmt.Fprintf(os.Stderr, "Open File(outTest2.txt) Error: %s\n", outputError)
	panic(outputError)
	return
}
defer outputFile.Close()
//创建文件写入器
fileWriter := bufio.NewWriter(outputFile)
//把数据写入缓存
_, writerError := fileWriter.WriteString("12312312\nweeewww")
if writerError != nil {
	fmt.Fprintf(os.Stderr, "Write File(outTest2.txt) Error: %s\n", writerError)
	panic(outputError)
	return
}
//将缓存上的所有数据写入到底层的io.Writer中
fileWriter.Flush()


//获取文件句柄，以只写模式打开文件，如果不存在就创建文件
outputFile3, outputError3 := os.OpenFile("file/outTest3.txt", os.O_WRONLY|os.O_CREATE, 0666)
if outputError3 != nil {
	fmt.Fprintf(os.Stderr, "Open File(outTest3.txt) Error: %s\n", outputError3)
	panic(outputError3)
	return
}
defer outputFile3.Close()
//直接写文件
fmt.Fprintf(outputFile3, "Some test data222222222.\n")
//outputFile3.Write([]byte("wwwwwwwwwwwww\ntttttttt"))
//outputFile3.WriteString("wewww22312312\n1231123123")
```