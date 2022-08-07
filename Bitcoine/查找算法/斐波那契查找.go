package main

import "fmt"

//二分查找

func bin_search(arr []float32, data float32) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2 //mid := left + (-left + right) / 2
		if data > arr[mid] {
			left = mid + 1
		} else if data < arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//二分法的中值改进
func bin_search_zhongzhi(arr []float32, data float32) int {
	left := 0
	right := len(arr) - 1
	i := 0
	for left <= right {
		//mid := (left + right) / 2  //mid := left + diff / 2
		//根据其所占整体的比例来划分找中点，一次筛掉大半
		leftv := float64(data - arr[left])      //数据距离左边的
		allv := float64(arr[right] - arr[left]) //总长
		diff := float64(right - left)
		mid := int(float64(left) + leftv/allv*diff)
		i++
		if mid < 0 || mid >= len(arr) {
			return -1
		}

		if data > arr[mid] {
			left = mid + 1
		} else if data < arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//构造斐波那契数列

func makeFabArray(arr []int) []int {
	length := len(arr)
	flblen := 2
	first, second, third := 1, 2, 3
	for third <= length { //找出最接近的斐波那契
		third, first, second = first+second, second, third //叠加计算斐波那契
		flblen++
	}
	fb := make([]int, flblen) //开辟数组
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < flblen; i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}

//斐波那契查找

func fab_search(arr []int, val int) int {
	length := len(arr)
	fabArr := makeFabArray(arr) //定制匹配的斐波那契数组
	fmt.Println(fabArr)
	filllength := fabArr[len(fabArr)-1] //填充长度
	fillArr := make([]int, filllength)  //填充的数组
	for i, v := range arr {
		fillArr[i] = v
	}
	lastdata := arr[length-1] //填充最后一个大数
	for i := length; i < filllength; i++ {
		fillArr[i] = lastdata //填充数据
	}
	fmt.Println(fillArr)

	left, mid, right := 0, 0, length //
	kindex := len(fabArr) - 1        //游标
	for left <= right {
		mid = left + fabArr[kindex-1] - 1 //斐波那契切割
		if val < fillArr[mid] {
			right = mid - 1
			kindex--
		} else if val > fillArr[mid] {
			left = mid + 1
			kindex -= 2
		} else {
			if mid > right {
				return right //越界
			} else {
				return mid
			}
		}
	}
	return -1
}

func main() {
	n := 10
	arr := make([]int, n, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	//fmt.Println(makeFabArray([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	index := fab_search(arr, 12)
	fmt.Println("index", index)
	if index == -1 {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到")
	}

}
