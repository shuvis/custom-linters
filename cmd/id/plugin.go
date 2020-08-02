package main

import (
	"id-linter/internal/id"

	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin Plugin

type Plugin struct{}

func (p Plugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{id.Analyzer}
}
