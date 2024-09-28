package main

import (
	"fmt"
	"slices"
)

type User struct {
	name     string
	age      int64
	password string
	score []int
}
func (u User) GetName() string {
	return u.name
}
func (u *User) SetName(n string) {
	u.name = n
}
func (u *User) IsElder() bool {
	var isElder bool = false
	if (u.age >= 18) { isElder = true }
	return isElder  
}
func (u User) GetMaxScore() int {
	return slices.Max(u.score)
}
func (u User) GetSumScore() int {
	sum := 0
	for _, sc := range u.score {
		sum += sc
	}
	return sum
}
func TestStruckture() {
	var firstUser User = User{name: "jhon", age: 23, password: "sdfs", score:  []int{14,42,21}}
	secondUser := User{ name: "ekk", age: 22, password: "pas", score:  []int{22,34,121,4343}}
	firstUser.SetName("LOH")
	fmt.Println(firstUser, secondUser)
	println(firstUser.GetMaxScore(), firstUser.GetSumScore())
	
}