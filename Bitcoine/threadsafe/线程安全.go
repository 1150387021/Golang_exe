package main

import (
	"fmt"
	"time"
)

var money int = 0

func add(pint *int) {
	for i := 0; i < 1000; i++ {
		*pint++
	}
}

//线程安全，多个线程访问同一个资源会产生资源竞争，从而出错
func main() {
	for i := 0; i < 100; i++ {
		go add(&money)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(money)
}
