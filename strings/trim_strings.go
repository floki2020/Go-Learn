package main

import (
	"fmt"
	"strings"
)

//trim 修剪字符
func main() {

	var str string = "  Hey, I wanna be a hero!! "
	//var cutStr string = "Hey"

	fmt.Printf("The original string   is: %s\n", str)
	fmt.Printf("The trimspace string  is: %s\n", strings.TrimSpace(str))
	fmt.Printf("The sep string        is: %s\n", strings.Trim(str, " "))
	fmt.Printf("The trim left string  is: %s\n", strings.TrimLeft(str, " "))
	fmt.Printf("The trim right string is: %s\n", strings.TrimRight(str, " "))
}

//output
//	The original string   is:   Hey, I wanna be a hero!!
//	The trimspace string  is: Hey, I wanna be a hero!!
//	The sep string        is: Hey, I wanna be a hero!!
//	The trim left string  is: Hey, I wanna be a hero!!
//	The trim right string is:   Hey, I wanna be a hero!!
