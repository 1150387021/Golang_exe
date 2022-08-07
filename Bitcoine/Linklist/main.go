package main

import (
	"fmt"
)

func main1() {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)
	node1.data = 1
	node1.pNext = node2
	node2.data = 2
	node2.pNext = node3
	node3.data = 3
	node3.pNext = node4
	node4.data = 4
	fmt.Println(node1.data)
	fmt.Println(node2.data)
	fmt.Println(node3.data)
	fmt.Println(node4.data)
	fmt.Println("--------------")
	fmt.Println(node1.pNext.data)
	fmt.Println(node1.pNext.pNext.pNext.data)

}

func main2() {
	mystack := NewStack()
	for i := 0; i < 50; i++ {
		mystack.Push(i)
	}
	for data := mystack.Pop(); data != nil; data = mystack.Pop() {
		fmt.Println(data)
	}
}
func main() {
	myq := NewLinkQueue()
	for i := 0; i < 1000; i++ {
		myq.Enqueue(i)
	}
	for data := myq.Dequeue(); data != nil; data = myq.Dequeue() {
		fmt.Println(data)
	}
}

// 栈：先进后出，操作口只有一个，头部删除头部插入，尾部插入，尾部删除
// 队列：头部插入，尾部删除  ｜｜ 尾部插入，头部删除

//文件遍历
//轻量级 数组栈 深度遍历； 数组队列 广度遍历
//重量级 链表栈 深度遍历； 链表队列 广度遍历

//数组实战，队列，实现文件遍历，搜索*.txt
//链表实战，队列，实现文件遍历,搜索*.txt带"关键词"
