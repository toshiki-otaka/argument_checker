package main

import (
	"github.com/toshiki-otaka/argument_checker"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(argument_checker.Analyzer) }
