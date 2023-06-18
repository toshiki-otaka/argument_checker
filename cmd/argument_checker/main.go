package main

import (
	"argument_checker"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(argument_checker.Analyzer) }
