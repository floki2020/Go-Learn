package main

import (
	"fmt"
	"strings"
)

//repalce 用户替换字符 n为个数 n=-1 替换所有
func main() {

	var str string = "Tom is a girl"
	var old string = "girl"
	var news string = "boy"

	fmt.Printf("The old strings \"%s\" changed is: %s", str, strings.Replace(str, old, news, 1))
}

//output
//The old strings "Tom is a girl" changed is: Tom is a boy
