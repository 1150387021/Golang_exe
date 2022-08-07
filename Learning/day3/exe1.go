package main

import (
	"errors"
	"fmt"
)

func printtype(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "type is int")
		case string:
			fmt.Println(arg, "type is string")
		case float64:
			fmt.Println(arg, "type is float64")
		default:
			fmt.Println(arg, "is an unknow type")

			var s []string
			s = append(s, []string{"nihoa", "小猫", "加油"}...)
			fmt.Println(s)

		}
	}
}
func Showbook(name, author string) (string, error) {
	if name == "" {
		return "", errors.New("书名不能为空")
	}
	if author == "" {
		return "", errors.New("作者不能为空")
	}
	return (name + ", 作者：" + author), nil
}

func main() {
	//printtype(57, 3.14, "意义")
	b, e := Showbook("红楼梦", "曹雪芹")
	fmt.Println(b, e)

}
