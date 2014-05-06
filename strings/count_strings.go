package main

import (
	"fmt"
	"strings"
)

//查看字符 在字符串中出现的次数

func main() {

	var str string = "This is a test demo!!!"
	var sep string = "s"

	fmt.Printf("The count of \"%s\" in \"%s\" is %d", sep, str, strings.Count(str, sep))
}

//output
//The count of "s" in "This is a test demo!!!" is 3
