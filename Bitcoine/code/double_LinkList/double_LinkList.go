package double_LinkList

import "fmt"

//双链表的基本结构
type DoubleLinkList struct {
	head   *DoubleLinkNode
	length int
}

//创建新的双链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, 0}
}

//返回链表长度
func (dlist *DoubleLinkList) Getlength() int {
	return dlist.length
}

//返回第一个节点
func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next == nil {
		node.next = nil   //nil
		phead.next = node //只有一个节点直接连接上
		node.prev = phead //上一个节点
		dlist.length++
	} else {
		phead.next.prev = node //标记上一个节点
		node.next = phead.next //下一个节点

		phead.next = node //标记头部节点的下一个
		node.prev = phead //
		dlist.length++

	}
}

func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next == nil {
		node.next = nil   //nil
		phead.next = node //只有一个节点直接连接上
		node.prev = phead //上一个节点
		dlist.length++
	} else {
		for phead.next != nil { //循环下去
			phead = phead.next
		}
		phead.next = node //后缀
		node.prev = phead //前缀
		dlist.length++
	}
}
func (dlist *DoubleLinkList) String() string {
	var listString1 string
	var listString2 string
	phead := dlist.head
	for phead.next != nil {
		//正向循环
		listString1 += fmt.Sprintf("%v-->", phead.next.value)
		phead = phead.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	for phead != dlist.head {
		//反向循环
		listString2 += fmt.Sprintf("<--%v", phead.value)
		phead = phead.prev
	}
	listString1 += fmt.Sprintf("nil")
	return listString1 + listString2 + "\n"

}
func (dlist *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest {
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next.next != nil {
			phead.next.next.prev = node
		}
		node.next = phead.next.next
		phead.next.next = node
		node.prev = phead.next

		//dlist.head = phead
		dlist.length++
		return true
	} else {
		return false
	}
}
func (dlist *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest {
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next != nil {
			phead.next.prev = node
		}
		node.next = phead.next
		node.prev = phead
		phead.next = node

		//dlist.head = phead
		dlist.length++
		return true
	} else {
		return false
	}
}
