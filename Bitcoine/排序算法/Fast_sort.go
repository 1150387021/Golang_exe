package main

import "fmt"

//快速排序
func quick_sort(arr []float32, left int, right int) []float32 {
	if left < right {
		mid := partition(arr, left, right)
		quick_sort(arr, left, mid-1)
		quick_sort(arr, mid+1, right)

	}
	return arr
}

//快排前奏，归位函数
//让最左边的值归位，使其左边的比他小，右边的比他大
func partition(arr []float32, left int, right int) int {
	temp := arr[left]
	for left < right { //保证至少存在两个数
		for left < right && arr[right] >= temp {
			right--
		}
		arr[left] = arr[right] //将右边的小值写到左边
		//fmt.Println(arr, "left")
		for left < right && arr[left] <= temp {
			left++
		}
		arr[right] = arr[left] //将左边的大值写到右边
		//fmt.Println(arr, "right")
	}
	arr[left] = temp //将原始最左边的值归位
	return left      //返回temp值最终的下标，方便对左右两边递归排序
}

func QuickSort(arr []float32) []float32 {
	length := len(arr)
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0] //以第一个为基准
		low := make([]float32, 0, 0)
		high := make([]float32, 0, 0)
		mid := make([]float32, 0, 0)
		mid = append(mid, splitdata) //加入第一个相等的
		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

////加强版的快速排序
//
//func QuicksortPlus(arr []int) []int {
//
//}

//向下调整函数（建堆）
func sift(arr []float32, low int, high int) []float32 {
	i := low
	j := 2*i + 1
	temp := arr[low]
	for j <= high {
		if j+1 <= high && arr[j+1] > arr[j] {
			j++
		}
		if arr[j] > temp {
			arr[i] = arr[j]
			i = j
			j = 2*i + 1
		} else {
			arr[i] = temp
			break
		}
	}
	arr[i] = temp
	return arr
}

//func Heap_sort(arr []float32) {
//	n:=len(arr)
//	for i:= range ((n-2)/2,-1,-1){
//		sift(arr,i,n-1)
//
//	}
//	for i:=range(n-1,-1,-1){
//		arr[0],arr[i]=arr[i],arr[0]
//		sift(arr,0,i-1)
//	}
//}

//堆排序2

func HeapSort(arr []float32) []float32 {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastMesslen := length - i //每次截取一段
		HeapSortMax(arr, lastMesslen)
		if i < length {
			arr[0], arr[lastMesslen-1] = arr[lastMesslen-1], arr[0]
		}
	}
	return arr
}

func HeapSortMax(arr []float32, length int) []float32 {
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1         //深度
		for i := depth; i >= 0; i-- { //循环所有的三节点
			topmax := i //假定最大的在i的位置
			leftchild := 2*i + 1
			rightchild := 2*i + 2
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] {
				topmax = leftchild
			}
			if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
				topmax = rightchild
			}
			if topmax != i { //确保i的值就是最大的
				arr[i], arr[topmax] = arr[topmax], arr[i]
			}
		}
		return arr
	}

}

//奇偶排序

func OddEven(arr []float32) []float32 {
	isSorted := false //奇数位，偶数位都不需要排序时有序
	for isSorted == false {
		isSorted = true
		for i := 1; i < len(arr)-1; i = i + 2 { //奇数位
			if arr[i] > arr[i+1] { //需要交换
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		//fmt.Println("1", arr)
		for i := 0; i < len(arr)-1; i = i + 2 { //偶数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		//fmt.Println("0", arr)
	}
	return arr
}

//归并排序'
func merge(leftarr []float32, rightarr []float32) []float32 {
	leftindex := 0         //左边索引
	rightindex := 0        //右边索引
	lastarr := []float32{} //最终都数组
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else {
			lastarr = append(lastarr, rightarr[rightindex])
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(leftarr) { //	把没结束都归并过来
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr) { //把没结束都归并过来
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr []float32) []float32 {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])
		return merge(leftarr, rightarr)
	}
}

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
func main() {
	arr := []float32{5, 3, 9, 6, 7, 4, 8, 1, 2}
	fmt.Println("原数组为：", arr)
	fmt.Println("---------")

	//arr1 := quick_sort(arr, 0, len(arr)-1)
	//fmt.Println("快速排序：", arr1)
	//fmt.Println("---------")
	//
	//arr11 := QuickSort(arr)
	//fmt.Println("快速排序1：", arr11)
	//fmt.Println("---------")

	//a := sift(arr, 0, len(arr))
	//fmt.Println(a)
	//fmt.Println("---------")
	//
	//arr2 := HeapSort(arr)
	//fmt.Println("堆排序：", arr2)
	//fmt.Println("---------")
	//
	//arr3 := OddEven(arr)
	//fmt.Println("奇偶排序：", arr3)
	//fmt.Println("---------")

	arr4 := MergeSort(arr)
	fmt.Println("归并排序：", arr4)
	fmt.Println("---------")

	//index := bin_search(arr4, 11)
	index := bin_search_zhongzhi(arr4, 4)
	if index == -1 {
		fmt.Println("没有找到")
	} else {
		fmt.Println("找到数字：", arr4[index], "位于：", index)
	}
}
