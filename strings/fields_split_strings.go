package main

import (
	"fmt"
	"strings"
)

//Fields 去除空格
//split  分割sep的字符

func main() {
	var str string = "  Hey, I wanna be a hero!! "

	sl := strings.Fields(str)

	for _, v := range sl {
		fmt.Printf("%s", v)
	}

	fmt.Print("\n")

	sl1 := strings.Split(str, " ")
	for _, v := range sl1 {
		fmt.Printf("%s", v)
	}
}
