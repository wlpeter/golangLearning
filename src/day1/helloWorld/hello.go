package helloWorld

import (
	"fmt"
	"runtime"
)

func SayHello() {
	fmt.Printf("hello world\n")
	fmt.Printf("%s\n", runtime.Version())
}

func SayGoodBey() {
	fmt.Printf("good bey\n")
	fmt.Printf("%s\n", runtime.Version())
}