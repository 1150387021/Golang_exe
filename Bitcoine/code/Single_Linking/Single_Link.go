package Single_Linking

import (
	"fmt"
	"strings"
)

//接口
type SingleLink interface {
	//增删查改
	GetFirstNode() *SingleLinkNode        //抓取头部节点
	InsertNodeFront(node *SingleLinkNode) //头部插入
	InsertNodeBack(node *SingleLinkNode)  //尾部插入

	InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool //在一个节点前插入
	InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool  //在一个节点后插入

	GetNodeAtIndex(index int) *SingleLinkNode //根据索引抓取指定定位置的节点
	DeleteNode(dest *SingleLinkNode) bool     //删除指定为位置的节点
	DeleteAtIndext(index int)                 //g根据指定都索引删除节点
	String() string                           //返回链表字符串
	FindString(data string)
	GetMid() *SingleLinkNode //得到链表中间节点
	ReverseList()            //反转链表
}

//定义链表的结构类型
type SingleLinkList struct {
	head   *SingleLinkNode //链表的头节点
	length int
}

//创建链表
func NewSingLeLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)
	return &SingleLinkList{head, 0}
}

//返回第一个数据节点
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

//头部插入
func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
		list.length++
	} else {
		bak := list.head
		node.pNext = bak.pNext
		bak.pNext = node
		list.length++ //插入节点，数据追加
	}
}

//尾部插入
func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
		list.length++
	} else {
		bak := list.head
		for bak.pNext != nil { //循环到链表尾端
			bak = bak.pNext
		}
		bak.pNext = node
		list.length++
	}
}

//在节点前插入节点

func (list *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	phead := list.head
	isfind := false //是否找到数据
	for phead.pNext != nil {
		if phead.pNext.value == dest { //找到
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return isfind
	} else {
		return isfind
	}
}

//在节点后插入节点
func (list *SingleLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	phead := list.head
	isfind := false //是否找到数据
	for phead.pNext != nil {
		if phead.value == dest { //找到
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return true
	} else {
		return false
	}
}

//根据索引抓取指定定位置的节点
func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.length-1 || index < 0 {
		return nil
	} else {
		phead := list.head
		for index > -1 {
			phead = phead.pNext //向后循环
			index--
		}
		return phead
	}
}

//删除指定为位置的节点
func (list *SingleLinkList) DeleteNode(dest *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	phead := list.head
	for phead.pNext != nil && phead.pNext != dest {
		phead = phead.pNext
	}
	if phead.pNext == dest {
		phead.pNext = phead.pNext.pNext
		list.length--
		return true
	} else {
		return false
	}
}

//按索引删除节点
func (list *SingleLinkList) DeleteAtIndext(index int) {
	if index > list.length-1 || index < 0 {
		return
	} else {
		phead := list.head
		for index > 0 {
			phead = phead.pNext
			index--
		}
		phead.pNext = phead.pNext.pNext
		list.length--
		return
	}
}

//返回链表字符串
func (list *SingleLinkList) String() string {
	var listString string
	p := list.head //头部节点
	for p.pNext != nil {
		listString += fmt.Sprintf("%v-->", p.pNext.value)
		p = p.pNext //循环
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}

//找字符串
func (list *SingleLinkList) FindString(data string) {
	phead := list.head.pNext //指定头部
	for phead.pNext != nil { //循环所有数据
		if strings.Contains(phead.value.(string), data) { //包含
			fmt.Println(phead.value)
		}
		phead = phead.pNext
	}
}

//获得中间节点

func (list *SingleLinkList) GetMid() *SingleLinkNode {
	if list.head.pNext == nil {
		return nil
	} else {
		phead1 := list.head
		phead2 := list.head
		for phead2 != nil && phead2.pNext != nil {
			phead1 = phead1.pNext
			phead2 = phead2.pNext.pNext
		}
		return phead1 //返回中间节点
	}
}

//反转链表

func (list *SingleLinkList) ReverseList() {
	if list.head == nil || list.head.pNext == nil {
		return //链表为空或链表只有一个节点
	} else {
		var pre *SingleLinkNode                   //前节点
		var cur *SingleLinkNode = list.head.pNext //当前节点
		for cur != nil {
			curNext := cur.pNext //后续节点
			cur.pNext = pre      //反转第一步
			pre = cur            //持续推进
			cur = curNext        //持续推进
		}
		list.head.pNext = pre
	}
}
