package main

import (
	"awesomeProject1/Bitcoine/threadsafe/Queue"
	"fmt"
)

func main() {
	h := Queue.NewMinPriorityQueue()
	h.Insert(*Queue.NewPriorityItem(101, 16))
	h.Insert(*Queue.NewPriorityItem(102, 17))
	h.Insert(*Queue.NewPriorityItem(103, 13))
	h.Insert(*Queue.NewPriorityItem(104, 15))
	h.Insert(*Queue.NewPriorityItem(105, 14))
	h.Insert(*Queue.NewPriorityItem(106, 18))
	h.Insert(*Queue.NewPriorityItem(107, 19))
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
}
