package main

import (
	"fmt"
	engine2 "github.com/dengsgo/math-engine/engine"
	"strconv"
)

func main() {
	//计算 公式
	s := "2+1"
	fmt.Println(s)
	value, err := engine2.ParseAndExec(s)
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println("haha")
	}

	//得到byte的
	a := 0x15
	fmt.Printf("%b", a)
	fmt.Println()
	fmt.Printf("%d", a)
	fmt.Println()
	fmt.Printf("%o", a)
	fmt.Println()

	var b int64 = 11
	fmt.Printf("%d", b)
	fmt.Println()

	s2 := strconv.FormatInt(b, 2)
	fmt.Printf("%v", s2)

	var b1 int64 = 0x15
	fmt.Printf("%d", b1)
	fmt.Println()

	s21 := strconv.FormatInt(b1, 2)
	fmt.Printf("%v", s21)

}
