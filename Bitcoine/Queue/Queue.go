package Queue

type MyQueue interface {
	Size() int                //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个
	IsEmpty() bool            //是否为空
	EnQueue(data interface{}) //入队
	DeQueue() interface{}     //出队
	Clear()                   //清空
}

type Queue struct {
	dataStore []interface{} //队列的数据存储
	TheSize   int           //队列大小
}

func NewQueue() *Queue {
	myqueue := new(Queue)                      //初始化，开辟结构体
	myqueue.dataStore = make([]interface{}, 0) //
	myqueue.TheSize = 0
	return myqueue
}

func (myq *Queue) Size() int {
	return myq.TheSize
}

func (myq *Queue) Front() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[0]
}

func (myq *Queue) End() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[myq.Size()-1]
}

func (myq *Queue) IsEmpty() bool {
	return myq.TheSize == 0
}

func (myq *Queue) EnQueue(data interface{}) {
	myq.dataStore = append(myq.dataStore, data)
	myq.TheSize++
}

func (myq *Queue) DeQueue() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	data := myq.dataStore[0]
	if myq.Size() > 1 {
		myq.dataStore = myq.dataStore[1:myq.Size()]
	}
	myq.TheSize--
	return data //返回数据
}

func (myq *Queue) Clear() {
	myq.dataStore = make([]interface{}, 0)
	myq.TheSize = 0
}
