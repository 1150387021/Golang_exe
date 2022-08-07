package main

import (
	"fmt"
)

func test() {
	var arr1 = [...]int{2, 3, 4, 5, 6}
	arr2 := [5]float32{3.4, 5, 6, 77.8}
	arr3 := [][]string{
		{"1", "nihoao"},
		{"4", "jjjj"},
		{"nnn", "444"},
	}
	fmt.Println(arr1, arr2)
	fmt.Printf("type of arr3: %T\n", arr3)
	fmt.Println(arr3)
}

func showarr() {
	arr3 := [][]string{
		{"1", "nihoao"},
		{"4", "jjjj"},
		{"nnn", "444"},
	}
	for ind, arr := range arr3 {
		fmt.Println(ind, arr)
	}
}
func createslice() {
	var numberlist = []int{}
	numlist := make([]int, 3, 5)
	fmt.Println(numlist)
	fmt.Println(numberlist)
}

func createmap() {
	step1 := make(map[string]string)
	fmt.Println(step1)
	var step2 map[string]string = map[string]string{
		"敲门": "step1",
		"你好": "step2",
	}
	//短申明的方式
	step3 := map[string]string{
		"敲门": "step1",
		"你好": "step2",
	}
	//给map中添加键值对
	step3["进门"] = "step3"
	fmt.Println(step2)
	fmt.Println(step3)
}

func main() {
	//test()
	//showarr()
	//createslice()
	createmap()

}
