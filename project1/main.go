package main

import (
	"fmt"
	"mymath"
)

type student struct {
	Age  int
	name string
}

func (t *student) SayInfoAfterYears(deltyear int) {
	nowage := t.Age + deltyear
	fmt.Println("nowage", nowage)
	fmt.Printf("my name is %s and i will be %d years old after %d years\n", t.name, mymath.MyAdd(t.Age, deltyear), deltyear)
}

func main() {
	var st1 *student = &student{Age: 22, name: "xiaoming"}
	st1.SayInfoAfterYears(2)
	fmt.Println("hello world!")
}
