package main 
import(
	"fmt"
)

type TestInt interface {
	Sum() int 
	Muliply() float64 
	Division() float64
	Substract() int
}

type TestStruct struct {
	numberFirst int
	numberSecond int
}

func (n TestStruct) Sum() int {
	return n.numberFirst + n.numberSecond
}
func (n TestStruct) Substract() int {
	return n.numberFirst - n.numberSecond
}
func (n TestStruct) Division() float64 {
	return float64(n.numberFirst) / float64(n.numberSecond)
}
func (n TestStruct) Muliply() float64 {
	return float64(n.numberFirst) * float64(n.numberSecond)
}

func MakesInterfaces() {
	var inter TestInt
	struc := TestStruct{87,42}
	inter = struc 
	fmt.Println(inter.Division(),inter.Muliply(),inter.Sum(),inter.Substract())


}