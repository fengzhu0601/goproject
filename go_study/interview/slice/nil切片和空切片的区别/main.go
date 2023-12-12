package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)
	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n",
		*(*reflect.SliceHeader)(unsafe.Pointer(&s1)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&s4)),
	)
	fmt.Printf("%v\n",
		(*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n",
		(*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}

/**
1. nil切片表示该切片尚未初始化，仅有一个类型声明，并不存在实例。引用数组指针地址为0。
2. 空切片表示该切片以及被初始化，有一个切实存在的对象，只是该切片的元素个数为0。引用的数
组指针地址也是一个存在的内存地址
*/
