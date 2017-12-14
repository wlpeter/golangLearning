package code

import (
	"fmt"
	"runtime"
)

func SayHello1() {
	fmt.Printf("hello world1\n")
	fmt.Printf("%s\n", runtime.Version())
}

func SayGoodBey1() {
	fmt.Printf("good bey2\n")
	fmt.Printf("%s\n", runtime.Version())
}