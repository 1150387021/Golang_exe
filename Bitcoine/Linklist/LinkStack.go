package main

type Node struct {
	data  interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{} //返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

//n为头节点，在n之后进行添加节点

func (n *Node) Push(data interface{}) {
	newnode := &Node{data, nil}
	newnode.pNext = n.pNext
	n.pNext = newnode //头部插入
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.pNext.data   //要弹出的数据
	n.pNext = n.pNext.pNext //头节点的指针指向下下个数据，就完成了删除下一个数据
	return value
}

func (n *Node) Length() int {
	length := 0
	for n.pNext != nil {
		n = n.pNext //节点的循环跳跃
		length++    //追加
	}
	return length
}
