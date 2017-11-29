
test:
	@go test -v -race

coverage:
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic

