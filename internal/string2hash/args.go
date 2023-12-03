package string2hash

import (
	"fmt"
	"strings"
)

func CheckParameters(parameters *Parameters) {
	if parameters.error == false && parameters.help == false {
		hlf := strings.Split(HashFunctionsList, ",")
		parameters.error = true
		for _, v := range hlf {
			if v == parameters.hash {
				parameters.error = false
			}
		}
		if parameters.error == true {
			fmt.Println("The parameter " + parameters.hash + " is not supported")
			fmt.Println("The list of supported hash functions is: " + (strings.Join(hlf[:], ", ")))
			fmt.Println("Please, read the help: string2hash --help")
		}
	}

}
