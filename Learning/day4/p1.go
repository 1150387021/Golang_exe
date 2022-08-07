package main

import "fmt"

type Lesson struct {
	name   string
	target string
}

func (l Lesson) printinfo() {
	fmt.Println("name:", l.name)
	fmt.Println("target:", l.target)

}

type worker struct {
	name string
	book string
	by   string //方式
}

type study interface {
	learn()
}

type student struct {
	name string
	book string
}

func (s student) learn() {
	fmt.Printf("%s 在读 %s\n", s.name, s.book)
}
func (w *worker) learn() {
	fmt.Printf("%s 在读 %s 通过%s\n", w.name, w.book, w.by)
}

//showinterface接口看成type和value的组合，type时接口底层的具体类型，value是具体类型的值

func showinterface(s study) {
	fmt.Printf("接口类型：%T\n,接口值为%v\n", s, s)
}
func main() {
	//l := Lesson{
	//	name:   "go教程13讲",
	//	target: "学习编程",
	//}
	//l.printinfo()
	s1 := student{
		name: "张三",
		book: "红楼梦",
	}
	s1.learn()
	worker1 := worker{
		name: "老赵",
		book: "母猪的产后护理",
		by:   "春晚直播",
	}
	s2 := &worker1
	s2.learn()
	fmt.Println("-----------")

	var st1 study
	var st2 study
	student2 := student{
		name: "里斯",
		book: "三国演义",
	}
	st1 = student2
	st1.learn()

	worker2 := worker{
		name: "康师傅",
		book: "红烧牛肉面",
		by:   "卖广告",
	}
	st2 = &worker2
	st2.learn()
	fmt.Println("-----------")

	var s study
	s = student2
	showinterface(s)
	s.learn()
	fmt.Println("----------")
	x := make([]interface{}, 3)
	x[0] = "孙悟空"
	x[1] = 3.14
	x[2] = []int{1, 2, 3}
	for _, item := range x {
		fmt.Println(item)
	}
}
