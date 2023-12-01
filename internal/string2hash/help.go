package string2hash

import (
	"fmt"
	"strings"
)

func PrintHelp() {
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("author: Vladimir Mashkov")
	fmt.Println("mailto:vladimir@mashkov.com")
	fmt.Println(strings.Repeat("=", 10))
	//fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("string2hash -f md5 -s \"the test\" ")

}
