package main

//基于C标准库函数输出字符串

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("hello world!\n"))
}

/*
could not determine kind of name for C.Puts
必须要加这样的一句话：
//#include <stdio.h>
import "C"
*/
