package main

import "fmt"

type Handler interface {
	HandleRequest(int)
	SetNext(Handler)
}

type ConcreteHandler1 struct {
	next Handler
}

func (c *ConcreteHandler1) SetNext(next Handler) {
	c.next = next
}

func (c *ConcreteHandler1) HandleRequest(request int) {
	if request >= 0 && request <= 10 {
		fmt.Printf("%T 处理请求 %d\n", c, request)
	} else if c.next != nil {
		c.next.HandleRequest(request)
	} else {
		fmt.Println("无法处理该请求")
	}
}

type ConcreteHandler2 struct {
	next Handler
}

func (c *ConcreteHandler2) SetNext(next Handler) {
	c.next = next
}

func (c *ConcreteHandler2) HandleRequest(request int) {
	if request >= 11 && request <= 20 {
		fmt.Printf("%T 处理请求 %d\n", c, request)
	} else if c.next != nil {
		c.next.HandleRequest(request)
	} else {
		fmt.Println("无法处理该请求")
	}
}

type ConcreteHandler3 struct {
	next Handler
}

func (c *ConcreteHandler3) SetNext(next Handler) {
	c.next = next
}

func (c *ConcreteHandler3) HandleRequest(request int) {
	if request >= 21 && request <= 30 {
		fmt.Printf("%T 处理请求 %d\n", c, request)
	} else if c.next != nil {
		c.next.HandleRequest(request)
	} else {
		fmt.Println("无法处理该请求")
	}
}

func main() {
	h1 := &ConcreteHandler1{}
	h2 := &ConcreteHandler2{}
	h3 := &ConcreteHandler3{}

	h1.SetNext(h2)
	h2.SetNext(h3)

	requests := []int{2, 17, 25, 42}

	for _, request := range requests {
		h1.HandleRequest(request)
	}
}
