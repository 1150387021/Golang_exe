package ArrayList

import (
	"errors"
	"fmt"
)

//小写是私有只能内部使用，大写公有可调用
//接口

type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         //抓取第几个元素
	Set(index int, newval interface{}) error    //修改数据
	checkisFull()                               //检查是否空间满了
	Insert(index int, newval interface{}) error //插入数据
	Append(newval interface{})                  //追加
	Clear()                                     //清空
	Delete(index int) error                     //删除
	String() string                             //返回字符串
}

//数据结构
type ArrayList struct {
	dataStore []interface{} //数组存储
	TheSize   int           //数组大小

}

//创建新的数组

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      //初始化结构体
	list.dataStore = make([]interface{}, 0, 10) //开辟空间10个
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize //返回数组大小
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval
	return nil
}

func (list *ArrayList) checkisFull() {
	if list.TheSize == cap(list.dataStore) {
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize) //开辟双倍内存
		copy(newdataStore, list.dataStore)                                  //拷贝
		list.dataStore = newdataStore                                       //赋值
	}
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkisFull()
	list.dataStore = list.dataStore[:list.TheSize+1] //插入数据，内存移一位
	for i := list.TheSize; i > index; i-- {          //从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval //插入数据
	list.TheSize++
	return nil
}

func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval) //追加数据
	list.TheSize++
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //跳过index追加后面的数
	list.TheSize--
	return nil
}
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
