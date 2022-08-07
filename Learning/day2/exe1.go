package main

import "fmt"

func ptr() {
	x := "学习基础"
	ptr := &x                      //直接去拿x数据的内存地址
	fmt.Println("ptr的内存地址为：", ptr) //指针显示内存地址
	fmt.Println(*ptr)              //显示这个指针对应的数据是什么

	ptr1 := new(string)
	*ptr1 = "你好啊赛利亚！"
	fmt.Println("ptr1的内存地址为：", ptr1)
	fmt.Println(*ptr1)

	x2 := "飞流直下三千尺"
	var ptr2 *string
	ptr2 = &x2
	fmt.Println("ptr2的内存地址为：", ptr2)

}
func store() { //计算商品的总价
	var (
		name       string
		num, price float32
	)
	fmt.Println("请输入商品名称：")
	fmt.Scanln(&name)
	fmt.Println("请输入商品价格：")
	fmt.Scanln(&price)
	fmt.Println("请输入商品数量：")
	fmt.Scanln(&num)
	fmt.Println("商品名称为：", name, "商品价格为：", price, "商品数量为：", num)
	total := price * num
	fmt.Println("总价为：", total)

}

func main() {
	//ptr()
	store()

}
