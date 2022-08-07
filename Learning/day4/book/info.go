package book

import "errors"

func Showbook(name, author string) (string, error) {
	if name == "" {
		return "", errors.New("书名不能为空")
	}
	if author == "" {
		return "", errors.New("作者不能为空")
	}
	return (name + ", 作者：" + author), nil
}
