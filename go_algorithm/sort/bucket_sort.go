package main

import (
	"fmt"
)

// 桶排序函数
func bucketSort(arr []int) {
	n := len(arr)

	// 找到数组中的最大值和最小值
	minVal, maxVal := arr[0], arr[0]
	for _, num := range arr {
		if num < minVal {
			minVal = num
		} else if num > maxVal {
			maxVal = num
		}
	}

	// 计算桶的数量，并创建桶
	bucketSize := (maxVal-minVal)/n + 1
	buckets := make([][]int, bucketSize)
	for i := 0; i < bucketSize; i++ {
		buckets[i] = make([]int, 0)
	}

	// 将元素放入对应的桶中
	for _, num := range arr {
		index := (num - minVal) / n
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶进行排序，并将排序后的元素放回原数组
	index := 0
	for i := 0; i < bucketSize; i++ {
		insertionSort(buckets[i])

		for _, num := range buckets[i] {
			arr[index] = num
			index++
		}
	}
}

// 插入排序函数
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func main() {
	arr := []int{9, 5, 1, 8, 3, 7, 4, 6, 2}
	fmt.Println("原始数组:", arr)

	bucketSort(arr)

	fmt.Println("排序后数组:", arr)
}
