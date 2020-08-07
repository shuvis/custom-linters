package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"id-linter/internal/id"
)

func main() { singlechecker.Main(id.Analyzer) }
