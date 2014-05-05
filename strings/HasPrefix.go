package main

import (
	"fmt"
	"strings"
)

//判断字符串 s 是否是以prefix开头 HasPrefix
//判断字符串 s 是否是以suffix结束 HasSuffix
//判断是否包含 substr Contains

func main() {
	var str string = "This is an example of a string"

	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")

	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("\n %t", strings.HasSuffix(str, "n"))

	fmt.Printf("\n Does the strings \"%s\" contains %s? ", str, "ex")

	fmt.Printf("%t\n", strings.Contains(str, "ex"))
}

//output
//T/F? Does the string "This is an example of a string" have prefix Th? true  false
