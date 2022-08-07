package main

import (
	"awesomeProject1/Bitcoine/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/Users/st/Goland/Projects/Algo_models" //路径
	files := []string{}                             //数组字符串
	mystack := StackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)
		read, err := ioutil.ReadDir(path)
		if err != nil {
			errors.New("文件夹不可读取")
		}
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir)
				mystack.Push(fulldir)
			} else {
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
