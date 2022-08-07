package main

import "log"

type Edge struct {
	index interface{}
	next  *Edge
}

type VerNode struct {
	data  interface{}
	edges *Edge
}

type Graph struct {
	VerNodes []VerNode
	edgenum  int
	vernum   int
}

//接口
type mygraph interface {
	findindex(verval interface{})
	AddEdge(verval1 interface{}, verval2 interface{})
}

//找顶点的索引位置
func (g *Graph) findindex(verval interface{}) interface{} {
	for i := 0; i < g.vernum; i++ {
		if g.VerNodes[i].data == verval {
			return g.VerNodes[i].data
		}
	}
	return "No"
}
func (g *Graph) AddEdge(verval1 interface{}, verval2 interface{}) {
	index1, index2 := g.findindex(verval1), g.findindex(verval2)
	if index1 == "No" || index2 == "No" {
		log.Fatalln("非法顶点值")
	}
	edge2 := &Edge{index: index2}
	edge2.next = g.VerNodes[index1].edges
	g.VerNodes[index1].edges = edge2
	g.edgenum--
}

func main() {

}
