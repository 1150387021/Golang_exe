package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func GetALL(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fulldir)    //追加路径
			files, _ = GetALL(fulldir, files) //文件夹递归处理
		} else {
			fulldir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fulldir)    //追加路径
		}
	}
	return files, nil
}

func main() {
	path := "/Users/st/Goland/Projects/Algo_models" //路径
	files := []string{}                             //数组字符串
	files, _ = GetALL(path, files)                  //抓取所有文件
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
