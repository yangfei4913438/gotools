package main

import (
	"gotools/base"
	"fmt"
)

func main() {
	fmt.Println("this is shell test, begin!\n")
	res, out := base.ShExec(".", "pwd")
	fmt.Println("res: ", res)
	fmt.Println("out: ", out)
	fmt.Println("this is shell test, end!")
	
	
	
}
