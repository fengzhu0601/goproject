package main

//void SayHello(_GoString_ s);//在Go1.10中CGO新增加了一个_GoString_预定义的C语言类型，用来表示Go语言字符串。
import "C"

import (
	"fmt"
)

func main() {
	C.SayHello("Hello, World\n")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
