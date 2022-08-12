package main

import (
	"awesomeProject1/Bitcoine/threadsafe/Queue"
	"fmt"
)

func main() {
	h := Queue.NewMin()
	h.Insert(Queue.Int(8))
	h.Insert(Queue.Int(5))
	h.Insert(Queue.Int(6))
	h.Insert(Queue.Int(12))
	h.Insert(Queue.Int(13))
	h.Insert(Queue.Int(9))
	h.Insert(Queue.Int(4))
	h.Insert(Queue.Int(3))
	h.Insert(Queue.Int(2))
	h.Insert(Queue.Int(7))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
}
