package main

import (
	"fmt"
	"sort"
)

func SliceTest() {
	var array []string
	array = append(array, "asd")
	array = append(array, "b")
	array = append(array, "c")
	array = append(array, "abc")
	array = append(array, "asd")
	sort.Strings(array)
	fmt.Println(array)
}