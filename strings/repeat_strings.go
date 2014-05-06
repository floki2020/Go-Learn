package main

import (
	"fmt"
	"strings"
)

//重复字符串s 用repeat方法

func main() {

	var oldStr string = "Hi Im Steve! "
	var newStr string

	newStr = strings.Repeat(oldStr, 3)

	fmt.Printf("The new repeated string is:%s\n", newStr)
}

//output
//The new repeated string is:Hi Im Steve! Hi Im Steve! Hi Im Steve!
