package code

import "fmt"
import "os"
import "io"
import "io/ioutil"
//import "strings"
import "flag"

func Output() {
	fmt.Println("********************* 读写相关 **********************")
	fmt.Println("********************* 3 文件拷贝 **********************")
	fileCopyFunc()
	fmt.Println("********************* 4 命令行读取 **********************")
	commandLineIOfunc()
}

func fileCopyFunc() {
	copyFile("file/test1.txt", "file/test2.txt")
	buf, _ := ioutil.ReadFile("file/test2.txt")
	fmt.Println(string(buf))
	return
}

func copyFile(srcPath, targetPath string) (written int64, err error) {
	if srcPath == "" || targetPath =="" {
		fmt.Fprintf(os.Stderr, "Param error: %s\n", "参数有误")
	}
	src, srcErr := os.Open(srcPath)
	if srcErr != nil {
		fmt.Fprintf(os.Stderr, "Open src file error: %s\n", srcErr)
		return
	}
	defer src.Close()

	target, targetErr := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE, 0666)
	if targetErr != nil {
		fmt.Fprintf(os.Stderr, "Open target file error: %s\n", targetErr)
		return
	}
	defer target.Close()

	return io.Copy(target, src)
}

var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func commandLineIOfunc() {
	// people := "Peter "
	// if len(os.Args) > 1 {
	// 	people += strings.Join(os.Args[1:], " ")
	// }
	// fmt.Println("Hello ", people)
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
			if i > 0 {
					s += " "
					if *NewLine { // -n is parsed, flag becomes true
							s += Newline
					}
			}
			s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}