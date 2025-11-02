gen:
	oapi-codegen --config .openapi --include-tags tasks --package tasks openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --color=auto

run:
	go run ./cmd/main.go
