package main

import (
	"awesomeProject1/Learning/day4/book"
	"fmt"
)

func main() {
	book.Printtype(57, 3.14, "意义")
	b, e := book.Showbook("红楼梦", "曹雪芹")
	fmt.Println(b, e)

}
