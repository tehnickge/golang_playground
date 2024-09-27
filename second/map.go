package main

import "fmt"

func MapTest() {
	var money map[string]int = map[string]int{
		"dollars": 32131,
		"kejk": 33,
	}

	el, status := money["kek"]
	fmt.Println(el,status)
}

func TestFunc(a int, b string) {
	fmt.Println(fmt.Sprintf("%d",a), b)
}

func TestReturnManyArgs(a int, b int) (sum int, sub int, mul int, div int) {
	sum = a + b
	div = a / b
	sub = a - b
	mul = a * b
	return
}