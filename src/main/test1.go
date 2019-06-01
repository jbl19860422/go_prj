package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Walk(step int)
	Walk1(step int)
}

type People struct {
	Name     string
	CurrStep int
}

func (c People) PeInfo() {
	fmt.Println("people name:", c.Name, " currstep:", c.CurrStep)
}

func (c *People) SetName(name string) {
	c.Name = name
	fmt.Println("people new name:", name)
}

func (c People) Walk(step int) {
	c.PeInfo()
	c.CurrStep += step
	fmt.Println(" walk ", step)
}

func (c *People) Walk1(step int) {

}

func PrintMethodSet(i interface{}) {
	rt := reflect.TypeOf(i)
	fmt.Println("type is ", rt)
	for i, n := 0, rt.NumMethod(); i < n; i++ {
		fmt.Println("\tit has method:", rt.Method(i).Name)
	}
}

func LetWalk(a Animal, s int) {
	a.Walk(s)
}

func LetWalk1(p People, s int) {
	p.Walk(s)
}

func main() {
	var p = People{Name: "jiangbaolin"}
	PrintMethodSet(p)
	//下面两句等效
	p.PeInfo()
	/*
		指针调用值方法，指针包含所有方法集，指针调用值方法时，会通过指针创建值副本，然后再调用值方法;
		效果等同于：
		q := *(&p)
		q.PeInfo()
	*/
	(&p).PeInfo()
	People.PeInfo(p)
	// People.PeInfo(&p) //语句错误，这里需要显示传递指针
	//下面三句等效
	p.SetName("xiaoming") //p是值类型，不包含指针方法，但是go 隐式调用了(&p).SetName("jiang")语法糖，所以可以调用
	(&p).SetName("xiaojiang")
	(*People).SetName(&p, "liu")

	// var q Animal = p //p是值，不包含 Walk1指针方法，但Animal需要实现Walk1方法，所以报错
	// PrintMethodSet(q)
	var m Animal = &p //指针包含所有方法集，所以可以使用
	PrintMethodSet(m)
	//不会改变p
	LetWalk(&p, 10) //等效于LetWalk(*(&p), 10)
	p.PeInfo()
	// LetWalk(p, 10) //编译错误，p没法转为Animal
	// p.PeInfo()
	// var i interface{} = p //执行时报错，p没有实现Walk1
	// i.(Animal).Walk(10)

	var j interface{} = &p //执行时报错，p没有实现Walk1
	j.(Animal).Walk(10)    //正确

	LetWalk1(p, 10)
	p.PeInfo()
	// LetWalk1(&p, 10)  //语法错误，不接收值接收器
	p.PeInfo()
}
