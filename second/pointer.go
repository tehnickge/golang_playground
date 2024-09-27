package main

import "fmt"

func TestPointer() {
	a := 5
	pointer := &a
	fmt.Println("a:",a,"pointer a:",&a,"p:",pointer,"p*:",*pointer)
}
