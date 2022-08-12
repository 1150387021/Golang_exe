package main

import (
	"awesomeProject1/Bitcoine/ArrayList"
	"awesomeProject1/Bitcoine/CircleQueue"
	"awesomeProject1/Bitcoine/Queue"
	"awesomeProject1/Bitcoine/StackArray"
	Queue2 "awesomeProject1/Bitcoine/threadsafe/Queue"
	"fmt"
)

func main1() {
	list := ArrayList.NewArrayList()
	list.Append("a2")
	list.Append("b1")
	list.Append(3)
	list.Append("e")
	fmt.Println(list)
	fmt.Println(list.TheSize)
}
func main2() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append(3)
	list.Append("e4")
	for i := 0; i < 10; i++ {
		list.Insert(1, "x5")
	}

	fmt.Println(list)
}
func main3() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	myq := Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue("b3")
	myq.EnQueue("c4")
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
}
func main4() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var myq CircleQueue.CircleQueue
	CircleQueue.InitQueue(&myq)
	CircleQueue.EnQueue(&myq, 1)
	CircleQueue.EnQueue(&myq, 2)
	CircleQueue.EnQueue(&myq, 3)
	CircleQueue.EnQueue(&myq, 4)
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
}
func main5() {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main6() {
	h := Queue2.NewMin()
	h.Insert(Queue2.Int(8))
	h.Insert(Queue2.Int(7))
	h.Insert(Queue2.Int(6))
	h.Insert(Queue2.Int(5))
	h.Insert(Queue2.Int(4))
	h.Insert(Queue2.Int(3))
	h.Insert(Queue2.Int(2))
	h.Insert(Queue2.Int(1))
	fmt.Println(h.Extract().(Queue2.Int))
}
