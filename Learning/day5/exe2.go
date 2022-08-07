package main

import (
	"fmt"
	"os"
)

type errorstring struct {
	s string
}

func New(text string) error {
	return &errorstring{s: text}
}
func (e *errorstring) Error() string {
	return e.s
}

func findfile() {
	file, err := os.Open("/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name(), "opened successfully")
}

func area(a, b int) (int, error) {
	if a < 0 || b < 0 {
		//return 0, errors.New("计算错误，长度或宽度，不能小于0")
		return 0, fmt.Errorf("计算错误，长度%d或宽度%d，不能小于0", a, b)
	}
	return a * b, nil
}
func compute() {
	a := 100
	b := -10
	r, err := area(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("area=", r)
}
func main() {
	//findfile()
	compute()

}
