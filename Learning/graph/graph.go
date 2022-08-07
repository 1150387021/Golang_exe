package main

import (
	"fmt"
	"log"
)

// 图边结构体

type EdgeNode struct {
	index int       // 顶点的图数组中的索引值
	next  *EdgeNode //边连接的下一个顶点信息，边是两个顶点之间的连接
}

// 图顶点结构体

type VerNode struct {
	data  int       // 顶点存储的数据，如果是很多数据，这里可以指向一个结构体的指针
	edges *EdgeNode // 结点的边
}

func (vn VerNode) String() string {
	return fmt.Sprintf("VerNode{data: %d, edges: %p}", vn.data, vn.edges)
}

// 图的结构体

type Graph struct {
	VerNodes []VerNode
	EdgeCnt  int
	VerCnt   int
}

// 定位顶点的索引位置
func (g *Graph) findIndex(verValue int) int {
	for i := 0; i < g.VerCnt; i++ {
		if g.VerNodes[i].data == verValue {
			return i
		}
	}
	return -1
}

func (g *Graph) AddEdge(verValue1 int, verValue2 int) {
	// 先定位
	index1, index2 := g.findIndex(verValue1), g.findIndex(verValue2)
	if index1 == -1 || index2 == -1 {
		log.Fatalln("非法顶点值")
	}

	// 准备好边节点
	//edgeNode1 := &EdgeNode{index: index1}
	edgeNode2 := &EdgeNode{index: index2} //边节点2，顶点1去连接2的边节点

	// 头插法，这样不用遍历节点链表，节省时间
	edgeNode2.next = g.VerNodes[index1].edges //顶点1去连接边节点2，顶点1的边连到边节点2的指针
	g.VerNodes[index1].edges = edgeNode2      //边节点2赋值给顶点1的边
	//实现无向图的双向连接
	//edgeNode1.next = g.VerNodes[index2].edges
	//g.VerNodes[index2].edges = edgeNode1
	g.EdgeCnt += 1
}

// 普通遍历

func (g *Graph) Show() {
	for i := 0; i < g.VerCnt; i++ {
		head := g.VerNodes[i].edges
		print(g.VerNodes[i].data)
		for head != nil {
			print("->")
			print(g.VerNodes[head.index].data)
			head = head.next
		}
		print("\n")
	}
}

// 广度优先遍历

func (g *Graph) BFSShow(node *VerNode) {
	// 用以记录已遍历的节点索引
	visited := make(map[*VerNode]struct{})
	queue := make([]*VerNode, 0)
	queue = append(queue, node)
	for len(queue) != 0 {
		// 模拟出队列，从头取出一个，然后队列就少了一个元素
		currNode := queue[0]
		queue = queue[1:]
		// 如果已遍历，则继续循环
		if _, ok := visited[currNode]; ok {
			continue
		}
		fmt.Print("->", currNode.data)
		visited[currNode] = struct{}{}
		head := currNode.edges
		for head != nil {
			queue = append(queue, &g.VerNodes[head.index])
			head = head.next
		}
	}
}
func (g *Graph) BFSShow2(node *VerNode) {
	// 用以记录已遍历的节点索引
	visited := make(map[*VerNode]struct{})
	queue := make([]*VerNode, 0)
	queue = append(queue, node)
	for len(queue) != 0 {
		// 模拟出队列，从头取出一个，然后队列就少了一个元素
		currNode := queue[0]
		queue = queue[1:]
		// 如果已遍历，则继续循环
		if _, ok := visited[currNode]; ok {
			continue
		}
		fmt.Print("->", currNode.data)
		visited[currNode] = struct{}{}
		head := currNode.edges
		for head != nil {
			queue = append(queue, &g.VerNodes[head.index])
			head = head.next
		}
	}
}

// 深度优先遍历

func (g *Graph) DFSShow(node *VerNode) {
	// 也需要一个变量暂时保存是否访问过
	visited := make(map[*VerNode]struct{})
	// 用数组实现一个简单栈，先压入栈中
	stack := make([]*VerNode, 0)
	stack = append(stack, node)
	for len(stack) != 0 {
		maxIndex := len(stack) - 1
		// 先进后出，模拟出栈，取栈中的最后一个数据
		currNode := stack[maxIndex]
		stack = stack[:maxIndex]
		if _, ok := visited[currNode]; ok {
			continue
		}
		fmt.Print("->", currNode.data)
		visited[currNode] = struct{}{}
		head := currNode.edges
		for head != nil {
			stack = append(stack, &g.VerNodes[head.index])
			head = head.next
			// time.Sleep(time.Second)
		}
		// time.Sleep(time.Second)
	}
}

func NewGraph(vCnt int) *Graph {
	vNodes := make([]VerNode, 0, vCnt)
	for i := 1; i <= vCnt; i++ {
		vNodes = append(vNodes, VerNode{data: i})
	}
	return &Graph{
		EdgeCnt:  0,
		VerCnt:   vCnt,
		VerNodes: vNodes,
	}
}

func TestNewGraph() {
	vCnt := 8
	graph := NewGraph(vCnt)
	// 模拟上图的图连接，增加节点之间的边
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(5, 8)
	graph.AddEdge(4, 8)
	graph.AddEdge(3, 6)
	graph.AddEdge(3, 7)
	graph.AddEdge(6, 7)

	fmt.Println("顶点连接的边：")
	graph.Show()
	fmt.Println("广度优先遍历：")
	graph.BFSShow(&graph.VerNodes[0])
	fmt.Println()
	fmt.Println("深度优先遍历：")
	graph.DFSShow(&graph.VerNodes[0])
}
func main() {
	TestNewGraph()
}
