package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type uid struct {
	list   int
	id     int
	like   int
	school string
	is_985 string
}

const N = 637

func main() {
	allstrs := make([]uid, N, N) //初始化数组

	//文件读取
	path := "/Users/st/space/数据/data.csv"
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件
	i := 0
	br := bufio.NewReader(fi)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr := string(line)                     //读取，转化为字符串
		lines := strings.Split(linestr, ",")        //字符串切割
		allstrs[i].list, _ = strconv.Atoi(lines[0]) //字符串转整数
		allstrs[i].id, _ = strconv.Atoi(lines[1])
		allstrs[i].like, _ = strconv.Atoi(lines[2])
		allstrs[i].school = lines[3]
		allstrs[i].is_985 = lines[4]

		i++
	}
	fmt.Println(i)
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 2)

	for {
		fmt.Println("请输入要查询的用户ID：")
		var inputstr int
		fmt.Scanf("%d", &inputstr)

		starttime := time.Now()
		for j := 0; j < N; j++ {
			if allstrs[j].id == inputstr {
				fmt.Println(j, allstrs[j].school, allstrs[j].is_985)
			}
		}
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
