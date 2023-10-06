package main

import (
	"fmt"
)

// 获取待排序数组中最大的数
func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// 对数组按照指定位数进行计数排序
func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// 统计每个数字出现的次数
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	// 计算每个数字在输出数组中的位置
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// 将数字按照计算好的位置放入输出数组中
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	// 将输出数组复制回原数组
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// 基数排序函数
func radixSort(arr []int) {
	max := getMax(arr)

	// 按照位数依次进行计数排序
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("原始数组:", arr)

	radixSort(arr)

	fmt.Println("排序后数组:", arr)
}
