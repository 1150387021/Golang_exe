package main

import "fmt"

//选择排序
//全局中选择一个最大（小）值排到最前面，然后在剩余的数组中找最大排第二，以此类推完成排序

func Select_sort(arr []float32) []float32 {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		max := i
		for j := i + 1; j < length; j++ { //前面的0-i已完成排序不用懂
			if arr[j] < arr[max] {
				arr[j], arr[max] = arr[max], arr[j] //交换数字，较大的放前面去
			}
		}
		//fmt.Println(arr)
	}
	return arr
}

//插入排序
//扑克牌插入，从右到左依次选取根据大小依次插入到左边

func Insert_sort(arr []float32) []float32 {
	for i := 1; i < len(arr); i++ { //i为摸到的牌下标
		backup := arr[i]
		j := i - 1
		for j >= 0 && backup < arr[j] { //从i的左边向左扫描找到arr[i]的位置
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = backup
	}
	return arr
}

//冒泡排序
//一趟实现从左至右两两相邻到大小交换，实现部分有序和找到一个最大值。
func bubble_sort(arr []float32) []float32 {
	exchange := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				exchange = true
			}
		}
		//fmt.Println(arr)
	}
	if !exchange {
		return arr
	}
	return arr
}

func main() {
	arr := []float32{5, 3, 9, 6, 7, 4, 8, 1, 2}
	fmt.Println("原数组为：", arr)
	fmt.Println("---------")

	//arr1 := Select_sort(arr)
	//fmt.Println("选择排序：", arr1)
	//fmt.Println("---------")

	//arr2 := Insert_sort(arr)
	//fmt.Println("插入排序：", arr2)
	//fmt.Println("---------")

	arr4 := bubble_sort(arr)
	fmt.Println("冒泡排序：", arr4)
	fmt.Println("---------")

}
