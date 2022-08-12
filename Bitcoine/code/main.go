package main

import (
	"awesomeProject1/Bitcoine/code/Single_Linking"
	"awesomeProject1/Bitcoine/code/double_LinkList"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main1() {
	list := Single_Linking.NewSingLeLinkList()
	node1 := Single_Linking.NewSingleLinkNode(1)
	node2 := Single_Linking.NewSingleLinkNode(2)
	node3 := Single_Linking.NewSingleLinkNode(3)
	list.InsertNodeFront(node1)
	fmt.Println(list)
	list.InsertNodeFront(node2)
	fmt.Println(list)
	list.InsertNodeFront(node3)
	fmt.Println(list)
}

func main2() {
	list := Single_Linking.NewSingLeLinkList()
	node1 := Single_Linking.NewSingleLinkNode(1)
	node2 := Single_Linking.NewSingleLinkNode(2)
	node3 := Single_Linking.NewSingleLinkNode(3)
	list.InsertNodeBack(node1)
	fmt.Println(list)
	list.InsertNodeBack(node2) //尾插法
	fmt.Println(list)
	list.InsertNodeBack(node3)
	fmt.Println(list)

	node4 := Single_Linking.NewSingleLinkNode(4)
	//list.InsertNodeValueBack(2, node4)
	list.InsertNodeValueFront(2, node4) //指定位置嵌入节点
	fmt.Println(list)
	list.ReverseList() //反转链表
	fmt.Println("反转：", list)

	fmt.Println(list.GetNodeAtIndex(3)) //根据索引查找节点

	list.DeleteNode(node4) //删除节点
	fmt.Println(list)
	fmt.Println(list.GetMid()) //获取中间节点

	list.DeleteAtIndext(1) //删除索引为1的节点
	fmt.Println(list)

}

//采用链表构建数据，并顺序查询
func main3() {
	const N = 637
	list := Single_Linking.NewSingLeLinkList()
	path := "/Users/st/space/数据/data.csv"
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件
	br := bufio.NewReader(fi)
	//文件写入
	//path2 := "/Users/st/space/数据/write.txt"
	//savefile, _ := os.Create(path2)
	//defer savefile.Close()
	//save := bufio.NewWriter(savefile)
	i := 0
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		//fmt.Println(string(line))
		linestr := string(line)                              //读取，转化为字符串
		nodestr := Single_Linking.NewSingleLinkNode(linestr) //创建节点
		list.InsertNodeFront(nodestr)                        //把数据插入链表
		i++

	}
	fmt.Println(i, "内存载入完成")
	for {
		fmt.Println("请输入要查询的用户：")
		var inputstr string
		fmt.Scanf("%s", &inputstr)

		starttime := time.Now()
		list.FindString(inputstr)
		fmt.Println("一共用了", time.Since(starttime))

	}
}

/*
54	19489141	0	华中科技大学同济医学院	TRUE
55	19498675	1	信阳师范学院	FALSE
56	19498774	1	吉林建筑大学	FALSE
57	19501262	0	湖南工艺美术职业学院	FALSE
58	19521230	0	哈尔滨师范学院	FALSE
59	19525934	1	西北师范大学	FALSE
60	19529213	1	广西大学行健文理学院	FALSE
61	19557614	1	弗吉尼亚大学	FALSE
62	19560802	1	新疆财经大学	FALSE
63	19573166	1	辽宁师范大学	FALSE
64	19574818	1	燕京理工学院	FALSE
65	19581637	0	北京交通大学	FALSE
66	19585262	1	北京第二外国语学院	FALSE
*/

func main4() {
	dlist := double_LinkList.NewDoubleLinkList()
	node1 := double_LinkList.NewDoubleLinkNode(1)
	node2 := double_LinkList.NewDoubleLinkNode(2)
	node3 := double_LinkList.NewDoubleLinkNode(3)
	node4 := double_LinkList.NewDoubleLinkNode(4)
	node5 := double_LinkList.NewDoubleLinkNode(5)
	dlist.InsertHead(node1)
	dlist.InsertHead(node2)
	dlist.InsertHead(node3)
	dlist.InsertHead(node4)
	dlist.InsertHead(node5)
	node6 := double_LinkList.NewDoubleLinkNode(6)
	node7 := double_LinkList.NewDoubleLinkNode(7)
	//dlist.InsertValueHead(node3, node6)
	//dlist.InsertValueBack(node3, node7)
	dlist.InsertValueBackByValue(4, node7)
	dlist.InsertValueHeadByValue(3, node6)
	fmt.Println(dlist.String())
	//dlist.DeleteNode(node4)
	dlist.DeleteNodeAtIndex(2)
	fmt.Println(dlist.String())

}

func main() {
	const N = 637
	dlist := double_LinkList.NewDoubleLinkList()
	path := "/Users/st/space/数据/data.csv"
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件
	br := bufio.NewReader(fi)
	//文件写入
	//path2 := "/Users/st/space/数据/write.txt"
	//savefile, _ := os.Create(path2)
	//defer savefile.Close()
	//save := bufio.NewWriter(savefile)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		//fmt.Println(string(line))
		linestr := string(line)                            //读取，转化为字符串
		node := double_LinkList.NewDoubleLinkNode(linestr) //新建节点
		dlist.InsertHead(node)                             //插入节点数据

	}
	fmt.Println("内存载入完成", dlist.Getlength())
	for {
		fmt.Println("请输入要查询的用户：")
		var inputstr string
		fmt.Scanf("%s", &inputstr)

		starttime := time.Now()
		dlist.Findstring(inputstr)
		fmt.Println("一共用了", time.Since(starttime))

	}
}

/*
54	19489141	0	华中科技大学同济医学院	TRUE
55	19498675	1	信阳师范学院	FALSE
56	19498774	1	吉林建筑大学	FALSE
57	19501262	0	湖南工艺美术职业学院	FALSE
58	19521230	0	哈尔滨师范学院	FALSE
59	19525934	1	西北师范大学	FALSE
60	19529213	1	广西大学行健文理学院	FALSE
61	19557614	1	弗吉尼亚大学	FALSE
62	19560802	1	新疆财经大学	FALSE
63	19573166	1	辽宁师范大学	FALSE
64	19574818	1	燕京理工学院	FALSE
65	19581637	0	北京交通大学	FALSE
66	19585262	1	北京第二外国语学院	FALSE
*/
