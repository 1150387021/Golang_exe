package main

import "fmt"

type Lesson struct {
	name, target string
	spend        int
}

func s() {
	lesson1 := Lesson{
		name:   "红楼梦",
		target: "看书",
		spend:  6,
	}
	fmt.Println(lesson1)
	lesson2 := Lesson{}
	lesson2.name = "nihao"
	lesson2.target = "xuexi"
	lesson2.spend = 2
	fmt.Println(lesson2)
}

func main() {
	s()

}
