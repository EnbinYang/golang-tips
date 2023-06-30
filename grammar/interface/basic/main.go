package main

import "fmt"

type Operation interface {
	getSum(int, int) int
}

type Robot struct{}

func (r Robot) getSum(a int, b int) int {
	return a + b
}

func foo(r Operation) {
	res := r.getSum(1, 1)
	fmt.Println("sum =", res)
}

func main() {
	var r Operation
	r = Robot{}
	foo(r)
}
