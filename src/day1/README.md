# 第一天内容
***

# 安装
***
1、下载（版本自选）
```
wget http://www.golangtc.com/static/go/go1.6.2.linux-amd64.tar.gz
安装包下载地址：http://golangtc.com/download
```

2、解压安装包指/usr/local目录下
```
sudo tar -zxvf go1.6.2.linux-xxx.tar.gz -C /usr/local
```

3、创建go工作空间
```
mkdir /root/workspace/gowork
cd /root/workspace/gowork
mkdir bin
mkdir pkg
mkdir src
```

4、配置环境变量，指定go工作空间，用vim打开/etc/profile，添加如下文字
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/root/workspace/gowork
export PATH=$PATH:$GOPATH/bin
```
5、输入如下命令，重新加载环境变量
```
source /etc/profile
```
6、检查go
```
go version
```

## go环境变量详解(go env 可查看go环境变量)：
- $GOROOT 表示 Go 在你的电脑上的安装位置，它的值一般都是 $HOME/go，当然，你也可以安装在别的地方。
- $GOARCH 表示目标机器的处理器架构，它的值可以是 386、amd64 或 arm。
- $GOOS 表示目标机器的操作系统，它的值可以是 darwin、freebsd、linux 或 windows。
- $GOBIN 表示编译器和链接器的安装位置，默认是 $GOROOT/bin，如果你使用的是 Go 1.0.3 及以后的版本，一般情况下你可以将它的值设置为空，Go 将会使用前面提到的默认值。
- $GOPATH 默认采用和 $GOROOT 一样的值，但从 Go 1.1 版本开始，你必须修改为其它路径。它可以包含多个包含 Go 语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：src、pkg 和 bin，这三个目录分别用于存放源码文件、包文件和可执行文件。
- $GOARM 专门针对基于 arm 架构的处理器，它的值可以是 5 或 6，默认为 6。
- $GOMAXPROCS 用于设置应用程序可使用的处理器个数与核数

***


# go 相关命令详解
***
go命令详解链接：http://wiki.jikexueyuan.com/project/go-command-tutorial/0.0.html
## go build && go build xxx.go/xxx :
用于测试编译。在包的编译过程中，若有必要，会同时编译与之相关联的包。
  1. 如果是普通包，当你执行go build命令后，不会产生任何文件。
  2. 如果是main包，当只执行go build命令后，会在当前目录下生成一个可执行文件。
  3. 如果某个文件夹下有多个文件，而你只想编译其中某一个文件，可以在 go build 之后加上文件名，例如 go build a.go；go build 命令默认会编译当前目录下的所有go文件。
  4. 你也可以指定编译输出的文件名。比如，我们可以指定go build -o myapp.exe，默认情况是你的package名(非main包)，或者是第一个源文件的文件名(main包)。
  5. go build 会忽略目录下以”_”或者”.”开头的go文件。
  6. 如果你的源代码针对不同的操作系统需要不同的处理，那么你可以根据不同的操作系统后缀来命名文件。例如有一个读取数组的程序，它对于不同的操作系统可能有如下几个源文件：
```
array_linux.go 
array_darwin.go 
array_windows.go 
array_freebsd.go
```
go build的时候会选择性地编译以系统名结尾的文件（Linux、Darwin、Windows、Freebsd）。例如Linux系统下面编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。

## go install && go install xxx.go/xxx :
命令在内部实际上分成了两步操作：第一步是生成结果文件(可执行文件或者.a包)，第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。
.exe文件/可执行文件： 是go install带main函数的go文件产生的，有函数入口，可以直接运行。
.a应用包： 一般是 go install 不包含main函数的go文件产生的，没有函数入口，只能被调用。

## go test:
会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件，包含名为 TestXXX 且签名为 func (t *testing.T) 函数的文件来编写测试。 测试框架会运行每一个这样的函数；若该函数调用了像 t.Error 或 t.Fail 这样表示失败的函数，此测试即表示失败。

## go doc xxx && godoc -... xxx:
go doc 命令其实就是一个很强大的文档工具。
如何查看相应package的文档呢？ 例如builtin包，那么执行go doc builtin；如果是http包，那么执行go doc net/http；查看某一个包里面的函数，那么执行godoc fmt Printf；也可以查看相应的代码，执行godoc -src fmt Printf；
通过命令在命令行执行 godoc -http=:端口号 比如godoc -http=:8080。然后在浏览器中打开127.0.0.1:8080，你将会看到一个golang.org的本地copy版本，通过它你可以查询pkg文档等其它内容。如果你设置了GOPATH，在pkg分类下，不但会列出标准包的文档，还会列出你本地GOPATH中所有项目的相关文档，这对于经常被限制访问的用户来说是一个不错的选择。

## go fmt xxx.go/xxx && gofmt -w xxx.go/xxx:
格式化go文件或项目

## go run :
编译并运行Go程序

## go env:
查看当前go的环境变量

## go version:
查看当前go的版本信息

## go list:
列出当前全部安装的package

## go fix:
用来修复以前老版本的代码到新版本，例如go1之前老版本的代码转化到go1
***