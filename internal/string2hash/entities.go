package string2hash

import (
	"os"
	"strings"
)

const HashFunctionsList = "md5,sha1"

type Parameters struct {
	hash   string
	text   string
	result string
	base64 string
	help   bool
	error  bool
}

func NewParameters() Parameters {
	argsCount := len(os.Args)
	if argsCount == 3 {
		return Parameters{strings.ToLower(os.Args[1]), os.Args[2], "", "", false, false}
	} else {
		return Parameters{"", "", "", "", false, false}
	}
}
