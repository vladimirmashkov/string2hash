package string2hash

import (
	"fmt"
	"os"
	"strings"
)

func PrintHelp(parameters *Parameters) {
	var helpCmd string
	argsCount := len(os.Args)

	if argsCount == 1 {
		helpCmd = "--help"
		parameters.help = true
	}
	if argsCount == 2 {
		helpCmd = os.Args[1]
		if helpCmd == ("-h") || helpCmd == ("man") {
			helpCmd = "--help"
			parameters.help = true
		}
		if helpCmd != "--help" {
			parameters.error = true
		}
	}

	if argsCount > 3 {
		parameters.error = true
	}

	if parameters.error == true {
		fmt.Println("The number of parameters cannot exceed two")
		fmt.Println("use help: string2hash --help")
		return
	}

	if helpCmd == "--help" {
		parameters.help = true
		hlf := strings.Split(HashFunctionsList, ",")
		fmt.Println(`================================+
GNU v.3.0; string2hash 1.0      |
author: Vladimir Mashkov        |
mailto:vladimir@mashkov.com     |
================================+

Usage:   string2hash [hash function] [text for hash]
===========================================+
Examples:                                  |
    string2hash md5 text                   |
    string2hash md5 'the long long text'   |
    string2hash md5 "the long long text"   |
===========================================+

[hash function]
    - list of supported hash functions:
      ` + (strings.Join(hlf[:], ", ")))
	}
}
