test:
	@go test ./... -coverprofile=cover.out -coverpkg=./...

coverage:
	$(MAKE) test
	@go tool cover -func cover.out | grep total | awk '{print $3}'

run:
	@go run cmd/main.go
