package main

import (
	"fmt"
	
)

func PirntAnyText(a int, ch chan int) {
	ch <- a
	fmt.Println("any with :",a)
}

func TestThirds() {
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go PirntAnyText(i,ch)
		fmt.Println("ch",<-ch)
	}
}