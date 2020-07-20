package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_violates(t *testing.T) {
	tests := []struct {
		test     string
		name     string
		expected bool
	}{
		{"empty", "", false},
		{"violates", "Id", true},
		{"correct case", "ID", false},
		{"id in the end", "SomeId", true},
		{"id in the middle", "SomeIdViolation", true},
		{"id in the start", "IdViolation", true},
		{"id is part of another word", "IdentifierIdentifier", false},
		{"id is part of another word and also there is bad pattern", "IdentifierId", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := violates(tt.name)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSuggestedFixes(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(wd), "testdata")
	aggregator := &errorAggregator{}
	analysistest.RunWithSuggestedFixes(aggregator, testdata, Analyzer, "id")
	assert.Equal(t, expectedErrors, aggregator.errors)
}

type errorAggregator struct {
	errors []string
}

func (n *errorAggregator) Errorf(format string, args ...interface{}) {
	n.errors = append(n.errors, fmt.Sprintf(format, args...))
}

var (
	globalConstError              = "id/id.go:3:7: unexpected diagnostic: ID should be uppercase: constId -> constID"
	globalVarError                = "id/id.go:5:5: unexpected diagnostic: ID should be uppercase: myId -> myID"
	structNameError               = "id/id.go:7:6: unexpected diagnostic: ID should be uppercase: StructWithId -> StructWithID"
	structFieldError              = "id/id.go:8:2: unexpected diagnostic: ID should be uppercase: Id -> ID"
	interfaceNameError            = "id/id.go:11:6: unexpected diagnostic: ID should be uppercase: InterfaceWithId -> InterfaceWithID"
	interfaceMethodNameError      = "id/id.go:12:2: unexpected diagnostic: ID should be uppercase: GetById -> GetByID"
	interfaceMethodParameterError = "id/id.go:12:10: unexpected diagnostic: ID should be uppercase: myId -> myID"
	returnParameterError          = "id/id.go:12:24: unexpected diagnostic: ID should be uppercase: otherId -> otherID"
	funcNameError                 = "id/id.go:15:6: unexpected diagnostic: ID should be uppercase: GetId -> GetID"
	mapFuncNameError              = "id/id.go:17:6: unexpected diagnostic: ID should be uppercase: MapSomethingToId -> MapSomethingToID"
	localVariableError            = "id/id.go:18:2: unexpected diagnostic: ID should be uppercase: someId -> someID"
	returnVariableError           = "id/id.go:19:9: unexpected diagnostic: ID should be uppercase: myId -> myID"
	returnVariableError2          = "id/id.go:19:16: unexpected diagnostic: ID should be uppercase: someId -> someID"
	funcReturnParameterError      = "id/id.go:22:15: unexpected diagnostic: ID should be uppercase: MyId -> MyID"
	returnConstError              = "id/id.go:23:9: unexpected diagnostic: ID should be uppercase: constId -> constID"
	mapEmptyFunctionError         = "id/id.go:26:6: unexpected diagnostic: ID should be uppercase: MyIdAndSomeOtherIdAndId -> MyIDAndSomeOtherIDAndID"

	expectedErrors = []string{
		globalConstError, globalVarError, structNameError, structFieldError, interfaceNameError, interfaceMethodNameError,
		interfaceMethodParameterError, returnParameterError, funcNameError, mapFuncNameError, localVariableError,
		returnVariableError, returnVariableError2, funcReturnParameterError, returnConstError, mapEmptyFunctionError}
)
