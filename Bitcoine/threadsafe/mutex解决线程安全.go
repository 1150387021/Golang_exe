package main

import (
	"fmt"
	"sync"
	"time"
)

var money1 int = 0
var lock *sync.RWMutex = new(sync.RWMutex) //初始化

func add1(pint *int) {
	lock.Lock()
	for i := 0; i < 1000; i++ {
		*pint++
	}
	lock.Unlock()
}

//线程安全，多个线程访问同一个资源会产生资源竞争，从而出错
func main() {
	for i := 0; i < 100; i++ {
		go add1(&money1)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(money1)
}
