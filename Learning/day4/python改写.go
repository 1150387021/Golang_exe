package main

type Vertex struct {
	id          int
	connectedTo map[int]float32
}

func (v *Vertex) addNeighbor(nbr int, weight float32) {
	v.connectedTo[nbr] = weight
}
func (v *Vertex) getid() int {
	return v.id
}
func (v *Vertex) getweight(nbr int) float32 {
	return v.connectedTo[nbr]
}

//func (v *Vertex) getconnections() int {
//	for k, _ := range v.connectedTo {
//		return k
//	}
//}

type Graph struct {
	vertlist    map[int]float32
	numvertices int
}

//func (g *Graph) addVertex(key int) int {
//	g.numvertices += 1
//	newVertex := Vertex(key, nil)
//	g.vertlist[key] = newVertex
//	return newVertex
//
//}

func main() {

}
