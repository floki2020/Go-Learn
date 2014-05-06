package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string = "Hey,I'm Steve,Nice to meet you! "
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", str)

	lower = strings.ToLower(str)
	fmt.Printf("The lowercase string is: %s\n", lower)

	upper = strings.ToUpper(str)
	fmt.Printf("The uppercase string is: %s\n", upper)
}

//output
//	The original string is: Hey,I'm Steve,Nice to meet you!
//	The lowercase string is: hey,i'm steve,nice to meet you!
//	The uppercase string is: HEY,I'M STEVE,NICE TO MEET YOU!
