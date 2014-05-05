package main

import (
	"fmt"
	"strings"
)

//查询出所在字符串中的位置

func main() {

	var str string = "This is a test demo strings"

	fmt.Printf("The position of \"demo\" is :")
	fmt.Printf("%d\n", strings.Index(str, "demo"))

}

//output
//The position of "demo" is :15
