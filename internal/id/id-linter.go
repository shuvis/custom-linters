package id

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	invalidPattern = "Id"
	validPattern   = "ID"
)

var (
	wrongPatternRegex = regexp.MustCompile(invalidPattern)
	Analyzer          = &analysis.Analyzer{
		Name:     "id",
		Doc:      "Checks that ID is uppercase",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
)

func run(pass *analysis.Pass) (interface{}, error) {
	visitor := func(node ast.Node) bool {
		if identifier, ok := node.(*ast.Ident); ok {
			checkID(identifier, pass)
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, visitor)
	}
	return nil, nil
}

func checkID(ident *ast.Ident, pass *analysis.Pass) {
	if violates(ident.Name) {
		pass.Report(diagnose(ident))
	}
}

func violates(name string) bool {
	for _, bounds := range wrongPatternRegex.FindAllIndex([]byte(name), -1) {
		if substringViolation(name, bounds[0]) {
			return true
		}
	}
	return false
}

func substringViolation(name string, start int) bool {
	if hasSubsequentLetter(name, start) {
		return nextCharIsUppercase(name, start)
	}
	return true
}

func hasSubsequentLetter(name string, start int) bool {
	return start+len(invalidPattern) < len(name)
}

func nextCharIsUppercase(name string, start int) bool {
	return unicode.IsUpper(rune(name[start+len(invalidPattern)]))
}

func diagnose(ident *ast.Ident) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:            ident.Pos(),
		Message:        violationMessage(ident),
		SuggestedFixes: suggestFixes(ident),
	}
}

func violationMessage(ident *ast.Ident) string {
	return fmt.Sprintf("ID should be uppercase: %s -> %s", ident.Name, replaceID(ident.Name))
}

func suggestFixes(ident *ast.Ident) []analysis.SuggestedFix {
	return []analysis.SuggestedFix{
		{Message: "Rename", TextEdits: []analysis.TextEdit{
			{Pos: ident.Pos(), End: ident.End(), NewText: []byte(replaceID(ident.Name))},
		}},
	}
}

func replaceID(name string) string {
	return strings.Replace(name, invalidPattern, validPattern, -1)
}
