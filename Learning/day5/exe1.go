package main

import (
	"fmt"
	"sync"
)

// Bank 银行
type Bank struct {
	balance int //余额

}

//Deposit存款
func (b *Bank) Deposit(amount int) {
	b.balance += amount
}

//查询余额
func (b *Bank) Balance() int {
	return b.balance
}
func v1() {
	b := &Bank{}
	//b.Deposit(1000)
	//b.Deposit(1000)
	//b.Deposit(1000)
	//fmt.Println(b.Balance())
	var wg sync.WaitGroup
	n := 1000
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(1000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.Balance()) //973000  968000
}

//添加互斥锁

type Bankv2 struct {
	balance int        //余额
	m       sync.Mutex //互斥锁，保护一个资源，防止竞争
}

//Deposit存款
func (b *Bankv2) Deposit(amount int) {
	b.m.Lock()         //加锁
	defer b.m.Unlock() //解锁
	b.balance += amount
}

//查询余额
func (b *Bankv2) Balance() int {
	return b.balance
}
func v2() {
	b := &Bankv2{}
	//b.Deposit(1000)
	//b.Deposit(1000)
	//b.Deposit(1000)
	//fmt.Println(b.Balance())
	var wg sync.WaitGroup
	n := 1000
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(1000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.Balance()) //973000  968000
}

//添加互斥锁并提高性能

type Bankv3 struct {
	balance int          //余额
	rwMutex sync.RWMutex //读写锁，适用场景：读多写少（微博，b站）
}

//Deposit存款
func (b *Bankv3) Deposit(amount int) {
	b.rwMutex.Lock()         //加锁
	defer b.rwMutex.Unlock() //解锁
	b.balance += amount
}

//查询余额
func (b *Bankv3) Balance() (balance int) {
	b.rwMutex.RLock()
	balance = b.balance
	b.rwMutex.RUnlock()
	return
}
func v3() {
	b := &Bankv3{}
	//b.Deposit(1000)
	//b.Deposit(1000)
	//b.Deposit(1000)
	//fmt.Println(b.Balance())
	var wg sync.WaitGroup
	n := 1000
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(1000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.Balance()) //973000  968000
}
func main() {
	//v1()
	//v2()
	v3()

}
