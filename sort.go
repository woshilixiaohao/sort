package algothem

import (
	"fmt"
)

func BubbleSort(arr [12]int) {
	// n个数 进行n-1次比较，相邻的2个数比较，最大的数在最后
	// n-1个数 进行n-2
	// n-2个数 进行n-3
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func SelectionSort(arr [12]int) {
	// n个数 比较n-1次 选择最小的数，放在第一个
	// n-1个数
	for i := 0; i < len(arr)-1; i++ {
		// 最小值位置
		min := i
		for j := min + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	fmt.Println(arr)
}
