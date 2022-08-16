package main

import "fmt"

func GnomeSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++ //符合顺序，继续前进
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
			//fmt.Println(arr)
		}
	}
	return arr
}
func main() {
	arr := []int{5, 3, 9, 6, 7, 4, 8, 1, 2}
	fmt.Println(GnomeSort(arr))
}
