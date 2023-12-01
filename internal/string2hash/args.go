package string2hash

import (
	"fmt"
	"os"
)

func Argguments() {
	args := os.Args
	count := 0
	for _, v := range args {
		count++
		fmt.Println(v)
	}

	fmt.Println(count)

	//fmt.Println(args)
	//fmt.Println(args[1])
}
