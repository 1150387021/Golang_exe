package main

import (
	"fmt"
	"testing"
)

/*
有向图的实现：
定义了三个结构体， Graph是图的数据结构，它包含边数、顶点数和顶点的集合；
Vertex是顶点数据结构，包含顶点的数据内容和边的头指针，
Edge边的结构体定义了与此顶点相连的另一个顶点在图的顶点集合中的位置和下一个边的地址。
图是顶点和边的集合，adj对象就是顶点的集合，每一个顶点内部的e就是通过链表实现的边的集合。这也就实现了图。
*/

type Graph struct {
	edgnum int
	vexnum int
	adj    []Vertex
}

type Vertex struct {
	Data string
	e    *Edge
}

type Edge struct {
	ivex int
	next *Edge
}

//在图中插入顶点，这个比较简单，加入顶点的集合即可
func (this *Graph) InsertVertex(v Vertex) {
	this.adj = append(this.adj, v)
}

//定位顶点的索引位置
func (this *Graph) get_position(verData string) int {
	for i := 0; i < this.vexnum; i++ {
		if this.adj[i].Data == verData {
			return i
		}
	}
	return -1
}

//在图中插入边，需要指出相连的两个顶点的信息，然后把边的信息插入到链表的最后
func (this *Graph) InsertEdge(v1, v2 Vertex) {
	var p *Edge = this.adj[this.get_position(v1.Data)].e
	if p == nil {
		// fmt.Println("nil...", v1.Data, v2.Data)
		this.adj[this.get_position(v1.Data)].e = NewEdge(this.get_position(v2.Data))
	} else {
		for ; p.next != nil; p = p.next {
		}
		p.next = NewEdge(this.get_position(v2.Data))
		// fmt.Println("append...", v1.Data, v2.Data)
	}
}

//获取某个顶点的邻接顶点，找到顶点的边集合，就可以通过边集合找到与它相连的顶点了。
//也可以通过这个方法计算得到顶点的出度（Out Degree）

func (this *Graph) Adjacent(v Vertex) []Vertex {
	res := make([]Vertex, 0)
	p := this.adj[this.get_position(v.Data)].e
	for ; p != nil; p = p.next {
		res = append(res, this.adj[p.ivex])
	}
	return res
}

func (this *Graph) OutDegree(v Vertex) int {
	res := 0
	p := this.adj[this.get_position(v.Data)].e
	for p != nil {
		res++
		p = p.next
	}
	return res
}

//计算顶点的入度相比出度有点麻烦，因为图本身保存的边的集合是边的起始位置，结束位置，也就是入度，只能遍历整个边集合来计算

func (this *Graph) InDegree(v Vertex) int {
	res := 0
	pos := this.get_position(v.Data)
	for _, a := range this.adj {
		p := a.e
		for p != nil {
			if pos == p.ivex {
				res++
			}
			p = p.next
		}
	}
	return res
}

/*
广度优先遍历，这个遍历类似于层序遍历，每遍历一层，需要记住当前层的节点，然后与遍历当前层相连的节点，如此实现遍历。
需要一个队列来记住当前层，先进先出。我是通过Golang的slice实现的简单队列。还有一个问题，就是需要防止回路，
也就是说，一个节点不能遍历两次。这里用了Golang内置的map实现，随机存储、随机查找。遍历还要防止非完全图的情况，
如果只有结点没有边还需要处理，所以这里用for循环来处理
*/

func (this *Graph) Bfs() {
	res := map[int]Vertex{}
	for _, a := range this.adj {
		Q := []Vertex{a}
		for len(Q) != 0 {
			u := Q[0]
			Q = Q[1:]
			if _, ok := res[this.get_position(u.Data)]; !ok {
				Q = append(Q, this.Adjacent(u)...)
				res[this.get_position(u.Data)] = u
				fmt.Printf("%s ", u.Data)
			}
		}
	}
	fmt.Printf("\n")
}

//深度优先遍历，遍历完一个节点，接着遍历这个节点的这个节点相连的节点，直到结束。用递归非常简单
func (this *Graph) Dfs() {
	res := map[int]Vertex{}
	for _, a := range this.adj {
		this.dfs(a, res)
	}
	fmt.Printf("\n")
}

func (this *Graph) dfs(u Vertex, res map[int]Vertex) {
	if _, ok := res[this.get_position(u.Data)]; !ok {
		res[this.get_position(u.Data)] = u
		fmt.Printf("%s ", u.Data)
		p := u.e
		for p != nil {
			if _, ok := res[p.ivex]; !ok {
				this.dfs(this.adj[p.ivex], res)
			}
			p = p.next
		}
	}
}

//完整的测试代码，大Golang的单元测试工具，直接go test即可
func TestMain(t *testing.T) {
	g := create_example_lgraph()
	g.Print()
	println("Bfs.............")
	g.Bfs()
	println("Dfs.............")
	g.Dfs()
	if g.InDegree(Vertex{Data: "E"}) != 2 {
		t.Error("indegree of E wanted 2 bug got", g.InDegree(Vertex{Data: "E"}))
	}
	if g.OutDegree(Vertex{Data: "E"}) != 2 {
		t.Error("outdegree of E wanted 2 bug got", g.OutDegree(Vertex{Data: "E"}))
	}
}

func create_example_lgraph() *Graph {
	vexs := []string{"A", "B", "C", "D", "E", "F", "G"}
	edges := [][2]string{
		{"A", "B"},
		{"B", "C"},
		{"B", "E"},
		{"B", "F"},
		{"C", "E"},
		{"D", "C"},
		{"E", "B"},
		{"E", "D"},
		{"F", "G"}}
	vlen := len(vexs)
	elen := len(edges)

	pG := Create()
	// 初始化"顶点数"和"边数"
	pG.vexnum = vlen
	pG.edgnum = elen
	// 初始化"邻接表"的顶点
	for i := 0; i < pG.vexnum; i++ {
		pG.InsertVertex(*NewVertex(vexs[i]))
	}
	// 初始化"邻接表"的边
	for i := 0; i < pG.edgnum; i++ {
		// 读取边的起始顶点和结束顶点
		pG.InsertEdge(*NewVertex(edges[i][0]), *NewVertex(edges[i][1]))
	}
	return pG
}
