package main

import "fmt"

func shellSort(arr []int) []int {
	length := len(arr)
	gap := 4
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	fmt.Println("gap = ", gap)
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
		fmt.Println("gap1 = ", gap, arr)
	}
	return arr
}

func main() {
	arr := []int{9, 5, 1, 8, 3, 7, 4, 6, 2}
	fmt.Println("原始数组:", arr)

	shellSort(arr)

	fmt.Println("排序后数组:", arr)
}
