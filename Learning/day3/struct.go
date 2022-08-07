package main

import "fmt"

//自定义Person结构体

type Person struct {
	firstname string
	lastname  string
}

//(p *Person)是Print方法的接收者，p名称可以自定义理解为python中self，p通常定义为指针节省内存
func (p *Person) Printme() {
	fmt.Println("person name is :", p.firstname, p.lastname)
}

//定义接口Namer，调用方法Printme
type Namer interface {
	Printme()
}

//定义结构体Animal
type Animal struct {
	animals string
	sex     string
}

//定义方法Printme，传入的接受者类型是Animal指针类型
func (a *Animal) Printme() {
	fmt.Printf("Anilmal is: %s, sex is:%s \n", a.animals, a.sex)
}

func main() {
	//使用结构体Person初始化p，此处传递的是地址
	p := &Person{firstname: "黄", lastname: "飞鸿"}
	//调用方法Print，指针p会自动解引用
	p.Printme()

	//另外一种初始化结构体的简略写法
	a := &Animal{"猫", "雄性"}
	fmt.Println("--------------------")

	//使用接口Namer初始化变量aa，并赋值a
	var aa Namer = a
	//使用接口Namer初始化变量pp，并赋值p
	var pp Namer = p

	//根据传入的接受者调用对应的方法
	aa.Printme()
	pp.Printme()
	fmt.Println("--------------------")

	//初始化切片定义类型为接口
	who := []Namer{p, a}
	for k, _ := range who {
		who[k].Printme()
	}
}
