package main

import (
	"fmt"
)

// 邻接表
type Vertex struct {
	Key      string
	Parents  []*Vertex
	Children []*Vertex
	Value    interface{}
}

type DAG struct {
	Vertexes []*Vertex
}

func (dag *DAG) AddVertex(v *Vertex) {
	dag.Vertexes = append(dag.Vertexes, v)
}

func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)

	//to.Parents = append(from.Parents, from)
}

func (dag *DAG) BFS(root *Vertex) {
	//q := queue.New()
	q := make([]*Vertex, 0)
	visitMap := make(map[string]bool)
	visitMap[root.Key] = true

	q = append(q, root)

	for {
		if len(q) == 0 {
			fmt.Println("完成遍历")
			break
		}
		current := q[0]
		q = q[1:]
		//fmt.Println("当前顶点：", current.Key)
		for _, v := range current.Children {
			fmt.Printf("%v->%s\n", current.Key, v.Key)
			if v.Key == root.Key {
				panic("产生回路！")
			}
			if _, ok := visitMap[v.Key]; !ok {
				visitMap[v.Key] = true
				fmt.Println("广度优先遍历:", v.Key)
				q = append(q, v)
			}
		}
	}

}

func (dag *DAG) DFS(root *Vertex) {
	stack := []*Vertex{root}

	visitMap := make(map[string]bool)
	visitMap[root.Key] = true

	for i := 0; i < 10; i++ {
		if len(stack) == 0 {
			fmt.Println("完成遍历！")
			break
		}
		if len(stack)-1 < 0 {
			panic("unexpected")
		}
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, v := range current.Children {
			fmt.Printf("%v->%s\n", current.Key, v.Key)
			if v.Key == root.Key {
				panic("产生环路！")
			}
			if _, ok := visitMap[v.Key]; !ok {
				visitMap[v.Key] = true
				fmt.Println("深度优先遍历：", v.Key)
				if v.Key == root.Key {
					panic("产生环路！！")
				}
				stack = append(stack, v)
			}
		}
	}

}

func main() {

	dag := &DAG{}
	v1 := &Vertex{Key: "1"}
	v2 := &Vertex{Key: "2"}
	v3 := &Vertex{Key: "3"}
	v4 := &Vertex{Key: "4"}
	v5 := &Vertex{Key: "5"}

	// 对于有环图，从root点出发，最终会回到root
	// non dag
	//    5
	//   >
	//  /
	// 1 <----2
	//  \   .>  \
	//   > /     >
	//    3----  >4
	dag.AddEdge(v1, v5)
	dag.AddEdge(v2, v1)
	dag.AddEdge(v1, v3)
	dag.AddEdge(v3, v4)
	dag.AddEdge(v3, v2)
	dag.AddEdge(v2, v4)
	dag.BFS(v1)
	//dag.DFS(v1)
}
