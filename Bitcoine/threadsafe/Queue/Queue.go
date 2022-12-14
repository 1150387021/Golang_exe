package Queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

//新建一个队列

func NewQueue() *Queue {
	queue := &Queue{} //初始化，开辟结构体
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)
	return queue
}

//解决线程安全

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) Shift() (el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, el)
	q.len++
	return
}

func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue[0]
}
