package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strings"
	"time"
)

func f() {
	//t1 := time.Now()
	const N = 637
	allstrs := make([]string, N, N) //开辟内存空间
	fmt.Println(len(allstrs))
	time.Sleep(time.Second * 10)
	//time.Since(t1)

	//文件读取
	path := "/Users/st/space/数据/data.csv"
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件
	br := bufio.NewReader(fi)
	//文件写入
	path2 := "/Users/st/space/数据/write.txt"
	savefile, _ := os.Create(path2)
	defer savefile.Close()
	//save := bufio.NewWriter(savefile)
	i := 0
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		//fmt.Println(string(line))
		linestr := string(line) //读取，转化为字符串
		//mystrs := strings.Split(linestr, ",") //字符串切割
		allstrs[i] = linestr

	}
}

func main() {
	f()
}
