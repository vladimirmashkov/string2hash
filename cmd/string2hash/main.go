package main

import (
	"github.com/vladimirmashkov/string2hash/internal/string2hash"
)

func main() {
	Params := string2hash.NewParameters()
	string2hash.PrintHelp(&Params)
	string2hash.CheckParameters(&Params)
	string2hash.GetResult(&Params)
}
