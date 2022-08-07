package main

import (
	"fmt"
)

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { //插入排序的变种
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] { //从i的左边向左扫描找到arr[i]的位置
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup

	}

}

//希尔排序

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ { //处理每个元素的步长
				ShellSortStep(arr, i, gap)
			}
			gap /= 2 //gap--
		}
	}
	return arr

}

//选出数组中的最大值

func Select_sortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < length; i++ { //前面的0-i已完成排序不用懂
			if arr[i] > max {
				arr[i], max = max, arr[i] //交换数字，较大的放前面去
			}
		}
		return max
	}

}

// 基数排序

func RadixSort(arr []int) []int {
	max := Select_sortMax(arr) //寻找数组极大值
	for bit := 1; max/bit > 0; bit *= 10 {
		//按照数量级分段
		arr = BitSort(arr, bit) //每次处理一个位数的排序
		//fmt.Println(arr)
	}
	return arr
}

func BitSort(arr []int, bit int) []int {
	length := len(arr)           //数组长度
	bitcounts := make([]int, 10) //统计长度
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10 //分层处理，bit=1000的，三位数不参与排序
		bitcounts[num]++           //统计余数相等的个数
	}
	//fmt.Println(bitcounts)
	for i := 1; i < 10; i++ {
		bitcounts[i] += bitcounts[i-1] //叠加，计算位置
	}
	tmp := make([]int, 10)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitcounts[num]-1] = arr[i] //计算排序的位置
		bitcounts[num]--
	}
	for i := 0; i < length; i++ {
		arr[i] = tmp[i] //保存数组
	}
	return arr
}

//鸡尾酒排序
func cocktail(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ { //每次循环，正向冒泡一次，反向冒泡一次
		left := 0
		right := len(arr) - 1 //左边，右边
		for left <= right {   //结束条件
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			fmt.Println(left, arr)
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
			fmt.Println(right, arr)
		}
		//fmt.Println(i, arr)
	}
	return arr
}
func main() {
	arr := []int{5, 3, 9, 6, 7, 4, 8, 1, 2}
	fmt.Println("原数组为：", arr)
	fmt.Println("---------")

	//arr1 := ShellSort(arr)
	//fmt.Println("希尔排序：", arr1)
	//fmt.Println("---------")
	//
	//arr2 := ShellSort(arr)
	//fmt.Println("希尔排序：", arr2)
	//fmt.Println("---------")

	//arr3 := RadixSort(arr)
	//fmt.Println("基数排序：", arr3)
	//fmt.Println("---------")

	arr4 := cocktail(arr)
	fmt.Println("鸡尾酒排序：", arr4)
	fmt.Println("---------")

}
