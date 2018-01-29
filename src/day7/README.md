# 第七天内容
***
*所有详细代码都能在code中有展示，可查看*

# Go读写（I/O）
***

## 3 文件拷贝
- 实现文件的互相拷贝只需使用`io`标准库的`Copy`方法即可

demo:
```
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
```