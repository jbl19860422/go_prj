package main

import (
	"runtime"
	"log"
	"fmt"
	// "time"
)

func traceMemStats() {
    var ms runtime.MemStats
    runtime.ReadMemStats(&ms)
    log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

type PARENT struct {
	childs []*CHILD
}

func (p *PARENT) removeChild(c *CHILD) {
	for i := 0; i < len(p.childs); i++ {
		if c == p.childs[i] {
			p.childs = append(p.childs[0:i], p.childs[i+1:]...)
			fmt.Println("remove child ok,childs len=", len(p.childs))
			break;
		}
	}
	runtime.GC()
}

type CHILD struct {
	parent *PARENT
	data	int32
	arr 	[10000000]int
}

func (c *CHILD) removeSelf() {
	c.parent.removeChild(c)
	fmt.Println("child removed")
	runtime.GC()
	c.data = 111
	fmt.Println("child's data=", c.data)
}

func main() {
	// traceMemStats()
	{
		p := &PARENT{}
		p.childs = make([]*CHILD, 0)

		c := &CHILD{}
		c.parent = p
		p.childs = append(p.childs, c)
		for i := 0; i < 10000000; i++ {
			c.arr[i] = i
		}
		traceMemStats()
		c.removeSelf()
		traceMemStats()
	}
	traceMemStats()
	// traceMemStats()
	runtime.GC()
	traceMemStats()
	// traceMemStats()
	// for i := 0; i < 10000000; i++ {
	// 	a := &A{}
	// 	b := &B{}
	// 	a.b = b
	// 	b.a = a
	// 	runtime.GC()
	// 	time.Sleep(time.Millisecond*1000)
	// 	traceMemStats()
	// }
}