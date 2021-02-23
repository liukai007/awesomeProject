package main

import (
	"fmt"
	engine2 "src/github.com/dengsgo/math-engine/engine"
)

func main(){
	s:="2+1"
	fmt.Println(s)
	value,err:=engine2.ParseAndExec(s)
	if err==nil {
		fmt.Println(value)
	}else {
		fmt.Println("haha")
	}

}


