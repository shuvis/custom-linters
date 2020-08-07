package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/shuvis/custom-linters/internal/id"
)

func main() { singlechecker.Main(id.Analyzer) }
