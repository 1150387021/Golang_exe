package main

import "fmt"

//A找到第一个等于3的

//D找到最后一个小于7的数据

func bin_searchA(arr []float32, data float32) int {
	left := 0
	right := len(arr) - 1
	index := -1
	for left <= right {
		mid := (left + right) / 2 //mid := left + (-left + right) / 2
		if data > arr[mid] {
			left = mid + 1
		} else if data < arr[mid] {
			right = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != data {
				index = mid
				fmt.Println("index:", index)
				break
			} else {
				right = mid - 1 //递归继续查找
			}
		}
	}
	return index
}

//B找到最后一个等于3的
func bin_searchB(arr []float32, data float32) int {
	left := 0
	right := len(arr) - 1
	index := -1
	for left <= right {
		mid := (left + right) / 2 //mid := left + (-left + right) / 2
		if data > arr[mid] {
			left = mid + 1
		} else if data < arr[mid] {
			right = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != data {
				index = mid
				fmt.Println("index:", index)
				break
			} else {
				left = mid + 1 //递归继续查找
			}
		}
	}
	return index
}

//C找到第一个大于等于data的index
func bin_searchC(arr []float32, data float32) int {
	left := 0
	right := len(arr) - 1
	index := -1
	for left <= right {
		mid := (left + right) / 2 //mid := left + (-left + right) / 2
		if data > arr[mid] {
			left = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < data {
				index = mid
				fmt.Println("index:", index)
				break
			} else {
				right = mid - 1
			}
		}
	}
	return index
}
func main() {
	arr := []float32{1, 2, 3, 4, 4, 4, 5, 6, 6, 6, 7, 8, 9, 10}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(i, arr[i])
	//}
	bin_searchC(arr, 5)
}
