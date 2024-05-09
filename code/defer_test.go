package code

import (
	"fmt"
	"testing"
)

// https://www.yuque.com/aceld/golang/qnubsg
// 1
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

// 4
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// 3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

// t=1 0 1
func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

type Student struct {
	Name string
}

var list map[string]Student

func TestMain123(t *testing.T) {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	student.Name = "Aceld2"
	list["student"] = student
	fmt.Println(list["student"])
}
