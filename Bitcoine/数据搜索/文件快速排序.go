package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func f1() {
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
	save := bufio.NewWriter(savefile)

	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		//fmt.Println(string(line))
		linestr := string(line) //读取，转化为字符串
		//if strings.Contains(linestr, "北京大学") { //搜索特定子字符串的行
		//	fmt.Println(linestr)
		//
		//}
		mystr := strings.Split(linestr, ",") //字符串切割
		//fmt.Println(mystr[1], mystr[3], mystr[4])
		fmt.Fprintln(save, mystr[3]) //字符串保存
	}
	save.Flush() //刷新
}

func main() {
	start := time.Now()
	const N = 300
	f1()
	fmt.Println("一共用了", time.Since(start))

}
