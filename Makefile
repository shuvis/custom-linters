plugin:
	@go build -o out/plugins/id-plugin.so -buildmode=plugin cmd/id/plugin/id-plugin.go

cli:
	@go build -o out/cli/id-linter cmd/id/id-linter/id-linter.go
