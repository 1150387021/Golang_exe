package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func f3() {
	const N = 637
	allstrs := make([]string, N, N) //开辟内存空间
	fmt.Println(len(allstrs))
	path := "/Users/st/space/数据/data.csv"
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("路径有问题，文件读取失败", err)
		return
	}

	defer fi.Close()
	br := bufio.NewReader(fi)
	i := 0
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr := string(line) //读取，转化为字符串
		//lines := strings.Split(linestr, ",") //字符串切割
		allstrs[i] = linestr
		i++

	}
	for {
		fmt.Println("请输入要查询的大学，默认北大")
		var inputstr string = "北京大学"
		fmt.Scanln(&inputstr)

		start := time.Now()
		for j := 0; j < N; j++ {
			if strings.Contains(allstrs[j], inputstr) {
				fmt.Println(allstrs[j])
			}
		}
		fmt.Println("一共用了", time.Since(start))
	}
}

func main() {
	f3()
}
