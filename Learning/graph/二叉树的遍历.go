package main

import "fmt"

//使用单向链表

type TreeNode struct {
	val   int
	left  *TreeNode //left是TreeNode类型的指针
	right *TreeNode //右指针
}

func preOrderTraverse(cur *TreeNode) {
	if cur == nil {
		return
	}
	fmt.Println(cur.val)        //访问跟节点
	preOrderTraverse(cur.left)  //访问左子树
	preOrderTraverse(cur.right) //访问右子树
}

//栈的遍历

type Stack struct {
	top int // 栈顶指针
	size int // 栈内元素数量
	capacity int // stack 的容量
	data []*TreeNode // 具体存储数据放在⼀个不定⻓的切⽚ slice 中 }
}


func preordertraverse(cur *TreeNode){
	isVisible := make(map[*TreeNode]int)
	stack.Push(cur)
	for !stack.IsEmpty() { // 栈不为空
		var tmp *TreeNode = stack.Pop() // 从栈⾥出栈⼀个元素
		if tmp == nil {
			continue // 当前节点为空，则继续下⼀次循环
		}
		if _, ok := isVisible[tmp] ; ok { // 如果 isVisible ⾥⾯有 tmp
			// 证明此时是第⼆次”⻅到” tmp, 打印tmp.Val 即可
			fmt.Printf(“%d ”, tmp.Val)
		} else {
			isVisible[tmp] = true // 标记tmp已经”⻅过”⼀次了
			stack.Push(tmp.Right)
			stack.Push(tmp.Left)
			stack.Push(tmp)
		}
	}
}


func main() {
	stack := NewStack(100) // 定义⼀个⾜够⼤的栈
	isVisible := make(map[*TreeNode]int) // 定义⼀个isVisible map
	stack.Push(root) // ⾸先是根节点⼊栈
	for !stack.IsEmpty() { // 栈不为空
		var tmp *TreeNode = stack.Pop() // 从栈⾥出栈⼀个元素
		if tmp == nil {
			continue // 当前节点为空，则继续下⼀次循环
		}

}
