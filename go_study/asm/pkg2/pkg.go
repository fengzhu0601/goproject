package pkg2

import "image"

// 全局变量
var count = 0x1234

// 数组
var num [2]int

// bool变量
var (
	boolValue  bool
	trueValue  bool
	falseValue bool
)

// int型变量
var int32Value int32
var uint32Value uint32

// float型变量
var float32Value float32
var float64Value float64

// string型变量
var helloworld string

// slice
var sliceString []byte

// map/channel型变量
var m map[string]int
var ch chan int

// 标准库图像包
var pt image.Point

// 函数
//
//go:nosplit
func Swap(a, b int) (ret0, ret1 int) {
	return
}
