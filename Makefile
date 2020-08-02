plugin:
	@go build -o out/plugin/id.so -buildmode=plugin cmd/id/plugin.go

cli:
	@go build -o out/cli/id-linter cmd/id/cli.go
