package code

import "fmt"
import "os"
import "bufio"
import "io"
import "runtime"
import "io/ioutil"
import "strings"

func Output() {
	fmt.Println("********************* 读写相关 **********************")
	fmt.Println("********************* 1.垃圾回收 **********************")
	gcFunc()
	fmt.Println("********************* 2.读取用户的输入 **********************")
	// stdinFunc()
	fmt.Println("********************* 3.文件读写 **********************")
	fileRFunc()
	fileWFunc()
}

func gcFunc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	defer fmt.Printf("已分配堆内存 %d Kb\n", m.Alloc / 1024)
	defer fmt.Printf("未分配内存 %d Kb\n", m.Frees  / 1024)
	defer fmt.Printf("已分配堆内存累计值（包括已释放的） %d Kb\n", m.TotalAlloc  / 1024)
	defer fmt.Printf("系统保留的总内存 %d Kb\n", m.Sys  / 1024)
}

func stdinFunc() {
	var i, j string
	fmt.Scanln(&i, &j)
	fmt.Println(i + " and " + j)

	var num1, num2 int
	fmt.Scanf("%d + %d", &num1, &num2)
	fmt.Printf("%v\n", num1 + num2)

	var d int
	var s string
	input := "56 and Go"
	format := "%d and %s"
	fmt.Sscanf(input, format,  &d, &s)
	fmt.Println("From the string we read: ", d, s)
	
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	//btyeData, err := inputReader.ReadBytes('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", str)
	} else {
		fmt.Println(err)
	}
	//stdinTestFunc()
}

func stdinTestFunc() {
	var strNum, enNum, lineNum int
	lineNum = 1
	inputReader := bufio.NewReader(os.Stdin)
	bytesList, _ := inputReader.ReadBytes('S')
	for i, _ := range bytesList {
		if (bytesList[i] == 10) {
			lineNum++
			continue
		} else if (bytesList[i] == 13) {
			continue
		} else {
			strNum++
		}
		if ((bytesList[i] >= 65 && bytesList[i] <= 90) || (bytesList[i] >= 97 && bytesList[i] <= 122)) {
			enNum++
		}
	}
	fmt.Printf("The result was: %d , %d, %d\n", strNum, enNum, lineNum)
}

func fileRFunc() {

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
			break
		}
		fmt.Printf("The %d line was: %s\n", i, str)
	}
	

	//按列读取文件
	fmt.Println("vvvv按列读取文件vvvv")
	file, _ := os.Open("file/test2.txt")
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, colErr := fmt.Fscanln(file, &v1, &v2, &v3)
		if (colErr != nil) {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)

	fileTestFunc();
}


type Menu struct {
	title, price, quantity string
}

func fileTestFunc() {
	file, err := os.Open("file/test3.txt")
	if err != nil {
		panic(err)
	}
	inputReader := bufio.NewReader(file)
	var arr []Menu
	for {
		res, _, error := inputReader.ReadLine()
		if error != nil {
			break
		}
		var menu Menu
		list := strings.Split(string(res), ";")
		menu.title = list[0]
		menu.price = list[1]
		menu.quantity = list[2]
		arr = append(arr, menu)
		fmt.Println(menu)
	}
	fmt.Println(arr)
}


func fileWFunc() {
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
		panic(writerError)
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

	fmt.Println("**************文件write练习*****************")
	p := new(Page)
	p.Title = "outTest4"
	p.Body = []byte("好好学习\n天天向上")
	p.save();

	p1 := load(p.Title)
	fmt.Printf("Title : %s\n", p1.Title)
	fmt.Printf("Body : %s\n", p1.Body)
}

type Page struct {
	Title string
	Body []byte
}

func (p Page)save() {
	if p.Title == "" || p.Body == nil || len(p.Body) == 0 {
		fmt.Fprintf(os.Stderr, "Param Error: %s\n", "参数不能为空")
	}
	path := "file/" + p.Title + ".txt"
	file, fileError := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if fileError != nil {
		fmt.Fprintf(os.Stderr, "Open File(%s) Error: %s\n", path, fileError)
		panic(fileError)
		return
	}
	defer file.Close()
	
	//创建文件写入器
	fileWriter := bufio.NewWriter(file)
	_, writerError := fileWriter.Write(p.Body)
	if writerError != nil {
		fmt.Fprintf(os.Stderr, "Write File(%s) Error: %s\n", path, writerError)
		panic(writerError)
		return
	}
	//将缓存上的所有数据写入到底层的io.Writer中
	fileWriter.Flush()
}

func load(title string) (p *Page) {
	if title == "" {
		fmt.Fprintf(os.Stderr, "Param Error: %s\n", "文件名不能为空")
	}
	path := "file/" + title + ".txt"
	buf, error := ioutil.ReadFile(path)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %s\n", error)
		return
	}
	pa := Page{ Title: title, Body: buf }
	p = &pa
	return
}
